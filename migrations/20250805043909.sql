-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `avatar_url` varchar(255) NULL AFTER `external_id`, ADD COLUMN `provider` varchar(255) NOT NULL DEFAULT "local";
