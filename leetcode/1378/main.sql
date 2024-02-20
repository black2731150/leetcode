SELECT e.unique_id,s.name
FROM Employees s LEFT JOIN EmployeeUNI e ON s.id=e.id