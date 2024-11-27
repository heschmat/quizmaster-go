sqlite3 quiz.db <<EOF
CREATE TABLE questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    prompt TEXT NOT NULL,
    option1 TEXT NOT NULL,
    option2 TEXT NOT NULL,
    option3 TEXT NOT NULL,
    option4 TEXT NOT NULL,
    answer INTEGER NOT NULL
);

INSERT INTO questions (prompt, option1, option2, option3, option4, answer) VALUES
('What is the capital of France?', 'Berlin', 'Madrid', 'Paris', 'Rome', 3),
('Which programming language is known as "Go"?', 'Python', 'JavaScript', 'Ruby', 'Golang', 4),
('What is 5 + 3?', '6', '7', '8', '9', 3),
('What is the largest planet in our Solar System?', 'Earth', 'Mars', 'Jupiter', 'Saturn', 3),
('What is the speed of light?', '3x10^8 m/s', '5x10^8 m/s', '1x10^6 m/s', '2x10^8 m/s', 1),
('Who wrote "To Kill a Mockingbird"?', 'Harper Lee', 'Mark Twain', 'George Orwell', 'Jane Austen', 1),
('What is the square root of 64?', '6', '7', '8', '9', 3),
('Who painted the Mona Lisa?', 'Van Gogh', 'Leonardo da Vinci', 'Picasso', 'Rembrandt', 2),
('What is the chemical symbol for water?', 'HO', 'H2O', 'O2', 'H2', 2),
('Which country hosted the 2016 Summer Olympics?', 'China', 'USA', 'Brazil', 'Russia', 3),
('What is the capital of Japan?', 'Beijing', 'Seoul', 'Tokyo', 'Bangkok', 3),
('What is 12 * 12?', '124', '144', '154', '164', 2),
('What is the freezing point of water in Celsius?', '0', '32', '100', '-1', 1),
('What year did the Titanic sink?', '1910', '1912', '1914', '1916', 2),
('What is the tallest mountain in the world?', 'K2', 'Everest', 'Kilimanjaro', 'Makalu', 2),
('Which gas is most abundant in the Earth’s atmosphere?', 'Oxygen', 'Nitrogen', 'Carbon Dioxide', 'Hydrogen', 2),
('Who discovered gravity?', 'Newton', 'Einstein', 'Galileo', 'Tesla', 1),
('What is the capital of Italy?', 'Milan', 'Venice', 'Rome', 'Naples', 3),
('Who is known as the Father of Computers?', 'Charles Babbage', 'Alan Turing', 'Bill Gates', 'Steve Jobs', 1),
('Which planet is known as the Red Planet?', 'Earth', 'Mars', 'Venus', 'Jupiter', 2);

EOF
