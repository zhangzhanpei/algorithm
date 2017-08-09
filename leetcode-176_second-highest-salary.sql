-- 查询表中第二高的薪水
-- 先在子查询中查询最高的薪水, 然后查询小于它的最高薪水即为第二高
select max(salary) as SecondHighestSalary from Employee where salary < (select max(salary) from Employee);
