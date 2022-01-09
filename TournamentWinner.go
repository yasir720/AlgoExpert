//  VERY DIRTY CODE. MUST BE RE-DONE
package TournamentWinner

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentBestTeam := ""
	scoreboard := make(map[string]int)

	for round, _ := range competitions {
			result := results[round]
			homeTeam := competitions[round][0]
			awayTeam := competitions[round][1]

			winningTeam := homeTeam
			if result == 0 {
				winningTeam = awayTeam
			}

			scoreboard[winningTeam] += 3

			if scoreboard[winningTeam] > scoreboard[currentBestTeam] {
				currentBestTeam = winningTeam
			}
	}
	return currentBestTeam
}

