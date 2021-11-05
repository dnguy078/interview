package main

func combinationSum(candidates []int, target int) [][]int {
	var list [][]int
	var temp []int
	backtrack(&list, &temp, candidates, target, 0)
	return [][]int{}
}

func backtrack(list *[][]int, temp *[]int, candidates []int, remaining, start int) {
	if remaining < 0 {
		return
	} else if remaining == 0 {
		*list = append(*list, *temp)
	} else {
		for i := start; i < len(candidates); i++ {
			*temp = append(*temp, candidates[i])
			backtrack(list, temp, candidates, remaining-candidates[i], i)
			// temp = *temp[:len(*temp)-1]
		}
	}

}

func combinationSum3(k int, n int) [][]int {

	// var output []int
	// output := func dfs(build []int, num int, currentSum int) {

	// }

	return [][]int{}
}
