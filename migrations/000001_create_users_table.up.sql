
-- Buat enum types
CREATE TYPE preference_type AS ENUM ('CARDIO', 'WEIGHT');
CREATE TYPE weight_unit_type AS ENUM ('KG', 'LBS');
CREATE TYPE height_unit_type AS ENUM ('CM', 'INCH');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(60),
    password VARCHAR(32) NOT NULL,

    preference preference_type,
    weightUnit weight_unit_type, 
    heightUnit height_unit_type,
    weight INT default(1), 
    height INT default(1),
    imageUri VARCHAR(255)
);