package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // to initialize connection
	"github.com/pressly/goose/v3"

	"github.com/marioscordia/chat/internal/api"
	"github.com/marioscordia/chat/internal/closer"
	"github.com/marioscordia/chat/internal/config"
	repository "github.com/marioscordia/chat/internal/repository"
	"github.com/marioscordia/chat/internal/repository/postgres"
	"github.com/marioscordia/chat/internal/service"
	"github.com/marioscordia/chat/internal/service/chat"
)

const (
	dbPostgresDriverName   = "postgres"
	migrationsPostgresPath = "db/migrations"
)

type provider struct {
	config *config.Config

	db *sqlx.DB

	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatHandler *api.Handler
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) NewDB() *sqlx.DB {
	if p.db == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			p.config.PostgresHost, p.config.PostgresPort, p.config.PostgresUser, p.config.PostgresPassword, p.config.PostgresDb, p.config.PostgresSslMode)

		db, err := sql.Open(dbPostgresDriverName, psqlInfo)
		if err != nil {
			log.Fatalf("failed to connect to postgres db: %v", err)
		}

		dbx := sqlx.NewDb(db, dbPostgresDriverName)

		if err = dbx.Ping(); err != nil {
			log.Fatalf("failed to verify connection: %v", err)

		}

		if p.config.PostgresMigrate {
			if err = goose.SetDialect("postgres"); err != nil {
				log.Fatalf("failed to set postgres dialect for goose: %v", err)
			}

			if err = goose.Up(db, migrationsPostgresPath); err != nil && !errors.Is(err, goose.ErrAlreadyApplied) {
				log.Fatalf("failed to apply migrations: %v", err)
			}
		}

		closer.Add(dbx.Close)

		p.db = dbx
	}

	return p.db
}

func (p *provider) ChatRepository() repository.ChatRepository {
	if p.chatRepository == nil {
		repo, err := postgres.New(p.NewDB())
		if err != nil {
			log.Fatalf("failed to initialize chat repository: %v", err)
		}

		p.chatRepository = repo
	}

	return p.chatRepository
}

func (p *provider) ChatService() service.ChatService {
	if p.chatService == nil {
		p.chatService = chat.New(p.ChatRepository())
	}

	return p.chatService
}

func (p *provider) ChatHandler() *api.Handler {
	if p.chatHandler == nil {
		p.chatHandler = api.New(p.ChatService())
	}

	return p.chatHandler
}
