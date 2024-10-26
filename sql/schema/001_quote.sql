-- +goose Up
CREATE TABLE Quotes (
    Id SERIAL PRIMARY KEY,
    Quote TEXT NOT NULL,
    Author TEXT NOT NULL,
    Book TEXT NOT NULL,
    Post_Date TIMESTAMP
);
-- +goose Down
DROP TABLE Quotes;