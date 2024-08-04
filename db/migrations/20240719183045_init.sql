-- +goose Up
-- +goose StatementBegin
create table if not exists chats (
    id serial primary key,
    title text not null,
    creator_id bigint not null,
    total_msg_count bigint default 0,
    chat_type varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

create table if not exists chat_members (
    chat_id bigint not null,
    member_id bigint not null,
    msg_count bigint default 0,
    roles varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    foreign key (chat_id) references chats(id) on delete cascade
);

create table if not exists messages (
    id serial primary key,
    chat_id bigint not null,
    author_id bigint not null,
    msg_text text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    foreign key (chat_id) references chats(id) on delete cascade
);

create unique index if not exists chat_members_chat_id_member_id_uindex
    on chat_members (chat_id, member_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists chat_members_channel_id_member_id_uindex;

drop table if exists messages;

drop table if exists chat_members;

drop table if exists chats;
-- +goose StatementEnd
