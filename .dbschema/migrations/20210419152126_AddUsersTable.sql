
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'メールアドレス',
  `password_digest` varchar(255) NOT NULL DEFAULT '' COMMENT 'パスワードダイジェスト',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `index_users_on_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;
