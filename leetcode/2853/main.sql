WITH
EMAX AS (
    SELECT MAX(salary) AS max_E
    FROM Salaries
    WHERE department='Engineering'
),
MMAX AS (
    SELECT MAX(salary) AS max_M
    FROM Salaries
    WHERE department='Marketing'
)

SELECT ABS(max_E-max_M) AS salary_difference
FROM EMAX,MMAX