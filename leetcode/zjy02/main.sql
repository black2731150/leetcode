WITH
first AS(
    SELECT MAX(salary) AS fristS
    FROM Employee
)
SELECT MAX(e.salary) AS SecondHighestSalary
FROM Employee e,first f
WHERE e.salary < f.fristS