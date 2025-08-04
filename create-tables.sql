create schema merchandise;

CREATE TABLE `merchendise`.`merch` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(450) NOT NULL,
  `type` VARCHAR(45) NULL,
  `price` DECIMAL(5,2) NULL,
  PRIMARY KEY (`id`));

  CREATE TABLE `merchendise`.`stock` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `merch_fkid` INT NOT NULL,
  `size` VARCHAR(45) NULL,
  `quantity` INT NULL,
  PRIMARY KEY (`id`));
