WITH
tmp AS(
    SELECT problem_id,ROUND(likes*100/(likes+dislikes),3) AS xhl
    FROM Problems
)

SELECT problem_id
FROM tmp
WHERE xhl<60
ORDER BY problem_id 