-- 查询面积大于300万平方公里或人口大于2500万的国家[大国]
SELECT name, population, area FROM World WHERE area >= 3000000 OR population >= 25000000
