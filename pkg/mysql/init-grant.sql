CREATE USER IF NOT EXISTS 'project_timesheet_admin'@'localhost' IDENTIFIED BY 'P@ssw0rd';


GRANT ALL PRIVILEGES ON project_timesheet.* TO 'project_timesheet_admin'@'localhost' WITH GRANT OPTION;

