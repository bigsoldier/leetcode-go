#### 题目
<p>编写一个 SQL 查询，获取 <code>Employee</code> 表中第&nbsp;<em>n&nbsp;</em>高的薪水（Salary）。</p>

<pre>+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
</pre>

<p>例如上述&nbsp;<code>Employee</code>&nbsp;表，<em>n = 2&nbsp;</em>时，应返回第二高的薪水&nbsp;<code>200</code>。如果不存在第&nbsp;<em>n&nbsp;</em>高的薪水，那么查询应返回&nbsp;<code>null</code>。</p>

<pre>+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
</pre>


 #### 题解
 1、单表查询
 ```mysql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    Set N := N-1; // 因为limit和offset字段后面只接受正整数
  RETURN (
      # Write your MySQL query statement below.
      select salary 
      from Employee
      group by salary
      order by salary desc
      limit N,1
  );
END
```
 2、子查询
 排名N的薪水意味着存在N-1高的薪水，所以这里去重后的N-1个
 ```mysql
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      select distinct e.salary 
      from Employee e
      where 
      (select count(distinct salary) from Employee where salary>e.salary) = N-1
  );
END
```