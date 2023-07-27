CREATE TABLE IF NOT EXISTS `accounts` (
  `id` BIGINT(20) UNSIGNED NOT NULL,
  `owner` VARCHAR(255) NOT NULL,
  `balance` BIGINT(20) UNSIGNED NOT NULL,
  `currency` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX `owner` (`owner` ASC) VISIBLE,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;