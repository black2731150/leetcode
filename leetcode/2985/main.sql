
WITH 
tmp1 AS (SELECT item_count,item_count*order_occurrences AS chengji FROM Orders),
fenzi AS (SELECT SUM(chengji) AS he FROM tmp1),
fenmu AS (SELECT SUM(order_occurrences) AS fm FROM Orders)

SELECT ROUND(he/fm,2) AS average_items_per_order 
FROM fenzi,fenmu