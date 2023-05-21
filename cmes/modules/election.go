package modules

import (
	"fmt"
	"time"
)

type Candidate struct {
	Name          string
	Party         string
	CampaignTheme string
}

type Voter struct {
	Name       string
	Age        int
	Occupation string
}

type Election struct {
	Date       time.Time
	Location   string
	Candidates []Candidate
	Voters     []Voter
	Results    map[string]int
}

func NewElection(date time.Time, location string) *Election {
	return &Election{
		Date:     date,
		Location: location,
		Results:  make(map[string]int),
	}
}

func (e *Election) AddCandidate(name, party, theme string) {
	e.Candidates = append(e.Candidates, Candidate{
		Name:          name,
		Party:         party,
		CampaignTheme: theme,
	})
}

func (e *Election) AddVoter(name string, age int, occupation string) {
	e.Voters = append(e.Voters, Voter{
		Name:       name,
		Age:        age,
		Occupation: occupation,
	})
}

func (e *Election) RecordVote(candidateName string) {
	e.Results[candidateName]++
}

func (e *Election) GetResults() map[string]int {
	return e.Results
}

func (e *Election) PrintResults() {
	fmt.Println("Election Results:")
	for _, candidate := range e.Candidates {
		fmt.Printf("%s (%s): %d votes\n", candidate.Name, candidate.Party, e.Results[candidate.Name])
	}
}
