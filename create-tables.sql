create schema merchandise;

CREATE TABLE `merchandise`.`merch` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(450) NOT NULL,
  `type` VARCHAR(45) NULL,
  `price` DECIMAL(5,2) NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE `merchandise`.`stock` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `merch_fkid` INT NOT NULL,
  `size` VARCHAR(45) NULL,
  `quantity` INT NULL,
  PRIMARY KEY (`id`));

ALTER TABLE `merchandise`.`merch` 
ADD COLUMN `color` VARCHAR(100) NULL AFTER `price`;
