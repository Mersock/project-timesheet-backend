DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users`(
    id int NOT NULL AUTO_INCREMENT,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255)  NOT NULL,
    firstname varchar(255) NOT NULL,
    lastname varchar(255) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    role_id int NOT NULL,
    CONSTRAINT PK_Users PRIMARY KEY (id),
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

INSERT INTO project_timesheet.users (email, password, firstname, lastname, created_at, role_id) VALUES('admin@admin.com', '1234', 'admin', 'admin', CURRENT_TIMESTAMP(), 1);
