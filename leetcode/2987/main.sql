WITH tmp AS (
    SELECT city,AVG(price) AS avg
    FROM Listings
    GROUP BY city
)

SELECT city
FROM tmp
WHERE avg>(
    SELECT AVG(price)
    FROM Listings
)
ORDER BY city