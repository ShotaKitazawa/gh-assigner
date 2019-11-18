CREATE TABLE IF NOT EXISTS actions (
  id              INT(11)  unsigned NOT NULL AUTO_INCREMENT,
  pullreq_id      INT(11)  unsigned NOT NULL,
  request_user_id INT(11)  unsigned NOT NULL,
  requested_at    DATETIME          NOT NULL DEFAULT CURRENT_TIMESTAMP,
  review_user_id  INT(11)  unsigned          DEFAULT NULL,
  reviewed_at     DATETIME                   DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (pullreq_id) REFERENCES pullrequests(id),
  FOREIGN KEY (request_user_id) REFERENCES users(id),
  FOREIGN KEY (review_user_id) REFERENCES users(id)
);
