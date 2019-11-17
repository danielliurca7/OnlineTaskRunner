CREATE TABLE IF NOT EXISTS Users (
    UserName VARCHAR(30) NOT NULL,
    FirstName VARCHAR(30),
    LastName VARCHAR(30),
    PasswordHash CHAR(64),
    PRIMARY KEY (UserName)
);

CREATE TABLE IF NOT EXISTS Series (
    SchoolYear INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,
    PRIMARY KEY (SchoolYear, SeriesName)
);

CREATE TABLE IF NOT EXISTS Classes (
    GroupNumber INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,    
    SchoolYear INT NOT NULL,
    PRIMARY KEY (GroupNumber, SeriesName, SchoolYear),
    FOREIGN KEY (SchoolYear, SeriesName) REFERENCES Series (SchoolYear, SeriesName)
);

CREATE TABLE IF NOT EXISTS Students (
    StudentName VARCHAR(30) NOT NULL,
    GroupNumber INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,    
    SchoolYear INT NOT NULL,
    PRIMARY KEY (StudentName, GroupNumber, SeriesName, SchoolYear),
    FOREIGN KEY (StudentName) REFERENCES Users (UserName),
    FOREIGN KEY (GroupNumber, SeriesName, SchoolYear) REFERENCES Classes (GroupNumber, SeriesName, SchoolYear)
);

CREATE TABLE IF NOT EXISTS Assistents (
    AssistentName VARCHAR(30) NOT NULL,
    PRIMARY KEY (AssistentName)
);

CREATE TABLE IF NOT EXISTS Professors (
    ProfessorName VARCHAR(30) NOT NULL,
    PRIMARY KEY (ProfessorName)
);

CREATE TABLE IF NOT EXISTS Courses (
    CourseName VARCHAR(50) NOT NULL,
    SchoolYear INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,
    Abbreviation VARCHAR(4),    
    ProfessorName VARCHAR(30),
    PRIMARY KEY (CourseName, SchoolYear, SeriesName),
    FOREIGN KEY (SchoolYear, SeriesName) REFERENCES Series (SchoolYear, SeriesName),
    FOREIGN KEY (ProfessorName) REFERENCES Professors (ProfessorName)
);

CREATE TABLE IF NOT EXISTS Assignments (
    AssignmentName VARCHAR(50) NOT NULL,
    CourseName VARCHAR(50) NOT NULL,
    SchoolYear INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,
    PathToRequirement VARCHAR(100),
    PRIMARY KEY (AssignmentName, CourseName, SchoolYear, SeriesName),
    FOREIGN KEY (CourseName, SchoolYear, SeriesName) REFERENCES Courses (CourseName, SchoolYear, SeriesName)
);

CREATE TABLE IF NOT EXISTS Grades (
    AssignmentName VARCHAR(50) NOT NULL,
    CourseName VARCHAR(50) NOT NULL,
    SchoolYear INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,
    StudentName VARCHAR(30) NOT NULL,
    GroupNumber INT NOT NULL,
    Grade INT NOT NULL,
    AssistentName VARCHAR(30) NOT NULL,
    GradeTime DATETIME,
    PRIMARY KEY (AssignmentName, CourseName, SchoolYear, SeriesName, StudentName, GroupNumber),
    FOREIGN KEY (AssignmentName, CourseName, SchoolYear, SeriesName) REFERENCES Assignments (AssignmentName, CourseName, SchoolYear, SeriesName),
    FOREIGN KEY (StudentName, GroupNumber, SeriesName, SchoolYear) REFERENCES Students (StudentName, GroupNumber, SeriesName, SchoolYear),
    FOREIGN KEY (AssistentName) REFERENCES Assistents (AssistentName)
);

CREATE TABLE IF NOT EXISTS AssistentsToCourses (
    AssistentName VARCHAR(30) NOT NULL,
    CourseName VARCHAR(50) NOT NULL,
    SchoolYear INT NOT NULL,
    PRIMARY KEY (AssistentName, CourseName, SchoolYear),
    FOREIGN KEY (AssistentName) REFERENCES Assistents (AssistentName),
    FOREIGN KEY (CourseName, SchoolYear) REFERENCES Courses (CourseName, SchoolYear)
);

CREATE TABLE IF NOT EXISTS StudentsToCourses (
    StudentName VARCHAR(30) NOT NULL,
    CourseName VARCHAR(50) NOT NULL,
    SchoolYear INT NOT NULL,
    PRIMARY KEY (StudentName, CourseName, SchoolYear),
    FOREIGN KEY (StudentName) REFERENCES Students (StudentName),
    FOREIGN KEY (CourseName, SchoolYear) REFERENCES Courses (CourseName, SchoolYear)
);
