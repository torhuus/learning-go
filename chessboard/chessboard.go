package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	toBeCounted, exists := cb[file]
	if !exists {
		fmt.Println("This file does not exist")
	}

	occupied := 0

	for _, v := range toBeCounted {
		if v {
			occupied = occupied + 1
		}
	}

	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	// Return 0 if the rank is invalid
	if rank < 1 || rank > 8 {
		return 0
	}

	occupied := 0

	for _, v := range cb {
		if v[rank-1] {
			occupied = occupied + 1
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squares := 0
	for range cb {
		squares = squares + 8
	}

	fmt.Println(squares)
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0
	// k = A
	// v = [true, false...]
	for _, v := range cb {
		for _, squares := range v {
			if squares {
				occupied = occupied + 1
			}
		}
	}
	return occupied
}
