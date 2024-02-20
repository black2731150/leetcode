WITH
METAL AS (
    SELECT symbol AS metal
    FROM Elements
    WHERE type='Metal'
),
Nonmetal AS (
    SELECT symbol AS nonmetal
    FROM Elements
    WHERE type='Nonmetal'
)

SELECT metal,nonmetal
FROM METAL JOIN Nonmetal