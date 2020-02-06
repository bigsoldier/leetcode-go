
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