SELECT DISTINCT a.email AS 'Email'
FROM Person AS a,Person AS b
WHERE a.id != b.id AND a.email = b.email;

SELECT Email FROM
(
    SELECT Email,COUNT(Email) AS num
    FROM Person
    GROUP BY Email
) as statistic
WHERE num>1;