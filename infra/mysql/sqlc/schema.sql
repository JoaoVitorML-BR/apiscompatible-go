CREATE DATABASE IF NOT EXISTS apiscompatiblego;
USE apiscompatiblego;
CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);

-- OLD SCHEMA THAT WORKED 100% BEFORE THE CHANGES TO LOGIN.

-- CREATE DATABASE IF NOT EXISTS apiscompatiblego;
-- USE apiscompatiblego;
-- CREATE TABLE users (
--   id INT NOT NULL AUTO_INCREMENT,
--   name VARCHAR(255) NOT NULL,
--   password BLOB NOT NULL,
--   created_at DATETIME NOT NULL,
--   updated_at DATETIME NOT NULL,
--   PRIMARY KEY (id)
-- );

