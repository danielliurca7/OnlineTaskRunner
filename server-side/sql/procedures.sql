DELIMITER $$

CREATE PROCEDURE CreateAssignment (
    IN  _AssignmentName VARCHAR(50),
    IN  _Type VARCHAR(10),
    IN  _Deadline DATETIME,
    IN  _CourseID INT,
    OUT ERROR VARCHAR(30)
)
BEGIN
    DECLARE courseExists BOOLEAN DEFAULT FALSE;

    SET courseExists = EXISTS(
        SELECT * FROM Courses
        WHERE CourseID = _CourseID
    );

    IF courseExists = 1 THEN
        INSERT INTO
        Assignments ( AssignmentName,  Type,  Deadline,  CourseID)
        VALUES      (_AssignmentName, _Type, _Deadline, _CourseID);

        SET ERROR = '';
    ELSE
        SET ERROR = 'Course not found';
    END IF;
END$$

CREATE PROCEDURE GradeAssignment (
    IN  _StudentName VARCHAR(30),
    IN  _Grade INT,
    IN  _GradeTime DATETIME,
    IN  _AssistentName VARCHAR(30),
    IN  _AssignmentID INT,
    OUT ERROR VARCHAR(30)
)
BEGIN
    DECLARE studentExists BOOLEAN DEFAULT FALSE;
    DECLARE assignmentExists BOOLEAN DEFAULT FALSE;
    DECLARE assistentExists BOOLEAN DEFAULT FALSE;

    SET studentExists = EXISTS(
        SELECT * FROM Students
        WHERE StudentName = _StudentName
    );

    IF studentExists THEN
        SET assignmentExists = EXISTS(
            SELECT * FROM Assignments
            WHERE AssignmentID = _AssignmentID
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
        Grades ( StudentName,  Grade,  GradeTime,  AssistentName,  AssignmentID)
        VALUES (_StudentName, _Grade, _GradeTime, _AssistentName, _AssignmentID);

        SET ERROR := '';
    END IF;
END$$

DELIMITER ;