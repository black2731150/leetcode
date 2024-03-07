CREATE PROCEDURE getUserIDs(startDate DATE, endDate DATE, minAmount INT)
BEGIN
	# Write your MySQL query statement below.
    SELECT DISTINCT user_id
    FROM Purchases
    WHERE amount >= minAmount AND startDate <= time_stamp AND endDate >= time_stamp
    ORDER BY user_id;
END