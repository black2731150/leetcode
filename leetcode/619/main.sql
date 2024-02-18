SELECT MAX(a.num) AS num
FROM (SELECT num,COUNT(*) AS nn
FROM MyNumbers
GROUP BY num
HAVING nn=1
) AS a
