SELECT actor_id,director_id
FROM (
    SELECT actor_id,director_id,COUNT(*) AS num
    FROM ActorDirector
    GROUP BY actor_id,director_id
) AS tmp
WHERE num>=3;