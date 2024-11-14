package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	v := &initialVotes
	return v
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	var v int
	if counter == nil {
		v = 0
	} else {
		v = *counter
	}
	return v
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	er := &ElectionResult{candidateName, votes}
	return er
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	str := result.Name + " (" + fmt.Sprint(result.Votes) + ")"
	return str
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	if _, exists := results[candidate]; exists {
		results[candidate]--
	}
}