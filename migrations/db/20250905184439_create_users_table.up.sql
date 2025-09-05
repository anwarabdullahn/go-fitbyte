-- Create enum types for user preferences
CREATE TYPE preference_type AS ENUM ('CARDIO', 'WEIGHT');
CREATE TYPE weight_unit_type AS ENUM ('KG', 'LBS');
CREATE TYPE height_unit_type AS ENUM ('CM', 'INCH');

-- Create users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    preference preference_type,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(60),
    weightunit weight_unit_type,
    heightunit height_unit_type,
    weight INTEGER,
    height INTEGER,
    imageuri VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);