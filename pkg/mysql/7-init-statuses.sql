DROP TABLE IF EXISTS `statuses`;

CREATE TABLE IF NOT EXISTS `statuses`(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    CONSTRAINT PK_Roles PRIMARY KEY (id)
);

INSERT INTO `statuses` (name,created_at) values ('to do',CURRENT_TIMESTAMP());
INSERT INTO `statuses` (name,created_at) values ('in progress',CURRENT_TIMESTAMP());
INSERT INTO `statuses` (name,created_at) values ('blocker',CURRENT_TIMESTAMP());
INSERT INTO `statuses` (name,created_at) values ('done',CURRENT_TIMESTAMP());


