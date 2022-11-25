CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "topics" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "topic_name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "owner_id" uuid NOT NULL,
  "topic_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "content" varchar NOT NULL,
  "title" varchar NOT NULL
);

CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "owner_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  "main_comment_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "content" varchar NOT NULL
);

CREATE TABLE "likes" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "user_id" uuid,
  "liked_id" uuid,
  "type" tinyint,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "followship" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4 ()),
  "follower_id" uuid,
  "topic_id" uuid,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "comment"."main_comment_id" IS 'if this field is 0 mean it is the top-level comment';

COMMENT ON COLUMN "likes"."type" IS '1 is posts, 2 is comment';

ALTER TABLE "posts" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("id") REFERENCES "comment" ("main_comment_id");

ALTER TABLE "likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "followship" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "followship" ADD FOREIGN KEY ("topic_id") REFERENCES "topics" ("id");
