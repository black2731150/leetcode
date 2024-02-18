SELECT  p.product_name,s.year,s.price
FROM Sales s,Product p
WHERE p.product_id=s.product_id