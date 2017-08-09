-- 查询表中第N高的薪水
-- 先按Salary倒序排序, 然后 limit N-1, 1 即可取到第 N 高的薪水
create function getNthHighestSalary(N INT) returns INT
begin
    set N = N - 1;
    return (
        select distinct Salary from Employee order by Salary desc limit N, 1
    );
end
