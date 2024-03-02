-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "name" character varying(100) NULL, PRIMARY KEY ("id"));
-- Create "blog_posts" table
CREATE TABLE "blog_posts" ("id" uuid NOT NULL, "title" character varying(100) NULL, "body" text NULL, "author_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "author_fk" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
