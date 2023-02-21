DROP TABLE IF EXISTS `duties`;

CREATE TABLE IF NOT EXISTS `duties`(
    project_id int NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY (project_id, user_id),
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);