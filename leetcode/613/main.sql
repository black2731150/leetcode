

SELECT MIN(ABS(p1.x-p2.x)) AS shortest
FROM Point p1 JOIN Point p2 On p1.x!=p2.x