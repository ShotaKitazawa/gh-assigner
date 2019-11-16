CREATE TABLE IF NOT EXISTS pullrequests (
  id            INT(11)      unsigned NOT NULL AUTO_INCREMENT,
  user_id       INT(11)      unsigned NOT NULL,
  repository_id INT(11)      unsigned NOT NULL,
  issue_id      INT(11)      unsigned NOT NULL,
  title         VARCHAR(255)          NOT NULL,
  state         VARCHAR(255)          NOT NULL,
  created_at    DATETIME              NOT NULL DEFAULT CURRENT_TIMESTAMP,
  closed_at     DATETIME                       DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (repository_id) REFERENCES repositories(id)
)
