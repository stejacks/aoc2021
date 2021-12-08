package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	number int
	marked bool
}

type Board struct {
	squares [5][5]Square
}

func readinput(input string) []string {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func stampboard(board *Board, draw int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.squares[i][j].number == draw {
				board.squares[i][j].marked = true
			}
		}
	}
}

func iswinner(board Board) bool {
	// Are columns all marked?
	winner := false
	for column := 0; column < 5; column++ {
		winner = board.squares[column][0].marked && board.squares[column][1].marked && board.squares[column][2].marked && board.squares[column][3].marked && board.squares[column][4].marked
		if winner == true {
			break
		}
	}
	if winner != true {
		for row := 0; row < 5; row++ {
			winner = board.squares[0][row].marked && board.squares[1][row].marked && board.squares[2][row].marked && board.squares[3][row].marked && board.squares[4][row].marked
			if winner == true {
				break
			}
		}
	}
	return winner
}

func calculatescore(board Board, number int) int {
  var score int
	for row := 0; row < 5; row++ {
	  for column := 0; column < 5; column++ {
      if board.squares[column][row].marked == false {
        score=score + board.squares[column][row].number
      }
    }
  }
  return score * number
}
    

func main() {

	// Read in the boards

	input := "input.txt"
	inputlines := readinput(input)
	rawdraws := strings.Split(inputlines[0], ",")

	var drawnnumbers []int
	for i := 0; i < len(rawdraws); i++ {
		draw, err := strconv.Atoi(rawdraws[i])
		if err != nil {
			log.Fatal(err)
		}
		drawnnumbers = append(drawnnumbers, draw)
	}

	inputboards := inputlines[2:]

	var boards []Board
	var tmpboard Board
	var emptyboard Board
	for i := 0; i < len(inputboards); i++ {
		if inputboards[i] == "" {
			boards = append(boards, tmpboard)
			tmpboard = emptyboard
		} else {
			rawnumbers := strings.Fields(inputboards[i])
			for j := 0; j < len(rawnumbers); j++ {
				boardnumber := i % 6
				number, err := strconv.Atoi(rawnumbers[j])
				if err != nil {
					log.Fatal(err)
				}
				tmpboard.squares[boardnumber][j] = Square{number, false}
			}
		}
	}

bingo:
	for i := 0; i < len(drawnnumbers); i++ {
		for j := 0; j < len(boards); j++ {
			stampboard(&boards[j], drawnnumbers[i])
			result := iswinner(boards[j])
			if result == true {
				winningboard := boards[j]
				winningnumber := drawnnumbers[i]
				score := calculatescore(winningboard, winningnumber)
				fmt.Println(score)
				break bingo
			}
		}
	}
}
