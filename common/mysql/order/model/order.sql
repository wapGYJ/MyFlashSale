CREATE TABLE theorder  (
    id int(0) NOT NULL AUTO_INCREMENT ,
    userid int(0) NOT NULL ,
    goodsid int(0) NOT NULL ,
    content varchar(36)  NOT NULL ,
    PRIMARY KEY (id),
    UNIQUE name_index (goodsid)
) ENGINE = InnoDB COLLATE = utf8_general_ci