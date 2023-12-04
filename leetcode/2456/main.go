package main

import "fmt"

type data struct {
	creator  string
	allViews int
	maxView  int
	maxIdes  string
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {

	theMap := make(map[string]data)
	for i := 0; i < len(creators); i++ {
		d := theMap[creators[i]]
		d.creator = creators[i]
		d.allViews += views[i]

		if len(d.maxIdes) == 0 {
			d.maxIdes = ids[i]
		}

		if views[i] > d.maxView || (views[i] == d.maxView && ids[i] < d.maxIdes) {
			d.maxView = views[i]
			d.maxIdes = ids[i]
		}
		theMap[creators[i]] = d
	}

	var maxViews int
	tmp := make([]data, 0, len(theMap))
	for _, d := range theMap {
		tmp = append(tmp, d)
		if d.allViews > maxViews {
			maxViews = d.allViews
		}
	}

	var answer [][]string
	for _, d := range tmp {
		if d.allViews == maxViews {
			answer = append(answer, []string{d.creator, d.maxIdes})
		}
	}

	return answer
}

func main() {
	creators := []string{"alice", "bob", "alice", "chris"}
	ids := []string{"one", "tow", "three", "four"}
	views := []int{5, 10, 5, 4}
	fmt.Println(mostPopularCreator(creators, ids, views))
}
