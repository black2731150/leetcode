package main

import (
	"fmt"
	"sort"
	"strings"
)

type IDAndScore struct {
	id    int
	score int
}

func topStudents(positive_feedback []string, negative_feedback []string, report []string, student_id []int, k int) []int {

	theMap := make([]IDAndScore, 0)

	positiveMap := make(map[string]int, len(positive_feedback))
	negativeMap := make(map[string]int, len(negative_feedback))

	for _, v := range positive_feedback {
		positiveMap[v] = 3
	}

	for _, v := range negative_feedback {
		negativeMap[v] = -1
	}

	for i, r := range report {

		score := 0

		rWords := strings.Split(r, " ")

		for _, word := range rWords {
			if x, ok := positiveMap[word]; ok {
				score = score + x
				continue
			}

			if x, ok := negativeMap[word]; ok {
				score = score + x
				continue
			}
		}

		theMap = append(theMap, IDAndScore{id: student_id[i], score: score})
	}

	sort.Slice(theMap, func(i, j int) bool {
		if theMap[i].score == theMap[j].score {
			return theMap[i].id < theMap[j].id
		}
		return theMap[i].score > theMap[j].score
	})

	answer := []int{}
	for i := 0; i < k; i++ {
		answer = append(answer, theMap[i].id)
	}

	return answer
}

func main() {
	positive_feedback := []string{"m", "eveszfubew"}
	negative_feedback := []string{"iq", "etwuedg", "egpakyk", "da", "qkmhvgxg", "q", "zs", "ujmy", "mh"}
	report := []string{"eveszfubew jebebqp iq eveszfubew eveszfubew iq daej eveszfubew q da", "ohfz zs ujmy egpakyk eveszfubew pffeq q qkmhvgxg kdgqq ipp", "cceierguau mh da eveszfubew m etwuedg ikeft egpakyk ltnibxljfi m", "km m iq rab inooo ujmy tlrdyu yqhn m xlkhebs", "q etwuedg m eveszfubew ixrfzwmb m jyltumdwt dacmewk odbllqdiq eveszfubew"}
	student_id := []int{643903773, 468275834, 993893529, 509587004, 61125507}
	k := 5

	fmt.Println(topStudents(positive_feedback, negative_feedback, report, student_id, k))
}
