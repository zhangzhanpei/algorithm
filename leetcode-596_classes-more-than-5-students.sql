--找出 courses 表中学生人数大于等于 5 的班级(同一个学生多次报一个班级算一个)
select class from courses group by class having count(distinct(student)) >= 5;
