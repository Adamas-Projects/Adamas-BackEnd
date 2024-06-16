DROP DATABASE IF EXISTS adamas_db;

CREATE DATABASE adamas_db;

USE adamas_db;


CREATE TABLE INSTITUTION_USER(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(64) NOT NULL,
    cnpj bigint NOT NULL
);

CREATE TABLE COMMON_USER(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    institution_id int NULL,
    password varchar(64) NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES INSTITUTION_USER(id)
);


CREATE TABLE REPOSITORY(
    id int auto_increment NOT NULL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    content varchar(255) NOT NULL
);
CREATE TABLE OWNERS_REPOSITORY(
    repository_id int NOT NULL REFERENCES REPOSITORY(id),
    owner_id int NOT NULL REFERENCES COMMON_USER(id),
    PRIMARY KEY(repository_id, owner_id)
);

CREATE TABLE EVENT(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL,
    address varchar(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    description varchar(255) NOT NULL
);
CREATE TABLE OWNER_EVENT(
    event_id int NOT NULL REFERENCES EVENT(id),
    owner_id int NOT NULL REFERENCES INSTITUTION_USER(id),
    PRIMARY KEY(event_id, owner_id)
);
CREATE TABLE ROOM_IN_EVENT(
    id int auto_increment NOT NULL PRIMARY KEY,
    event_id int NOT NULL,
    name varchar(50) NOT NULL,
    quantity_repos int NOT NULL,
    FOREIGN KEY (event_id) REFERENCES EVENT(id) 
);
CREATE TABLE REPOSITORY_IN_ROOM(
    repository_id int NOT NULL,
    room_id int NOT NULL,
    PRIMARY KEY(repository_id, room_id)
);
CREATE TABLE CATEGORY_REPO(
    category_id int NOT NULL,
    repository_id int NOT NULL,
    PRIMARY KEY(category_id,repository_id)
);

DELIMITER $$
CREATE TRIGGER limit_category_count
BEFORE INSERT ON CATEGORY_REPO
FOR EACH ROW
BEGIN
    DECLARE category_count INT;
    SELECT COUNT(*) INTO category_count 
    FROM CATEGORY_REPO 
    WHERE repository_id = NEW.repository_id;
    IF category_count >= 3 THEN
        SIGNAL SQLSTATE '45000' 
        SET MESSAGE_TEXT = 'Limite de 3 categorias por repositório atingido';
    END IF;
END$$
DELIMITER ;

CREATE TABLE CATEGORY(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(200) NOT NULL
); 

INSERT INTO CATEGORY(name) values 
("Saúde"),
("Agricultura"),
("Ferramenta"),
("Música"),
("TI"),
("Marketing"),
("Mecânica");

                                                                                                                     