package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var cnt int
	for _, val := range cb[file] {
		if val {
			cnt++
		}
	}
	return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var cnt int

	for key := range cb {
		if rank > len(cb[key]) || rank < 1 {
			return 0
		}
		if cb[key][rank - 1] {
			cnt++
		}
	}
	return cnt
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var cnt int

	for key := range cb {
		for range cb[key] {
			cnt++
		}
	}
	return cnt
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var cnt int

	for key := range cb {
		for _, val := range cb[key] {
			if val {
				cnt++
			}
		}
	}
	return cnt
}
