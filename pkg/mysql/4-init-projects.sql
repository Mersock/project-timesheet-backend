DROP TABLE IF EXISTS `projects`;

CREATE TABLE IF NOT EXISTS `projects`(
    id int NOT NULL AUTO_INCREMENT,
    code varchar(50) UNIQUE NOT NULL,
    name varchar(255)  NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    PRIMARY KEY (id)
);