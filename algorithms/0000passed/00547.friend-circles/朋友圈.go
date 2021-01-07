package code

func findCircleNum(isConnected [][]int) (ans int) {
	isVisited := make([]bool, len(isConnected))
	var dfs func(i int)
	dfs = func(i int) {
		isVisited[i] = true
		for index, connect := range isConnected[i] {
			if connect == 1 && !isVisited[index] {
				dfs(index)
			}
		}
	}
	for i, visited := range isVisited {
		if !visited {
			ans++
			dfs(i)
		}
	}
	return
}
