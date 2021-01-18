package code

import "sort"

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailToIndex := map[string]int{}   // 记录有多少个邮箱
	emailToName := map[string]string{} // 记录邮箱对应的用户
	for _, account := range accounts {
		name := account[0]
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = name
			}
		}
	}

	fa := make([]int, len(emailToIndex))
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	merge := func(x, y int) {
		fa[find(x)] = find(y)
	}

	for _, account := range accounts {
		index := emailToIndex[account[1]]
		for _, email := range account[2:] {
			merge(emailToIndex[email], index)
		}
	}

	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}
