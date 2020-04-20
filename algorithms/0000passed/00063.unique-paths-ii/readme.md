#### 题目
<p>一个机器人位于一个 <em>m x n </em>网格的左上角 （起始点在下图中标记为&ldquo;Start&rdquo; ）。</p>

<p>机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为&ldquo;Finish&rdquo;）。</p>

<p>现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？</p>

<p><img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/robot_maze.png" style="height: 183px; width: 400px;"></p>

<p>网格中的障碍物和空位置分别用 <code>1</code> 和 <code>0</code> 来表示。</p>

<p><strong>说明：</strong><em>m</em>&nbsp;和 <em>n </em>的值均不超过 100。</p>

<p><strong>示例&nbsp;1:</strong></p>

<pre><strong>输入:
</strong>[
&nbsp; [0,0,0],
&nbsp; [0,1,0],
&nbsp; [0,0,0]
]
<strong>输出:</strong> 2
<strong>解释:</strong>
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 <code>2</code> 条不同的路径：
1. 向右 -&gt; 向右 -&gt; 向下 -&gt; 向下
2. 向下 -&gt; 向下 -&gt; 向右 -&gt; 向右
</pre>


 #### 题解
 动态规划
 第一行和第一列的如果存在障碍，那么后续的所有都置0
 与62题类似，dp[i][j] = dp[i-1][j]+dp[i][j-1],如果该点是障碍，那么dp[i][j]=0
 ```go
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	
	height := len(obstacleGrid)
	width := len(obstacleGrid[0])
	
	obstacleGrid[0][0] = 1
	
	// 找出第一列的障碍，如果有一个有障碍，那么后续全部置0
	for i := 1; i < height; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	// 同理，找出第一行的障碍，如果有一个有障碍，那么后续全部置0
	for i := 1; i < width; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < height; i++ {
		for j := 0; j < width; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	
	return obstacleGrid[height-1][width-1]
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-04-20/006301.png)
时间复杂度O(n^2^),空间复杂度O(1)