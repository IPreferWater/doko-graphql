GRANT ALL PRIVILEGES ON doko  TO 'user';

CREATE TABLE posts(
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(30),
    txt TEXT,
    latitude FLOAT,
    longitude FLOAT
);

INSERT INTO posts(title,txt,latitude,longitude)
VALUES ("1","blablabla",50.597186,3.112793),
("2","il pleut sa mere",50.797791,4.769001),
("3","deja en Allemagne",50.930222,6.242893),
("4","en foret c'est cool",51.034399,8.336635);