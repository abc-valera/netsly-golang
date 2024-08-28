-- +goose Up
-- +goose StatementBegin

CREATE TABLE `users` (
  `id` text NOT NULL, 
  `username` text NOT NULL, 
  `email` text NOT NULL, 
  `hashed_password` text NOT NULL, 
  `fullname` text NOT NULL, 
  `status` text NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  PRIMARY KEY (`id`), 
  CONSTRAINT `uni_users_username` UNIQUE (`username`), 
  CONSTRAINT `uni_users_email` UNIQUE (`email`)
);

CREATE TABLE `jokes` (
  `id` text NOT NULL, 
  `title` text NOT NULL, 
  `text` text NOT NULL, 
  `explanation` text NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `user_id` text NOT NULL, 
  PRIMARY KEY (`id`), 
  CONSTRAINT `fk_users_jokes` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE UNIQUE INDEX `unique_title_per_user` ON `jokes`(`title`, `user_id`);

CREATE TABLE `likes` (
  `created_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `user_id` text NOT NULL, 
  `joke_id` text NOT NULL, 
  PRIMARY KEY (`user_id`, `joke_id`), 
  CONSTRAINT `fk_jokes_likes` FOREIGN KEY (`joke_id`) REFERENCES `jokes`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `fk_users_likes` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE `comments` (
  `id` text NOT NULL, 
  `text` text NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `user_id` text NOT NULL, 
  `joke_id` text NOT NULL, 
  PRIMARY KEY (`id`), 
  CONSTRAINT `fk_jokes_comments` FOREIGN KEY (`joke_id`) REFERENCES `jokes`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `fk_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE `rooms` (
  `id` text NOT NULL, 
  `name` text NOT NULL, 
  `description` text NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `creator_user_id` text NOT NULL, 
  PRIMARY KEY (`id`), 
  CONSTRAINT `fk_users_created_rooms` FOREIGN KEY (`creator_user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `uni_rooms_name` UNIQUE (`name`)
);

CREATE TABLE `room_members` (
  `created_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `user_id` text NOT NULL, 
  `room_id` text NOT NULL, 
  PRIMARY KEY (`user_id`, `room_id`), 
  CONSTRAINT `fk_users_room_members` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `fk_rooms_members` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE
);

CREATE TABLE `file_infos` (
  `id` text NOT NULL, 
  `name` text NOT NULL, 
  `type` integer NOT NULL, 
  `size` integer NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS "room_messages" (
  `id` text NOT NULL, 
  `text` text NOT NULL, 
  `created_at` datetime NOT NULL, 
  `updated_at` datetime NOT NULL, 
  `deleted_at` datetime NOT NULL, 
  `user_id` text NOT NULL, 
  `room_id` text NOT NULL, 
  PRIMARY KEY (`id`), 
  CONSTRAINT `fk_rooms_messages` FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `fk_users_room_messages` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE, 
  CONSTRAINT `fk_room_members_messages` FOREIGN KEY (`user_id`, `room_id`) REFERENCES `room_members`(`user_id`, `room_id`)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "FileInfo";
DROP TABLE IF EXISTS "RoomMember";
DROP TABLE IF EXISTS "RoomMessage";
DROP TABLE IF EXISTS "Room";
DROP TABLE IF EXISTS "Comment";
DROP TABLE IF EXISTS "Like";
DROP TABLE IF EXISTS "Joke";
DROP TABLE IF EXISTS "User";

-- +goose StatementEnd
