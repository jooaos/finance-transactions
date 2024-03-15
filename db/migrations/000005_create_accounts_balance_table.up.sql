CREATE TABLE IF NOT EXISTS `accounts_balance` (
    `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `account_id` INT NOT NULL UNIQUE,
    `balance` DECIMAL(11,2) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    CONSTRAINT `fk_accounts_balance_account_id` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`)
);