package TournamentWinner

// O(n) time | O(k) space - where n is the number of competitions and k is the number of teams
func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentBestTeam := "" // keep track of the current winning team
	scoreboard := make(map[string]int) // make use of a hash/map to keep track of scores

	for round, _ := range competitions { // iterate through the rounds
			result := results[round] // retrive the results for the current round
			homeTeam := competitions[round][0]
			awayTeam := competitions[round][1]

			winningTeam := homeTeam // magic logic
			if result == 0 {
				winningTeam = awayTeam
			}

			scoreboard[winningTeam] += 3 // allocate winning points

			if scoreboard[winningTeam] > scoreboard[currentBestTeam] {
				currentBestTeam = winningTeam // check if round winner is now in the lead
			}
	}
	return currentBestTeam // return last current winner
}

