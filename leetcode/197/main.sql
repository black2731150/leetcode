SELECT w2.id
FROM Weather w1,Weather w2
WHERE (DATE(w1.recordDate) = DATE(w2.recordDate)-INTERVAL 1 DAY) AND w2.temperature>w1.temperature;