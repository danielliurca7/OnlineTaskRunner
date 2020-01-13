INSERT INTO Users VALUES ('prof1', 'Random', 'Professor1', '9beb8428008fba903a067a9ce23828f5c4196e74cabaf7131233bfee3aa3a991');
INSERT INTO Professors VALUES ('prof1');

INSERT INTO Users VALUES ('prof2', 'Random', 'Professor2', 'dc7a2db63253c482d862cd4bdfc81b960fbf54aa58eee1f1c31a3e7ce5e00d38');
INSERT INTO Professors VALUES ('prof2');

INSERT INTO Series (SchoolYear, SeriesName) VALUES (2020, 'CB');
INSERT INTO Series (SchoolYear, SeriesName) VALUES (2018, 'CB');

INSERT INTO Classes (GroupNumber, SeriesID) VALUES (315, 1);
INSERT INTO Classes (GroupNumber, SeriesID) VALUES (335, 1);

INSERT INTO Courses (CourseName, SeriesID, Abbreviation, ProfessorName)
VALUES ('Programarea Calculatoarelor', 1, 'PC', 'prof1');
INSERT INTO Courses (CourseName, SeriesID, Abbreviation, ProfessorName)
VALUES ('Algoritmi Paraleli si Distribuiti', 1, 'APD', 'prof2');

INSERT INTO Users VALUES ('stud1', 'Random', 'Student1', '7ee252e54e0742935727c722fd0fdae554fe01ddbbf961f315f9f1fa89af2eb5');
INSERT INTO Students VALUES ('stud1', 1);

INSERT INTO Users VALUES ('stud2', 'Random', 'Student2', '93f09f31a7bd22079a4078616be52012aa5c5a8a1e8be231d3ab45187ae1fbd9');
INSERT INTO Students VALUES ('stud2', 1);

INSERT INTO Users VALUES ('stud3', 'Random', 'Student3', 'c2a2c9903026fad6193d1f2476ee7e082ec76c5b2abc78670905225fd41f3b1d');
INSERT INTO Students VALUES ('stud3', 1);

INSERT INTO Users VALUES ('stud4', 'Random', 'Student4', '8014109378959842b7ed494a2bd25ad24872976976be8db0bd3cde676256c50f');
INSERT INTO Students VALUES ('stud4', 2);

INSERT INTO Users VALUES ('stud5', 'Random', 'Student5', '756986f2670ee60b14f37c58e240cf26fa77d3a3f8cff2db357dd406f73f9f91');
INSERT INTO Students VALUES ('stud5', 2);

INSERT INTO Assistents VALUES ('stud4');

INSERT INTO Assistents VALUES ('prof2');

INSERT INTO Users VALUES ('asistent1', 'Random', 'Assistent1', '37f514205db02ad131172e6db7534a13712e9a8f28c59a388289270d66b9b786');
INSERT INTO Assistents VALUES ('asistent1');

INSERT INTO Users VALUES ('asistent2', 'Random', 'Assistent2', '2789116cfc062aca0eae664c6d037390e45b146aa7c287cf056b9d86bb6e08bc');
INSERT INTO Assistents VALUES ('asistent2');

INSERT INTO StudentsToCourses VALUES ('stud1', 1, 2020);
INSERT INTO StudentsToCourses VALUES ('stud2', 1, 2020);
INSERT INTO StudentsToCourses VALUES ('stud3', 1, 2020);
INSERT INTO StudentsToCourses VALUES ('stud4', 2, 2020);
INSERT INTO StudentsToCourses VALUES ('stud5', 2, 2020);

INSERT INTO AssistentsToCourses VALUES ('asistent1', 2, 2020);
INSERT INTO AssistentsToCourses VALUES ('asistent2', 1, 2020);
INSERT INTO AssistentsToCourses VALUES ('stud4', 1, 2020);
INSERT INTO AssistentsToCourses VALUES ('prof2', 2, 2020);

INSERT INTO Assignments (AssignmentName, Type, Deadline, CourseID, SeriesID)
VALUES ('Laborator 1', 'Laborator', '16-10-2019 00:00:00', 1, 1);

INSERT INTO Assignments (AssignmentName, Type, Deadline, CourseID, SeriesID)
VALUES ('Tema 1', 'Tema', '2019-11-09', 1, 1);

INSERT INTO Assignments (AssignmentName, Type, Deadline, CourseID, SeriesID)
VALUES ('Laborator 1', 'Laborator', '2020-3-19', 2, 1);

INSERT INTO Assignments (AssignmentName, Type, Deadline, CourseID, SeriesID)
VALUES ('Laborator 2', 'Laborator', '2020-3-26', 2, 1);

INSERT INTO Assignments (AssignmentName, Type, Deadline, CourseID, SeriesID)
VALUES ('Tema 1', 'Tema', '2020-10-04', 2, 1);