DROP TABLE posts CASCADE IF EXISTS; -- CASCADEで関係を持っているテーブルのレコード？も連鎖して削除させる？
DROP TABLE comments IF EXISTS;
CREATE TABLE posts (
  id serial primary key,
  content TEXT,
  author VARCHAR(255)
);
CREATE TABLE comments (
  id serial primary key,
  content TEXT,
  author VARCHAR(255),
  post_id INTEGER references posts(id)
)