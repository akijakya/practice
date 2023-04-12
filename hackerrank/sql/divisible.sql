select *, 
CASE
    WHEN n%x=0 AND n%y=0 
        THEN 'true'
    ELSE 'false'
END AS res
from kata;

SELECT id, n % x = 0 AND n % y = 0 AS res FROM kata