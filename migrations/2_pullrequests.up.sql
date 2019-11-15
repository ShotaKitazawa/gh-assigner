CREATE TABLE IF NOT EXISTS pullrequests (
  id INT(11) unsigned NOT NULL,
  title VARCHAR(255),
  user_id INT(11) unsigned NOT NULL,
  opened_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  merged_at DATETIME DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
)
