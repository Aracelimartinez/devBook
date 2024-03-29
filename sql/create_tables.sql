DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

-- users
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  nick VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(150) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- followers
CREATE TABLE
 IF NOT EXISTS followers (
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  follower_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, follower_id)
);

-- posts
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  content VARCHAR(300) NOT NULL,
  author_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  likes INT DEFAULT 0,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
