package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/marioscordia/chat"
	"github.com/marioscordia/chat/internal/constants"
	"github.com/marioscordia/chat/pkg/chat_v1"
)

// New is ...
func New(ctx context.Context, db *sqlx.DB) (chat.Repository, error) {
	stmtDeleteChat, err := db.PreparexContext(ctx, "update chats set deleted_at=$1 where id=$2")
	if err != nil {
		return nil, err
	}

	stmtDeleteMember, err := db.PreparexContext(ctx, "update chat_members set deleted_at=$1 where chat_id=$2 and member_id=$3")
	if err != nil {
		return nil, err
	}

	stmtCreateMsg, err := db.PreparexContext(ctx, `insert into messages (chat_id, author_id, msg_text, created_at, updated_at)
											       values($1, $2, $3, $4, $5)`)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:               db,
		stmtDeleteChat:   stmtDeleteChat,
		stmtDeleteMember: stmtDeleteMember,
		stmtCreateMsg:    stmtCreateMsg,
	}, nil
}

type repository struct {
	db               *sqlx.DB
	stmtDeleteChat   *sqlx.Stmt
	stmtDeleteMember *sqlx.Stmt
	stmtCreateMsg    *sqlx.Stmt
}

func (r *repository) CreateChat(ctx context.Context, chat *chat_v1.CreateRequest) (int64, error) {
	t := time.Now()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	var id int64

	query := `insert into chats (title, creator_id, chat_type, created_at, updated_at)
			  values ($1, $2, $3, $4, $5) returning id`

	if err := tx.GetContext(ctx, &id, query, chat.ChatName, chat.CreatorId, chat.ChatType, t, t); err != nil {
		return 0, err
	}

	query = `insert into chat_members (chat_id, member_id, roles, created_at, updated_at)
			 values ($1, $2, $3, $4, $5)`

	if _, err := tx.ExecContext(ctx, query, id, chat.CreatorId, constants.ChannelMemberRoleAdmin, t, t); err != nil {
		return 0, err
	}

	for _, userID := range chat.UserIds {
		if _, err := tx.ExecContext(ctx, query, id, userID, constants.ChannelMemberRoleMember, t, t); err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) DeleteMember(ctx context.Context, chatID, memberID int64) error {
	t := time.Now()

	_, err := r.stmtDeleteMember.ExecContext(ctx, t, chatID, memberID)

	return err
}

func (r *repository) DeleteChat(ctx context.Context, chatID int64) error {
	t := time.Now()

	_, err := r.stmtDeleteChat.ExecContext(ctx, t, chatID)

	return err
}

func (r *repository) CreateMessage(ctx context.Context, msg *chat_v1.Message) error {
	t := time.Now()

	_, err := r.stmtCreateMsg.ExecContext(ctx, msg.ChatId, msg.AuthorId, msg.Text, t, t)

	return err
}
