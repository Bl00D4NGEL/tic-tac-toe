package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player int

type Grid struct {
	fields [9]Player
}

const (
	NoPlayer = iota
	FirstPlayer
	SecondPlayer
)

func main() {
	player1 := Player(FirstPlayer)
	player2 := Player(SecondPlayer)

	winConditions := GetWinConditions()

	currentPlayer := player1

	var grid Grid
	fmt.Printf("Player %s please start.\n", strconv.FormatInt(int64(currentPlayer), 10))
	PrintGrid(grid)

	for {
		fmt.Printf("Current player: %s\n", strconv.FormatInt(int64(currentPlayer), 10))

		index := GetUserSelection(grid)

		grid.fields[index] = currentPlayer

		PrintGrid(grid)

		if IsDraw(grid) {
			fmt.Println("Game is draw")
			return
		}

		if IsWon(grid, winConditions) {
			fmt.Println("Game won")
			return
		}

		if currentPlayer == player1 {
			currentPlayer = player2
		} else {
			currentPlayer = player1
		}
	}
}

func IsDraw(grid Grid) bool {
	for _, v := range grid.fields {
		if v == NoPlayer {
			return false
		}
	}
	return true
}

func IsWon(grid Grid, winConditions []WinCondition) bool {
	for _, winCondition := range winConditions {
		if winCondition.IsWon(grid) {
			return true
		}
	}

	return false
}

func GetUserSelection(grid Grid) int {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter field to occupy: ")
		text, _ := reader.ReadString('\n')

		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("Please enter a value (0-8)")
		} else {
			index, _ := strconv.Atoi(text)
			selectedField := grid.fields[index]
			if selectedField == NoPlayer {
				return index
			}
			fmt.Println("Field is already taken or invalid")
		}
	}
}

func PrintGrid(grid Grid) {
	fmt.Println("Current game status:")
	fmt.Println("-------")
	for i := 0; i < len(grid.fields); i += 3 {
		PrintRow(grid.fields[i:i+3], i)
		fmt.Println("-------")
	}
}

func PrintRow(fields []Player, fieldIndex int) {
	if len(fields) != 3 {
		return
	}

	fmt.Printf(
		"|%s|%s|%s|\n",
		GetPlayerIdentifier(fields[0], fieldIndex),
		GetPlayerIdentifier(fields[1], fieldIndex+1),
		GetPlayerIdentifier(fields[2], fieldIndex+2),
	)
}

func GetPlayerIdentifier(player Player, fieldIndex int) string {
	if player == NoPlayer {
		return strconv.FormatInt(int64(fieldIndex), 10)
	}

	if player == FirstPlayer {
		return "X"
	}

	if player == SecondPlayer {
		return "O"
	}

	return ""
}
