
ALTER TABLE `accounts` DROP INDEX `owner_currency_index`;
ALTER TABLE transferDB.accounts DROP FOREIGN KEY accounts_ibfk_1;
DROP TABLE IF EXISTS transferDB.users;