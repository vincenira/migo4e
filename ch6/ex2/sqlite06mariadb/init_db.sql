-- MariaDB script

-- Drop the database if it exists
-- DROP DATABASE IF EXISTS social_media;

-- Create the database
-- CREATE DATABASE social_media;

-- Use the database
-- USE social_media;

-- Table 'UserData'
-- DROP TABLE IF EXISTS UserData;

-- CREATE TABLE UserData (
--     id CHAR(36) NOT NULL PRIMARY KEY,
--     data VARCHAR(160) NOT NULL
-- );

-- Drop the database if it exists
DROP DATABASE IF EXISTS social_media;

-- Create the database
CREATE DATABASE social_media;

-- Use the database
USE social_media;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Userdata;

CREATE TABLE Users (
    ID SERIAL,
    Username VARCHAR(100) PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE Userdata (
    UserID Int NOT NULL,
    Name VARCHAR(100),
    Surname VARCHAR(100),
    Description VARCHAR(200)
);
