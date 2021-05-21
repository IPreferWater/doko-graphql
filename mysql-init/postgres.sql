CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    title VARCHAR(30),
    txt TEXT,
    latitude FLOAT,
    longitude FLOAT
);

CREATE TABLE users(
   id SERIAL PRIMARY KEY,
  username varchar(64) UNIQUE,
  email varchar(64) UNIQUE,
  password varchar(255) NOT NULL
);

INSERT INTO users(username,email,password)
VALUES ('ipreferwater','random@email.com','password');