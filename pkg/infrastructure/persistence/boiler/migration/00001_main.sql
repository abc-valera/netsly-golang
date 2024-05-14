-- +goose Up
-- +goose StatementBegin
CREATE TABLE "User" (
  "id" uuid PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "fullname" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "Joke" (
  "id" uuid PRIMARY KEY,
  "title" varchar NOT NULL,
  "text" varchar NOT NULL,
  "explanation" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "Likes" (
  "created_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL,
  "joke_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "joke_id")
);

CREATE TABLE "Comment" (
  "id" uuid PRIMARY KEY,
  "text" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL,
  "joke_id" uuid NOT NULL
);

CREATE TABLE "Room" (
  "id" uuid PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "creator_user_id" uuid NOT NULL
);

CREATE TABLE "RoomMessage" (
  "id" uuid PRIMARY KEY,
  "text" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL,
  "room_id" uuid NOT NULL
);

CREATE TABLE "RoomMember" (
  "created_at" timestamp NOT NULL,
  "user_id" uuid NOT NULL,
  "room_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "room_id")
);

CREATE UNIQUE INDEX ON "Joke" ("title", "user_id");

ALTER TABLE "Joke" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Likes" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Likes" ADD FOREIGN KEY ("joke_id") REFERENCES "Joke" ("id");

ALTER TABLE "Comment" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Comment" ADD FOREIGN KEY ("joke_id") REFERENCES "Joke" ("id");

ALTER TABLE "Room" ADD FOREIGN KEY ("creator_user_id") REFERENCES "User" ("id");

ALTER TABLE "RoomMessage" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "RoomMessage" ADD FOREIGN KEY ("room_id") REFERENCES "Room" ("id");

ALTER TABLE "RoomMember" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "RoomMember" ADD FOREIGN KEY ("room_id") REFERENCES "Room" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "RoomMember";
DROP TABLE IF EXISTS "RoomMessage";
DROP TABLE IF EXISTS "Room";
DROP TABLE IF EXISTS "Comment";
DROP TABLE IF EXISTS "Like";
DROP TABLE IF EXISTS "Joke";
DROP TABLE IF EXISTS "User";
-- +goose StatementEnd
