SELECT DISTINCT l1.user_id
FROM Loans l1,Loans l2
WHERE l1.user_id=l2.user_id AND l1.loan_type='Mortgage' AND l2.loan_type='Refinance'
ORDER BY l1.user_id 