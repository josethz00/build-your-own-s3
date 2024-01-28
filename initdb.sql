CREATE TABLE IF NOT EXISTS `buckets` (
  `id` bigint unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `public` boolean NOT NULL DEFAULT false,
  `description` text,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
);

CREATE TABLE IF NOT EXISTS `files` (
  `id` bigint unsigned NOT NULL,
  `bucket_id` bigint unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `size` bigint unsigned NOT NULL,
  `mime_type` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `bucket_id` (`bucket_id`),
  CONSTRAINT `files_ibfk_1` FOREIGN KEY (`bucket_id`) REFERENCES `buckets` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);