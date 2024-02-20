SELECT ad_id,CASE 
    WHEN SUM(CASE 
        WHEN action = 'Clicked' THEN  1 
        ELSE  0
    END) + SUM(CASE 
        WHEN action = 'Viewed ' THEN  1
        ELSE  0
    END) !=0 THEN  ROUND(100 * SUM(CASE 
        WHEN action = 'Clicked' THEN  1
        ELSE  0
    END)/(SUM(CASE 
        WHEN action = 'Clicked' THEN  1 
        ELSE  0
    END) + SUM(CASE 
        WHEN action = 'Viewed ' THEN  1
        ELSE  0
    END)),2)
    ELSE  0
END ctr
FROM Ads
GROUP BY ad_id
ORDER BY ctr DESC,ad_id