
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名',
  `amount` mediumint unsigned NOT NULL DEFAULT 0 COMMENT '値段',
  `created_at` datetime(6) NOT NULL,
  `updated_at` datetime(6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE items;
