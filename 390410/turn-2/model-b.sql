CREATE INDEX idx_posts_user_id ON Posts (user_id);
CREATE INDEX idx_comments_post_id ON Comments (post_id);

EXPLAIN ANALYZE SELECT * FROM Posts JOIN Comments ON Posts.id = Comments.post_id WHERE Posts.user_id = 1;

SELECT * FROM Posts ORDER BY created_at DESC LIMIT 10 OFFSET 0;
