CREATE TABLE IF NOT EXISTS reviewed_actions (
  id INT(11) unsigned NOT NULL AUTO_INCREMENT,
  pullreq_id INT(11) unsigned NOT NULL,
  user_id INT(11) unsigned NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (pullreq_id) REFERENCES pullrequests(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
