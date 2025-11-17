-- Active: 1762854347508@@127.0.0.1@3306@SHORTEN_URL
CREATE DATABASE shorten_url;

USE shorten_url;

CREATE TABLE url_encode (
    `shorten_id` VARCHAR(50) NOT NULL UNIQUE COLLATE utf8mb4_bin,
    `shorten_number` INT UNSIGNED NOT NULL UNIQUE,
    `original_url` VARCHAR(255) NOT NULL COLLATE utf8mb4_bin,
    PRIMARY KEY (`shorten_id`)
);
