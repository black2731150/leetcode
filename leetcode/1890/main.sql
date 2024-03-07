WITH 
tmp AS (
    SELECT user_id,MAX(time_stamp) AS last_stamp
    FROM Logins
    WHERE YEAR(time_stamp) = 2020
    GROUP BY user_id
    
)

SELECT user_id,MAX(time_stamp)
FROM tmp
WHERE year = 2020