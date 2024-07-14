CREATE TABLE IF NOT EXISTS Users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Movies (
    id CHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    director VARCHAR(255) NOT NULL,
    year INT NOT NULL
);

INSERT INTO Movies (id, title, director, year) VALUES
('13d7bd2f-732b-465a-bcbe-4b0bc58c3fad', 'Inception', 'Christopher Nolan', 2010),
('adc0c387-cd27-47a7-a9a3-1fcf7a6a8fab', 'The Matrix', 'Lana Wachowski, Lilly Wachowski', 1999),
('d704f9e8-ed68-4655-b471-feaa3aad955a', 'Interstellar', 'Christopher Nolan', 2014);
