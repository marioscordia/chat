package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // to initialize connection
	"github.com/pressly/goose/v3"

	"github.com/marioscordia/chat/closer"
	"github.com/marioscordia/chat/delivery/grpc"
	"github.com/marioscordia/chat/facility"
	repo "github.com/marioscordia/chat/repository"
	"github.com/marioscordia/chat/repository/postgres"
	"github.com/marioscordia/chat/service"
	"github.com/marioscordia/chat/service/chat"
)

const (
	dbPostgresDriverName   = "postgres"
	migrationsPostgresPath = "db/migrations"
)

type provider struct {
	config *facility.Config

	db *sqlx.DB

	chatRepo repo.ChatRepository

	chatService service.ChatService

	chatHandler *grpc.Handler
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
			if err := goose.SetDialect("postgres"); err != nil {
				log.Fatalf("failed to set postgres dialect for goose: %v", err)
			}
			if err := goose.Up(db, migrationsPostgresPath); err != nil && !errors.Is(err, goose.ErrAlreadyApplied) {
				log.Fatalf("failed to apply migrations: %v", err)
			}
		}

		closer.Add(dbx.Close)

		p.db = dbx
	}

	return p.db
}

func (p *provider) ChatRepository(ctx context.Context) repo.ChatRepository {
	if p.chatRepo == nil {
		repo, err := postgres.New(ctx, p.NewDB())
		if err != nil {
			log.Fatalf("failed to initialize chat repository: %v", err)
		}

		p.chatRepo = repo
	}

	return p.chatRepo
}

func (p *provider) ChatService(ctx context.Context) service.ChatService {
	if p.chatService == nil {
		p.chatService = chat.New(p.ChatRepository(ctx))
	}

	return p.chatService
}

func (p *provider) ChatHandler(ctx context.Context) *grpc.Handler {
	if p.chatHandler == nil {
		p.chatHandler = grpc.New(p.ChatService(ctx))
	}

	return p.chatHandler
}
