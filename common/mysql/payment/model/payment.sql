CREATE TABLE payment  (
                           id int(0) NOT NULL AUTO_INCREMENT ,
                           userid int(0) NOT NULL ,
                           goodsid int(0) NOT NULL ,
                           status tinyint default 0,
                           PRIMARY KEY (id),
                           UNIQUE name_index (goodsid)
) ENGINE = InnoDB COLLATE = utf8_general_ci