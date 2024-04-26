package main

import "fmt"

func main() {
	Board := [3][3]string{}
	player := "x"
	for {
		fmt.Println("\nplayer:", player)
		var row int
		fmt.Println("enter number of row 0 or 1 or 2")
		fmt.Scanln(&row)
		if row < 0 || row > 2 {
			fmt.Println("invalid number of row : please enter 0 , 1 or 2")
			continue
		}
		var column int
		fmt.Println("enter number of col 0 or 1 or 2")
		fmt.Scanln(&column)
		if column < 0 || column > 2 {
			fmt.Println("invalid number of column : please enter 0 , 1 or 2")
			continue
		}
		if Board[row][column] == "" {
			Board[row][column] = player
		} else {
			fmt.Println("invalid row and column", row, column)
			continue
		}
		fmt.Println(Board[0])
		fmt.Println(Board[1])
		fmt.Println(Board[2])
		for i := 0; i < 3; i++ {
			if (Board[i][0] == Board[i][1] && Board[i][1] == Board[i][2] && Board[i][2] == player) ||
				(Board[0][i] == Board[1][i] && Board[1][i] == Board[2][i] && Board[2][i] == player) {
				fmt.Println("Game ended, winner is player:", player)
				return
			}
		}
		if (Board[0][0] == Board[1][1] && Board[1][1] == Board[2][2] && Board[2][2] == player) || (Board[0][2] == Board[1][1] && Board[1][1] == Board[2][0] && Board[2][0] == player) {
			fmt.Println("Game ended, winner is player:", player)
			return
		}
		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}

	}
}

