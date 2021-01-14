#### 题目
<p>在<strong>二维</strong>平面上计算出两个<strong>由直线构成的</strong>矩形重叠后形成的总面积。</p>

<p>每个矩形由其左下顶点和右上顶点坐标表示，如图所示。</p>

<p><img alt="Rectangle Area" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rectangle_area.png"></p>

<p><strong>示例:</strong></p>

<pre><strong>输入:</strong> -3, 0, 3, 4, 0, -1, 9, 2
<strong>输出:</strong> 45</pre>

<p><strong>说明:</strong> 假设矩形面积不会超出&nbsp;<strong>int&nbsp;</strong>的范围。</p>


 #### 题解
 如果两个矩阵不相交，面积为 (C-A)*(D-B)+(G-E)*(H-F)
 
 如果两个矩阵相交，那么就是两个矩阵的面积减去相交的面积。
 
 相交的情况我们可以观察到,第二个矩阵的横纵坐标包含在第一个矩阵中
 ```go
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	lx := max(A,E)
	rx := min(C,G)
	ly := max(B,F)
	ry := min(D,H)
	area1,area2 := (C-A)*(D-B),(G-E)*(H-F)
	if lx > rx || ly > ry {
		return area1+area2
	}
	return area1+area2-(rx-lx)*(ry-ly)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```