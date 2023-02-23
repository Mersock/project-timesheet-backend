DROP TABLE IF EXISTS `time_entries`;

CREATE TABLE IF NOT EXISTS `time_entries`(
    id int NOT NULL AUTO_INCREMENT,
    status_id int NOT NULL,
    work_type_id int NOT NULL,
    user_id int NOT NULL,
    start_time datetime NOT NULL,
    end_time datetime,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (status_id) REFERENCES statuses(id) ON DELETE CASCADE,
    FOREIGN KEY (work_type_id) REFERENCES worktypes(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);