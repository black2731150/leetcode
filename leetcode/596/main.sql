
SELECT class
FROM (
    SELECT class,COUNT(*) AS class_num
    FROM Courses
    GROUP BY class
) AS help
WHERE class_num>=5
