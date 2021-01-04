#### 题目
<p>编写一个 SQL 查询，获取 <code>Employee</code>&nbsp;表中第二高的薪水（Salary）&nbsp;。</p>

<pre>+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
</pre>

<p>例如上述&nbsp;<code>Employee</code>&nbsp;表，SQL查询应该返回&nbsp;<code>200</code> 作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回 <code>null</code>。</p>

<pre>+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
</pre>


 #### 题解
 1、使用limit获取
 ```mysql
select distinct Salary 
from Employee 
order by Salary desc 
limit 1 offset 1
```
因为本表可能只有一条记录，不能返回null，所以使用临时表
 
 ```mysql
select 
(select distinct Salary 
from Employee 
order by Salary desc 
limit 1 offset 1) as SecondHighestSalary
```

 2、使用ifnull来判断
 ```mysql
select 
ifnull ((select distinct Salary 
from Employee 
order by Salary desc 
limit 1 offset 1),null) as SecondHighestSalary
```