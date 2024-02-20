WITH
tmp AS (
    SELECT user_id,SUM(distance) AS traveled_distance
    FROM Rides
    GROUP BY user_id
)

SELECT u.user_id,u.name,COALESCE(t.traveled_distance,0)  AS 'traveled distance' 
FROM Users u LEFT JOIN tmp t ON t.user_id=u.user_id
ORDER BY u.user_id