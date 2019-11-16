CREATE TABLE IF NOT EXISTS repositories (
  id           INT(11)      unsigned NOT NULL AUTO_INCREMENT,
  organization VARCHAR(255)          NOT NULL,
  repository   VARCHAR(255)          NOT NULL,
  created_at   DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE fullname (organization, repository)
)
