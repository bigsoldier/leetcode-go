#### 题目
罗马数字包含以下七种字符：I,V,X,L,C,D,M
```$xslt
字符      数值
I           1
V           5
X           10
L           50
C           100
D           500
M           1000
```
例如，罗马数字2写作II，12写作XII，27写作XXVII。

通常情况下，罗马数字中小的数字在大的数字右边。但存在特例，例如4写作IV。这种特殊规则适用于以下6种情况：
- I 可以放在 V(5) 和 X(10) 的左边，来表示 4 和 9
- X 可以放在 L(50) 和 C(100) 的左边，来表示 40 和 90
- C 可以放在 D(500) 和 M(1000) 的左边，来表示 400 和 900

给定一个整数，将其转为罗马数字，输入确保在1到3999的范围内。

#### 题解
贪心算法，将所有的情况列举出来
map = {1000: "M", 900: "CM", 500: "D", 400: "CD", 100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}

```go
func intToRoman(num int) string {
	var nums = []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	var romans = []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
	var i int
	var Roman string

	for i < len(nums) {
		for num >= nums[i] {
			Roman += romans[i]
			num -= nums[i]
		}
		i++
	}
	return Roman
}
```
![](https://raw.githubusercontent.com/betterfor/cloudImage/master/images/2020-02-06/001201.png)
时间复杂度和空间复杂度都是O(1)