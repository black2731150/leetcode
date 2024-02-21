WITH
NG AS (
    SELECT COUNT(*) AS ngoods
    FROM NewYork
    WHERE score>=90
),
CG AS (
    SELECT COUNT(*) AS cgoods
    FROM California
    WHERE score>=90
)


SELECT CASE 
    WHEN ngoods>cgoods THEN  "New York University"
    WHEN ngoods<cgoods THEN "California University"
    ELSE  "No Winner"
END AS winner
FROM NG,CG