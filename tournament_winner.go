package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	scores := make(map[string]int, len(competitions))
	for i, c := range competitions {
		home, away := c[0], c[1]
		if results[i] == 0 {
			scores[away]++
		} else {
			scores[home]++
		}
	}
	max := 0
	var winner string
	for team, score := range scores {
		if score >= max {
			max = score
			winner = team
		}
	}

	return winner
}
