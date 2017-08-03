--删除 person 表中重复的 email, 但保留 id 最小的 email 记录
delete p from person p, person tmp where p.email = tmp.email and p.id > tmp.id；
