SELECT *
FROM Employees e LEFT JOIN Salaries s ON e.employee_id=s.employee_id
UNION
SELECT *
FROM Employees e RIGHT JOIN Salaries s ON e.employee_id=s.employee_id
WHERE e.name IS NULL OR s.salary IS NULL

