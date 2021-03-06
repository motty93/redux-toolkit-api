-- +migrate Up
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` INT(11) unsigned NOT NULL,
  `title` VARCHAR(50) DEFAULT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP NULL DEFAULT NULL,
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS `tasks`;
