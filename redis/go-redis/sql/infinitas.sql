CREATE DATABASE IF NOT EXISTS infinitas;
USE infinitas;

CREATE TABLE IF NOT EXISTS product (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  quantity INT NOT NULL
);