WITH 
tmp AS (
    SELECT customer_id,MAX(amount) AS max_amount
    FROM Store
    GROUP BY customer_id
)

SELECT COUNT(*) AS rich_count
FROM tmp
WHERE max_amount>500