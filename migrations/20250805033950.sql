-- Create "language_detections" table
CREATE TABLE `language_detections` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `query` varchar(255) NOT NULL,
  `detected_languages` json NOT NULL,
  `duration` bigint NOT NULL,
  `model_name` varchar(50) NOT NULL,
  `input_token` bigint NOT NULL,
  `output_token` bigint NOT NULL,
  `cached_token` bigint NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  `cost` double NOT NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "translations" table
CREATE TABLE `translations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `content` varchar(255) NOT NULL,
  `choices` json NOT NULL,
  `external_id` varchar(255) NOT NULL,
  `language_pair` varchar(255) NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `external_id` (`external_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `full_name` varchar(255) NULL,
  `is_active` bool NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `role` varchar(255) NULL DEFAULT "user",
  `external_id` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email` (`email`),
  UNIQUE INDEX `external_id` (`external_id`),
  UNIQUE INDEX `username` (`username`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "transcription_stats" table
CREATE TABLE `transcription_stats` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `duration` bigint NOT NULL,
  `input_token` bigint NOT NULL,
  `output_token` bigint NOT NULL,
  `model_name` varchar(50) NOT NULL,
  `cost` double NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  `translation_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `transcription_stats_translations_transcription_stats` (`translation_id`),
  CONSTRAINT `transcription_stats_translations_transcription_stats` FOREIGN KEY (`translation_id`) REFERENCES `translations` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
