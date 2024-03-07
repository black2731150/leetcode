WITH
tmp1 AS (
    SELECT employee_id
    FROM Employees
    WHERE employee_id NOT IN (SELECT employee_id FROM Salaries)
),
tmp2 AS (
    SELECT employee_id
    FROM Salaries
    WHERE employee_id NOT IN (SELECT employee_id FROM Employees)
)

SELECT employee_id
FROM tmp1
UNION
SELECT employee_id
FROM tmp2
ORDER BY employee_id