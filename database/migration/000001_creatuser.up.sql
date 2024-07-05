CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id),
    full_name VARCHAR(255),
    bio TEXT,
    expertise VARCHAR(255),
    location VARCHAR(255),
    avatar_url TEXT
);
