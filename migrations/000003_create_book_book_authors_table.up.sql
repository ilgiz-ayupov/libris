CREATE TABLE book_book_authors (
    book_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id),
    FOREIGN KEY (author_id) REFERENCES book_authors(author_id)
);

INSERT INTO book_book_authors (book_id, author_id) VALUES
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 4),
    (3, 5),
    (3, 6),
    (4, 7),
    (4, 8),
    (5, 9),
    (5, 10),
    (6, 1),
    (6, 3),
    (7, 2),
    (7, 4),
    (8, 5),
    (8, 7),
    (9, 6),
    (9, 8),
    (10, 9),
    (10, 10);