CREATE TABLE IF NOT EXISTS Users (
    UserName VARCHAR(30) NOT NULL,
    FirstName VARCHAR(30),
    LastName VARCHAR(30),
    PasswordHash CHAR(64),
    PRIMARY KEY (UserName)
);

CREATE TABLE IF NOT EXISTS Series (
    SeriesID INT NOT NULL AUTO_INCREMENT,
    SchoolYear INT NOT NULL,
    SeriesName VARCHAR(2) NOT NULL,
    PRIMARY KEY (SeriesID)
);

CREATE TABLE IF NOT EXISTS Classes (
    GroupID INT NOT NULL AUTO_INCREMENT,
    GroupNumber INT NOT NULL,
    SeriesID INT NOT NULL,    
    PRIMARY KEY (GroupID),
    FOREIGN KEY (SeriesID) REFERENCES Series (SeriesID)
);

CREATE TABLE IF NOT EXISTS Students (
    StudentName VARCHAR(30) NOT NULL,
    GroupID INT NOT NULL,
    PRIMARY KEY (StudentName),
    FOREIGN KEY (StudentName) REFERENCES Users (UserName),
    FOREIGN KEY (GroupID) REFERENCES Classes (GroupID)
);

CREATE TABLE IF NOT EXISTS Assistents (
    AssistentName VARCHAR(30) NOT NULL,
    PRIMARY KEY (AssistentName),
    FOREIGN KEY (AssistentName) REFERENCES Users (UserName)
);

CREATE TABLE IF NOT EXISTS Professors (
    ProfessorName VARCHAR(30) NOT NULL,
    PRIMARY KEY (ProfessorName),
    FOREIGN KEY (ProfessorName) REFERENCES Users (UserName)
);

CREATE TABLE IF NOT EXISTS Courses (
    CourseID INT NOT NULL AUTO_INCREMENT,
    CourseName VARCHAR(50) NOT NULL,
    SeriesID INT NOT NULL,
    Abbreviation VARCHAR(5),    
    ProfessorName VARCHAR(30),
    PRIMARY KEY (CourseID),
    FOREIGN KEY (SeriesID) REFERENCES Series (SeriesID),
    FOREIGN KEY (ProfessorName) REFERENCES Professors (ProfessorName)
);

CREATE TABLE IF NOT EXISTS Assignments (
    AssignmentID INT NOT NULL AUTO_INCREMENT,
    AssignmentName VARCHAR(50) NOT NULL,
    Type VARCHAR(10) NOT NULL,
    Deadline DATETIME,
    CourseID INT NOT NULL,
    PRIMARY KEY (AssignmentID),
    FOREIGN KEY (CourseID) REFERENCES Courses (CourseID),
    FOREIGN KEY (SeriesID) REFERENCES Series (SeriesID)
);

CREATE TABLE IF NOT EXISTS Grades (
    GradeID INT NOT NULL AUTO_INCREMENT,
    StudentName VARCHAR(30) NOT NULL,
    Grade INT NOT NULL,
    GradeTime DATETIME,
    AssistentName VARCHAR(30),
    AssignmentID INT NOT NULL,
    PRIMARY KEY (GradeID),
    FOREIGN KEY (StudentName) REFERENCES Students (StudentName),
    FOREIGN KEY (AssistentName) REFERENCES Assistents (AssistentName),
    FOREIGN KEY (AssignmentID) REFERENCES Assignments (AssignmentID),
    FOREIGN KEY (StudentName) REFERENCES Students (StudentName)
);

CREATE TABLE IF NOT EXISTS StudentsToCourses (
    StudentName VARCHAR(30) NOT NULL,
    CourseID INT NOT NULL,
    SchoolYear INT NOT NULL,
    PRIMARY KEY (StudentName, CourseID),
    FOREIGN KEY (StudentName) REFERENCES Students (StudentName),
    FOREIGN KEY (CourseID) REFERENCES Courses (CourseID)
);

CREATE TABLE IF NOT EXISTS AssistentsToCourses (
    AssistentName VARCHAR(30) NOT NULL,
    CourseID INT NOT NULL,
    SchoolYear INT NOT NULL,
    PRIMARY KEY (AssistentName, CourseID),
    FOREIGN KEY (AssistentName) REFERENCES Assistents (AssistentName),
    FOREIGN KEY (CourseID) REFERENCES Courses (CourseID)
);
