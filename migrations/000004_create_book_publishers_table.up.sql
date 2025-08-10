CREATE TABLE book_publishers (
    publisher_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_date TIMESTAMP WITH TIME ZONE NULL
);

ALTER TABLE books
ADD COLUMN publisher_id INTEGER REFERENCES book_publishers(publisher_id);

INSERT INTO book_publishers (name, description) VALUES
('Эксмо', 'Крупнейшее издательство России, выпускающее художественную и научную литературу.'),
('АСТ', 'Одно из ведущих российских издательств, специализирующееся на классике и современной прозе.'),
('Азбука', 'Издательство, известное качественными изданиями русской и зарубежной литературы.'),
('Просвещение', 'Издательство, выпускающее учебную и образовательную литературу.'),
('Рипол Классик', 'Издательство, специализирующееся на классической литературе и биографиях.');

UPDATE books SET publisher_id = 1 WHERE title = 'Война и мир';
UPDATE books SET publisher_id = 2 WHERE title = 'Преступление и наказание';
UPDATE books SET publisher_id = 3 WHERE title = 'Мастер и Маргарита';
UPDATE books SET publisher_id = 4 WHERE title = 'Евгений Онегин';
UPDATE books SET publisher_id = 5 WHERE title = 'Отцы и дети';
UPDATE books SET publisher_id = 1 WHERE title = 'Анна Каренина';
UPDATE books SET publisher_id = 2 WHERE title = 'Мёртвые души';
UPDATE books SET publisher_id = 3 WHERE title = 'Доктор Живаго';
UPDATE books SET publisher_id = 4 WHERE title = 'Герой нашего времени';
UPDATE books SET publisher_id = 5 WHERE title = 'Двенадцать стульев';
