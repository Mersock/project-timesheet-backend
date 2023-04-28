DROP TABLE IF EXISTS `work_types`;

CREATE TABLE IF NOT EXISTS `work_types`(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255)  NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    project_id int NOT NULL,
    CONSTRAINT PK_Worktypes PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
);