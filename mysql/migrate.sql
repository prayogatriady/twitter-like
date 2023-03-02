CREATE DATABASE IF NOT EXISTS twitter_like;

USE twitter_like;

CREATE TABLE IF NOT EXISTS  users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    profile_pict VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS  tweets (
    id BIGINT NOT NULL AUTO_INCREMENT,
    content TEXT,
    user_id BIGINT NOT null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (user_id),
    PRIMARY KEY (id),
    foreign key (user_id) references users (id) ON DELETE restrict on update CASCADE
);

CREATE TABLE IF NOT EXISTS  follows (
    id BIGINT NOT NULL AUTO_INCREMENT,
    follower_id BIGINT NOT null,
    following_id BIGINT NOT null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (id),
    foreign key (follower_id) references users (id) ON DELETE restrict on update CASCADE,
    foreign key (following_id) references users (id) ON DELETE restrict on update CASCADE
);