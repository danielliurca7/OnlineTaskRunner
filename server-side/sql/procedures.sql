DELIMITER $$

CREATE PROCEDURE GradeAssignment (
    IN  _AssignmentName VARCHAR(50),
    IN  _CourseName VARCHAR(50),
    IN  _SchoolYear INT,
    IN  _SeriesName VARCHAR(2),
    IN  _StudentName VARCHAR(30),
    IN  _GroupNumber INT,
    IN  _AssistentName VARCHAR(30),
    IN  _Grade INT,
    OUT ERROR VARCHAR(30)
)
BEGIN
    DECLARE studentExists BOOLEAN DEFAULT FALSE;
    DECLARE assignmentExists BOOLEAN DEFAULT FALSE;
    DECLARE assistentExists BOOLEAN DEFAULT FALSE;

    SET studentExists = EXISTS(
        SELECT * FROM Student
        WHERE StudentName = _StudentName AND
              GroupNumber = _GroupNumber AND
              SeriesName  = _SeriesName AND
              SchoolYear  = _SchoolYear
    );

    IF studentExists THEN
        SET assignmentExists = EXISTS(
            SELECT * FROM Assignemts
            WHERE AssignmentName = _AssignmentName AND
                  CourseName     = _CourseName AND
                  SchoolYear     = _SchoolYear AND
                  SeriesName     = _SeriesName 
        );

        IF assignmentExists THEN
            SET assistentExists = EXISTS(SELECT * FROM Assistents WHERE AssistentName = _AssistentName);

            IF NOT assistentExists THEN
                SET ERROR := 'Assistent not found';
            END IF;
        ELSE
            SET ERROR := 'Assignment not found';
        END IF;
    ELSE
        SET ERROR := 'Student not found';
    END IF;

    IF studentExists AND assignmentExists AND assistentExists THEN
        INSERT INTO
        Grades ( AssignmentName,  CourseName,  SchoolYear,  SeriesName,   StudentName,  GroupNumber,  AssistentName,  Grade, GradeTime)
        VALUES (_AssignmentName, _CourseName, _SchoolYear, _SeriesName,  _StudentName, _GroupNumber, _AssistentName, _Grade, NOW());

        SET ERROR := '';
    END IF;
END$$

CREATE PROCEDURE CreateAssignment (
    IN  _AssignmentName VARCHAR(50),
    IN  _CourseName VARCHAR(50),
    IN  _SchoolYear INT,
    IN  _SeriesName VARCHAR(2),
    IN  _PathToRequirement VARCHAR(100),
    OUT ERROR VARCHAR(30)
)
BEGIN
    DECLARE courseExists BOOLEAN DEFAULT FALSE;

    SET courseExists = EXISTS(
        SELECT * FROM Courses
        WHERE CourseName = _CourseName AND
              SchoolYear = _SchoolYear AND
              SeriesName = _SeriesName 
    );

    IF courseExists = 1 THEN
        INSERT INTO
        Assignments ( AssignmentName,  CourseName,  SchoolYear,  SeriesName,  PathToRequirement)
        VALUES      (_AssignmentName, _CourseName, _SchoolYear, _SeriesName, _PathToRequirement);

        SET ERROR = '';
    ELSE
        SET ERROR = 'Course not found';
    END IF;
END$$

DELIMITER ;