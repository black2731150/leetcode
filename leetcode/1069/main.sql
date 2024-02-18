SELECT p.product_id,SUM(s.quantity) AS total_quantity
FROM Sales s,Product p
WHERE  s.product_id=p.product_id
GROUP BY p.product_id