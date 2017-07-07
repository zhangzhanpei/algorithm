-- 查询奇数id并且描述不是boring的电影, 按评分由高到低排序
SELECT id, movie, description, rating FROM cinema WHERE description != 'boring' AND id % 2 = 1 ORDER BY rating DESC;
