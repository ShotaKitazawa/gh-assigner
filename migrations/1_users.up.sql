CREATE TABLE IF NOT EXISTS users (
  id INT(11) unsigned NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) UNIQUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
)
