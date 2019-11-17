DELIMITER $$

CREATE TRIGGER DeleteAssistent
AFTER DELETE
ON AssistentsToCourses
FOR EACH ROW
BEGIN
    DECLARE stillAssistent BOOLEAN DEFAULT FALSE;

    SET stillAssistent = EXISTS(SELECT * FROM AssistentsToCourses WHERE AssistentName = OLD.AssistentName);

    if NOT stillAssistent THEN
        DELETE FROM Assistents WHERE AssistentName = OLD.AssistentName;
    END IF;
END$$

DELIMITER ;