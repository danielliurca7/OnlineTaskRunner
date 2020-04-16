CREATE TABLE IF NOT EXISTS users (
    user_name varchar PRIMARY KEY,
    first_name varchar NOT NULL,
    last_name varchar NOT NULL,
    password_hash char(64) NOT NULL
);

CREATE TABLE IF NOT EXISTS series (
    id serial PRIMARY KEY,
    school_year integer NOT NULL,
    name varchar NOT NULL
);

CREATE TABLE IF NOT EXISTS groups (
    id serial PRIMARY KEY,
    group_no integer NOT NULL,
    series_id integer REFERENCES series(id)
);

CREATE TABLE IF NOT EXISTS students (
    name varchar REFERENCES users(user_name) PRIMARY KEY,
    group_id integer REFERENCES groups(id)
);

CREATE TABLE IF NOT EXISTS courses (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    series_id integer REFERENCES series(id),
    abbreviation varchar(5) NOT NULL,    
    professor varchar REFERENCES users(user_name)

);

CREATE TABLE IF NOT EXISTS students_courses (
    student_name varchar REFERENCES students(name),
    course_id integer REFERENCES courses(id),
    PRIMARY KEY (student_name, course_id)
);

CREATE TABLE IF NOT EXISTS assistants_courses (
    assistant_name varchar REFERENCES users(user_name),
    course_id integer REFERENCES courses(id),
    PRIMARY KEY (assistant_name, course_id)
);
