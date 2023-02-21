DROP TABLE IF EXISTS `roles`;

CREATE TABLE IF NOT EXISTS `roles`(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    CONSTRAINT PK_Roles PRIMARY KEY (id)
);

INSERT INTO `roles` (name,created_at) values ('administrator',CURRENT_TIMESTAMP());
INSERT INTO `roles` (name,created_at) values ('project manager',CURRENT_TIMESTAMP());
INSERT INTO `roles` (name,created_at) values ('project member',CURRENT_TIMESTAMP());


