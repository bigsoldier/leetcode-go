#### 题目
<p>编写一个 SQL 查询，查找所有至少连续出现三次的数字。</p>

<pre>+----+-----+
| Id | Num |
+----+-----+
| 1  |  1  |
| 2  |  1  |
| 3  |  1  |
| 4  |  2  |
| 5  |  1  |
| 6  |  2  |
| 7  |  2  |
+----+-----+
</pre>

<p>例如，给定上面的 <code>Logs</code> 表， <code>1</code> 是唯一连续出现至少三次的数字。</p>

<pre>+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
</pre>


 #### 题解
 # Write your MySQL query statement below
 select DISTINCT l1.Num as ConsecutiveNums
 from Logs l1,Logs l2,Logs l3
 where l1.Id=l2.Id-1 
 and l2.Id=l3.Id-1
 and l1.Num=l2.Num 
 and l2.Num=l3.Num