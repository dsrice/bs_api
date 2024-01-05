create table tokens (
    id int not null  PRIMARY KEY AUTO_INCREMETN,
    user_id int not null ,
    token varchar(100),
    refresh_token varchar(100),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    create_user int,
    updated_at timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    updated_user int,
    deleted_flg BOOLEAN,
    deleted_at timestamp NULL,
    deleted_user int,
    FOREIGN KEY (user_id) REFERENCES users (id)
);