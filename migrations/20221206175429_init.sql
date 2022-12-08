-- +goose Up
-- +goose StatementBegin
create table channels
(
    id int NOT NULL AUTO_INCREMENT,
    channel_id varchar(255),
    channel_name varchar(255),
    type varchar(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    UNIQUE (channel_id)
);
-- +goose StatementEnd
-- +goose StatementBegin
create table phrases
(
    id int NOT NULL AUTO_INCREMENT,
    sender_chan_id varchar(255) null,
    sender_id varchar(255) null,
    sender_name varchar(255) null,
    phrase text null,
    reply_id varchar(255) null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id),
    INDEX (sender_chan_id, sender_id),
    FOREIGN KEY (sender_chan_id)
        REFERENCES channels(channel_id)
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table phrases;
-- +goose StatementEnd
-- +goose StatementBegin
drop table channels;
-- +goose StatementEnd