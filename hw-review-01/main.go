package main

import "fmt"

// Use constants for score thresholds
const (
	excellentThreshold = 90
	goodThreshold      = 75
	passThreshold      = 60
	maxStudents        = 5
)

// TODO(human): Implement the analyzeScores function
// Requirements:
//   - Use a for loop to iterate over the scores slice
//   - For each score, use switch (no-expression form) to print the grade:
//     >= 90: "Excellent"
//     >= 75: "Good"
//     >= 60: "Pass"
//     < 60:  "Fail"
//   - Use an if statement with a short variable declaration (if x := ...; condition)
//     to track and print the highest score
//   - At the end, calculate and print the average score
//
// Expected output for the test data:
//
//	Score 85: Good
//	Score 92: Excellent
//	Score 58: Fail
//	Score 76: Good
//	Score 95: Excellent
//	Highest: 95
//	Average: 81
func analyzeScores(scores [maxStudents]int) {

}

func main() {
	fmt.Println("=== Score Analyzer ===")

	scores := [maxStudents]int{85, 92, 58, 76, 95}

	analyzeScores(scores)
}
