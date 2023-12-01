package main

import (
	"fmt"
	"sort"
)

type Team struct {
	team rune
	rank []int
}

func rankTeams(votes []string) string {
	theMap := make(map[rune][]int)

	for _, ch := range votes[0] {
		theMap[ch] = make([]int, len(votes[0]))
	}

	for _, vote := range votes {
		for rank, team := range vote {
			theMap[team][rank]++
		}
	}

	Teams := make([]Team, len(votes[0]))
	index := 0
	for team, rank := range theMap {
		Teams[index] = Team{
			team: team,
			rank: rank,
		}
		index++
	}

	sort.Slice(Teams, func(i, j int) bool {
		for i2 := range Teams[i].rank {
			if Teams[i].rank[i2] != Teams[j].rank[i2] {
				return Teams[i].rank[i2] > Teams[j].rank[i2]
			}
		}
		return Teams[i].team < Teams[j].team
	})

	answer := make([]rune, len(votes[0]))
	for i := range answer {
		answer[i] = Teams[i].team
	}

	return string(answer)

}

func main() {
	votes := []string{"FVSHJIEMNGYPTQOURLWCZKAX", "AITFQORCEHPVJMXGKSLNZWUY", "OTERVXFZUMHNIYSCQAWGPKJL", "VMSERIJYLZNWCPQTOKFUHAXG", "VNHOZWKQCEFYPSGLAMXJIUTR", "ANPHQIJMXCWOSKTYGULFVERZ", "RFYUXJEWCKQOMGATHZVILNSP", "SCPYUMQJTVEXKRNLIOWGHAFZ", "VIKTSJCEYQGLOMPZWAHFXURN", "SVJICLXKHQZTFWNPYRGMEUAO", "JRCTHYKIGSXPOZLUQAVNEWFM", "NGMSWJITREHFZVQCUKXYAPOL", "WUXJOQKGNSYLHEZAFIPMRCVT", "PKYQIOLXFCRGHZNAMJVUTWES", "FERSGNMJVZXWAYLIKCPUQHTO", "HPLRIUQMTSGYJVAXWNOCZEKF", "JUVWPTEGCOFYSKXNRMHQALIZ", "MWPIAZCNSLEYRTHFKQXUOVGJ", "EZXLUNFVCMORSIWKTYHJAQPG", "HRQNLTKJFIEGMCSXAZPYOVUW", "LOHXVYGWRIJMCPSQENUAKTZF", "XKUTWPRGHOAQFLVYMJSNEIZC", "WTCRQMVKPHOSLGAXZUEFYNJI"}
	fmt.Println(rankTeams(votes))
}
