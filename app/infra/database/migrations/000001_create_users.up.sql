create table users (
    id int not null  PRIMARY KEY AUTO_INCREMENT,
    login_id varchar(10) not null UNIQUE,
    name varchar(50),
    password varchar(100) not null ,
    mail varchar(100),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    create_user int,
    updated_at timestamp NULL ON UPDATE CURRENT_TIMESTAMP,
    updated_user int,
    deleted_flg BOOLEAN,
    deleted_at timestamp NULL,
    deleted_user int
);