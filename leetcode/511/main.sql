SELECT a1.player_id AS player_id,a1.event_date AS first_login
FROM Activity a1
LEFT JOIN Activity a2 ON a1.player_id = a2.player_id AND a1.event_date > a2.event_date
WHERE a2.player_id IS NULL
GROUP BY a1.player_id;