--找出 person 表中重复的 email
select email from person group by email having count(email) > 1;
