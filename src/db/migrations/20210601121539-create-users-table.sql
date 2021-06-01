-- +migrate Up
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(20) DEFAULT "",
  `email` VARCHAR(50) NOT NULL,
  `password` VARCHAR(60) NOT NULL
  `hobby` VARCHAR(60) DEFAULT "",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE IF EXISTS `users`;
