CREATE TABLE IF NOT EXISTS request_actions (
  id INT(11) unsigned NOT NULL AUTO_INCREMENT,
  pullreq_id INT(11) unsigned NOT NULL,
  user_id INT(11) unsigned NOT NULL,
  time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (pullreq_id) REFERENCES pullrequests(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
