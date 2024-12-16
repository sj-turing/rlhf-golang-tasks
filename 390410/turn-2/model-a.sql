SELECT * FROM users WHERE age > 30;

SELECT id, username, email FROM users WHERE age > 30;

CREATE INDEX idx_age ON users(age);

EXPLAIN SELECT id, username, email FROM users WHERE age > 30;
