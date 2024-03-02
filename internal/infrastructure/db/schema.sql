-- create "users" table
CREATE TABLE users (
  id uuid NOT NULL,
  name varchar(100) NULL,
  PRIMARY KEY (id)
);

-- create "blog_posts" table
CREATE TABLE blog_posts (
  id uuid NOT NULL,
  title varchar(100) NULL,
  body text NULL,
  author_id uuid NULL,
  PRIMARY KEY (id),
  CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES users(id)
);