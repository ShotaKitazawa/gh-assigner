CREATE TABLE IF NOT EXISTS repositories (
  id           INT(11)      unsigned NOT NULL AUTO_INCREMENT,
  domain       VARCHAR(255)          NOT NULL,
  organization VARCHAR(255)          NOT NULL,
  repository   VARCHAR(255)          NOT NULL,
  created_at   DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE fullname (domain, organization, repository)
)
