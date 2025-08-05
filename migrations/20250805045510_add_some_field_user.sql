-- Modify "users" table
ALTER TABLE `users` MODIFY COLUMN `provider` varchar(255) NULL DEFAULT "local", ADD COLUMN `provider_id` varchar(255) NULL;
