#### 题目
<p>给定一个<strong>非空</strong>整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。</p>

<p><strong>说明：</strong></p>

<p>你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [2,2,3,2]
<strong>输出:</strong> 3
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [0,1,0,1,0,1,99]
<strong>输出:</strong> 99</pre>


 #### 题解
 HashSet
 ```go
func singleNumber(nums []int) int {
	var total,sum int
	var m = make(map[int]bool)
	for _, num := range nums {
		total += num
		if _, ok := m[num]; ok {
			continue
		}
		sum += num
		m[num] = true
	}
	return (3*sum-total)/2
}
```
 时间复杂度O(n),空间复杂度O(n)
 
 位运算
 
 我们想要达到的效果
         a   b
 初始状态  0   0   
 第一次见  x   0
 第二次见  0   x
 第三次见  0   0
 
 而异或的规则
          a   
  初始状态  0      
  第一次见  x   
  第二次见  0   
  
  我们遍历元素时，第一次碰到将x赋值给a，第二次碰到将x赋值给b，
  所以我们可以写出
  ```text
    a = (a^num)
    b = (b^num)
```
  同时只有当a是x时，我们才会将0赋值给b，那要怎么写？
  我们知道如果b=0，那么b^num=x,而x&~x=0，而此时a=x,
  所以 第二行为 b=(b^num)&~a
  
  同理，a=(a^num)&~b
  
  所以最终的变换条件为
  ```text
    a=(a^num)&~b
    b=(b^num)&~a
```
  
  ```go
func singleNumber(nums []int) int {
	var once,twice int
	for _, num := range nums {
		once = ^twice&(once^num)
		twice = ^once&(twice^num)
	}
	return once
}
```
 时间复杂度O(n),空间复杂度O(1)