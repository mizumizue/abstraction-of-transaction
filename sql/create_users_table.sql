CREATE TABLE IF NOT EXISTS users(
   user_id BIGINT PRIMARY KEY,
   user_name VARCHAR (50) NOT NULL,
   email VARCHAR (300) NOT NULL,
   telephone VARCHAR(20) DEFAULT NULL
);