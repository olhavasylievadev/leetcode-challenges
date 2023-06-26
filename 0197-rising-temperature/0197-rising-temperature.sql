# Write your MySQL query statement below

SELECT 
    A.id
FROM 
    Weather A, Weather B
WHERE
    DATEDIFF(A.recordDate, B.recordDate) = 1
    AND
    B.temperature < A.Temperature
