DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS tests;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS enrollments;

CREATE TABLE students (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL
);

CREATE TABLE tests (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE questions (
    id VARCHAR(32) PRIMARY KEY,
    test_id VARCHAR(32) NOT NULL,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(255) NOT NULL,

    FOREIGN KEY (test_id) REFERENCES tests(id)
);

CREATE TABLE enrollments (
    student_id VARCHAR(32) NOT NULL,
    test_id VARCHAR(32) NOT NULL,

    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (test_id) REFERENCES tests(id)
)