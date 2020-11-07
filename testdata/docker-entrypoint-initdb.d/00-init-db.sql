CREATE DATABASE IF NOT EXISTS linkee_db;

use linkee_db;

CREATE TABLE users (
    id varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE pages (
    id varchar(255) NOT NULL,
    slug varchar(255) NOT NULL,
    user_id varchar(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT `fk_users` FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE (slug, user_id)
);

CREATE TABLE links (
    id varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    url varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    counter INT(10) NOT NULL DEFAULT 0,
    slug varchar(255) NOT NULL,
    page_id varchar(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT `fk_pages` FOREIGN KEY (page_id) REFERENCES pages(id),
    UNIQUE (slug, page_id)
);

CREATE INDEX idx_links_slug ON links(slug, page_id);
CREATE INDEX idx_pages_slug ON pages(slug, user_id);
CREATE INDEX idx_users_username ON users(username);