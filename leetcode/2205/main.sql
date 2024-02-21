CREATE FUNCTION getUserIDs(startDate DATE, endDate DATE, minAmount INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      SELECT COUNT(DISTINCT user_id ) AS user_cnt
      FROM Purchases
      WHERE DATE(startDate) <= DATE(time_stamp) AND DATE(endDate) > DATE(time_stamp) AND amount >=  minAmount
  );
END