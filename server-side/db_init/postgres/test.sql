INSERT INTO users VALUES ('prof1', 'Professor1', 'Random', '9beb8428008fba903a067a9ce23828f5c4196e74cabaf7131233bfee3aa3a991');
INSERT INTO users VALUES ('prof2', 'Professor2', 'Random', 'dc7a2db63253c482d862cd4bdfc81b960fbf54aa58eee1f1c31a3e7ce5e00d38');
INSERT INTO users VALUES ('asistent1', 'Asistent1', 'Random', '37f514205db02ad131172e6db7534a13712e9a8f28c59a388289270d66b9b786');
INSERT INTO users VALUES ('asistent2', 'Asistent2', 'Random', '2789116cfc062aca0eae664c6d037390e45b146aa7c287cf056b9d86bb6e08bc');
INSERT INTO users VALUES ('stud1', 'Student1', 'Random', '7ee252e54e0742935727c722fd0fdae554fe01ddbbf961f315f9f1fa89af2eb5');
INSERT INTO users VALUES ('stud2', 'Student2', 'Random', '93f09f31a7bd22079a4078616be52012aa5c5a8a1e8be231d3ab45187ae1fbd9');
INSERT INTO users VALUES ('stud3', 'Student3', 'Random', 'c2a2c9903026fad6193d1f2476ee7e082ec76c5b2abc78670905225fd41f3b1d');
INSERT INTO users VALUES ('stud4', 'Student4', 'Random', '8014109378959842b7ed494a2bd25ad24872976976be8db0bd3cde676256c50f');
INSERT INTO users VALUES ('stud5', 'Student5', 'Random', '756986f2670ee60b14f37c58e240cf26fa77d3a3f8cff2db357dd406f73f9f91');

INSERT INTO series (school_year, name) VALUES (2020, 'CB');
INSERT INTO series (school_year, name) VALUES (2020, 'CC');

INSERT INTO groups (group_no, series_id) VALUES (315, 1);
INSERT INTO groups (group_no, series_id) VALUES (335, 2);

INSERT INTO courses (name, series_id, abbreviation, professor) VALUES ('Programarea Calculatoarelor', 1, 'PC', 'prof1');
INSERT INTO courses (name, series_id, abbreviation, professor) VALUES ('Algoritmi Paraleli si Distribuiti', 1, 'APD', 'prof2');

INSERT INTO students VALUES ('stud1', 1);
INSERT INTO students VALUES ('stud2', 1);
INSERT INTO students VALUES ('stud3', 1);
INSERT INTO students VALUES ('stud4', 2);
INSERT INTO students VALUES ('stud5', 2);

INSERT INTO students_courses VALUES ('stud1', 1);
INSERT INTO students_courses VALUES ('stud2', 1);
INSERT INTO students_courses VALUES ('stud3', 1);
INSERT INTO students_courses VALUES ('stud4', 2);
INSERT INTO students_courses VALUES ('stud5', 2);

INSERT INTO assistants_courses VALUES ('asistent1', 2);
INSERT INTO assistants_courses VALUES ('asistent2', 1);
INSERT INTO assistants_courses VALUES ('stud4', 1);
INSERT INTO assistants_courses VALUES ('prof2', 2);