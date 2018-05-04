# 给定一个用户表 Customers 和一个订单表 Orders，找到未下过单的用户
select Name AS Customers from Customers where id not in (select CustomerId from Orders);
