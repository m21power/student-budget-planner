CREATE TABLE users (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing UID
    name VARCHAR(100) NOT NULL,  -- User name
    email VARCHAR(100) UNIQUE NOT NULL,  -- User email (unique constraint)
    password VARCHAR(255) NOT NULL,  -- User password
    badges TEXT[]  -- Array of badges (list of strings)
);