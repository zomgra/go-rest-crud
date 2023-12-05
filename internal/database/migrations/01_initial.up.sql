CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_id INT,
    isnbn VARCHAR(255) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors(id)
);
