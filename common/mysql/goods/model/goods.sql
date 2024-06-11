CREATE TABLE goods
(
  id bigint NOT NULL AUTO_INCREMENT ,
  name varchar(50)  NOT NULL,
  price bigint NOT NULL ,
  stock int NOT NULL,
  PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE = utf8_general_ci ;