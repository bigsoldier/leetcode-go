package code

func majorityElement(nums []int) (ans []int) {
	n := len(nums)
	// 最多两个候选人
	cand1, cand2 := nums[0], nums[0]
	var count1, count2 int
	// 选出候选人，遇到不同的候选人消除选票
	for _, num := range nums {
		if num == cand1 {
			count1++
			continue
		}
		if num == cand2 {
			count2++
			continue
		}
		if count1 == 0 {
			cand1 = num
			count1++
			continue
		}
		if count2 == 0 {
			cand2 = num
			count2++
			continue
		}
		count1--
		count2--
	}

	// 重新计数
	count1, count2 = 0, 0
	for _, num := range nums {
		if cand1 == num {
			count1++
		} else if cand2 == num {
			count2++
		}
	}

	// 判断1/3
	if count1 > n/3 {
		ans = append(ans, cand1)
	}
	if count2 > n/3 {
		ans = append(ans, cand2)
	}
	return
}
