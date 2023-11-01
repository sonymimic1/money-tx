-- -----------------------------------------------------
-- Table `transferDB`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `username` VARCHAR(255) NOT NULL,
  `hashed_password` VARCHAR(255) NOT NULL,
  `full_name` VARCHAR(255)  NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `password_change_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`))
ENGINE = InnoDB;
ALTER TABLE `accounts`
ADD FOREIGN KEY ( `owner` ) REFERENCES `users` (`username`);

CREATE UNIQUE INDEX `owner_currency_index`
ON `accounts` (`owner`,`currency`);