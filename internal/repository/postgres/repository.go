package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/marioscordia/chat/internal/constants"
	"github.com/marioscordia/chat/internal/model"
	repo "github.com/marioscordia/chat/internal/repository"
)

// New is a function that returns ChatRepository object
func New(db *sqlx.DB) (repo.ChatRepository, error) {

	return &repository{
		db: db,
	}, nil
}

type repository struct {
	db *sqlx.DB
}

func (r *repository) CreateChat(ctx context.Context, chat *model.ChatCreate) (int64, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	var id int64

	query := `insert into chats (title, creator_id, chat_type)
			  values ($1, $2, $3) returning id`

	if err := tx.GetContext(ctx, &id, query, chat.Name, chat.CreatorID, chat.Type); err != nil {
		return 0, err
	}

	query = `insert into chat_members (chat_id, member_id, roles)
			 values ($1, $2, $3)`

	if _, err := tx.ExecContext(ctx, query, id, chat.CreatorID, constants.ChannelMemberRoleAdmin); err != nil {
		return 0, err
	}

	for _, userID := range chat.UserIDs {
		if _, err := tx.ExecContext(ctx, query, id, userID, constants.ChannelMemberRoleMember); err != nil {
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

	deleteMember, err := r.db.PreparexContext(ctx, "update chat_members set deleted_at=$1 where chat_id=$2 and member_id=$3")
	if err != nil {
		return err
	}

	_, err = deleteMember.ExecContext(ctx, t, chatID, memberID)

	return err
}

func (r *repository) DeleteChat(ctx context.Context, chatID int64) error {
	t := time.Now()

	deleteChat, err := r.db.PreparexContext(ctx, "update chats set deleted_at=$1 where id=$2")
	if err != nil {
		return err
	}

	_, err = deleteChat.ExecContext(ctx, t, chatID)

	return err
}

func (r *repository) CreateMessage(ctx context.Context, msg *model.Message) error {
	createMsg, err := r.db.PreparexContext(ctx, `insert into messages (chat_id, author_id, msg_text)
											      				values($1, $2, $3)`)
	if err != nil {
		return err
	}

	_, err = createMsg.ExecContext(ctx, msg.ChatID, msg.UserID, msg.Text)

	return err
}
