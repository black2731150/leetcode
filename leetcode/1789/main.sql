WITH
tmp AS (
    SELECT employee_id,COUNT(employee_id) AS cnum
    FROM Employee
    WHERE cnum =1
    GROUP BY employee_id
)

SELECT DISTINCT employee_id,department_id
FROM Employee 
WHERE primary_flag = 'Y' 
UNION
SELECT e.employee_id,e.department_id
FROM Employee e,tmp t
WHERE t.cnum=1