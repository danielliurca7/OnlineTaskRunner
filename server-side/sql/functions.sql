DELIMITER $$

CREATE FUNCTION Login (
    _UserName VARCHAR(30),
    _PasswordHash CHAR(64)
)
RETURNS BOOLEAN
READS SQL DATA
BEGIN
    DECLARE isValid BOOLEAN DEFAULT FALSE;

    SET isValid = EXISTS(SELECT * FROM Users WHERE UserName = _UserName AND PasswordHash = _PasswordHash);

    RETURN isValid;
END$$

CREATE FUNCTION GetType (
    _UserName VARCHAR(30)
)
RETURNS VARCHAR(2)
READS SQL DATA
BEGIN
    DECLARE type VARCHAR(2) DEFAULT "";

    IF EXISTS(SELECT * FROM Students WHERE StudentName = _UserName) THEN
        SET type = CONCAT(type, "1");
    END IF;

    IF EXISTS(SELECT * FROM Assistents WHERE AssistentName = _UserName) THEN
        SET type = CONCAT(type, "2");
    END IF;
    
    IF EXISTS(SELECT * FROM Professors WHERE ProfessorName = _UserName) THEN
        SET type = CONCAT(type, "3");
    END IF;

    RETURN type;
END$$

CREATE FUNCTION IsValid (
    _UserName VARCHAR(30),
    _CourseName VARCHAR(50),
    _SchoolYear INT,
    _AssignmentName VARCHAR(50)
)
RETURNS BOOLEAN
READS SQL DATA
BEGIN
    DECLARE isStudent BOOLEAN;
    DECLARE isAssistent BOOLEAN;
    DECLARE isValid BOOLEAN DEFAULT FALSE;

    SET isStudent = EXISTS(
        SELECT * FROM
        Assignments NATURAL JOIN (
            SELECT * FROM
            Students NATURAL JOIN StudentsToCourses
            WHERE _UserName   = StudentName AND
                  _CourseName = CourseName  AND
                  _SchoolYear = SchoolYear
        ) qqqq WHERE AssignmentName = _AssignmentName 
    );

    IF isStudent THEN
        SET isValid = TRUE;
    ELSE
        SET isAssistent = EXISTS(
            SELECT * FROM
            Assignments NATURAL JOIN (
                SELECT * FROM
                Assistents NATURAL JOIN AssistentsToCourses
                WHERE _UserName   = AssistentName AND
                      _CourseName = CourseName  AND
                      _SchoolYear = SchoolYear
            ) qqqq WHERE AssignmentName = _AssignmentName 
        );

        IF isAssistent THEN
            SET isValid = TRUE;
        END IF;
    END IF;

    RETURN isValid;
END$$

DELIMITER ;