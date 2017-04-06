USE spary;
CREATE TABLE IF NOT EXISTS spa (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  address VARCHAR(255),
  url VARCHAR(255),
  tel VARCHAR(255),
  effect TEXT,
  fee INTEGER,
  image VARCHAR(255),
  equipment VARCHAR(255),
  workday VARCHAR(255),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255),
  address VARCHAR(255),
  PRIMARY KEY (id)
);

create table if not exists comments (
  id mediumint not null auto_increment,
  spa_id mediumint not null,
  user_id mediumint not null,
  comment text,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS favorite (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  user_id INTEGER,
  spa_id INTEGER,
  PRIMARY KEY (id)
);

USE spary_test;
CREATE TABLE IF NOT EXISTS spa (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  address VARCHAR(255),
  url VARCHAR(255),
  tel VARCHAR(255),
  effect TEXT,
  fee INTEGER,
  image VARCHAR(255),
  equipment VARCHAR(255),
  workday VARCHAR(255),
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255),
  email VARCHAR(255),
  address VARCHAR(255),
  PRIMARY KEY (id)
);

create table if not exists comments (
  id mediumint not null auto_increment,
  spa_id mediumint not null,
  user_id mediumint not null,
  comment text,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS favorite (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  user_id INTEGER,
  spa_id INTEGER,
  PRIMARY KEY (id)
);
