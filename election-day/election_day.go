package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	// For min egen del:
	// counter <- er adressen til selve tallet f.eks 2
	// &counter <- er adressen til variabelen counter
	// *counter <- er da verdien som er lagret i counter

	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter == nil {
		return
	}
	// Verdien av counter (*counter) skal være like verdien av counter (*counter) + increment
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}
