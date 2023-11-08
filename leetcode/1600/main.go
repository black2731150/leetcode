package main

//用树完成的方法，时间复杂度超标，官方答案用map实现的

// type TreeNodeMan struct {
// 	name     string
// 	death    bool
// 	age      int
// 	Clidrens []*TreeNodeMan
// }

// type ThroneInheritance struct {
// 	curOrder *TreeNodeMan
// }

// func Constructor(kingName string) ThroneInheritance {
// 	return ThroneInheritance{
// 		curOrder: &TreeNodeMan{
// 			name:     kingName,
// 			death:    false,
// 			age:      0,
// 			Clidrens: []*TreeNodeMan{},
// 		},
// 	}
// }

// func (t *ThroneInheritance) Birth(parentName string, childName string) {
// 	if t.curOrder.name == parentName {
// 		t.curOrder.Clidrens = append(t.curOrder.Clidrens, &TreeNodeMan{
// 			name:     childName,
// 			death:    false,
// 			age:      t.curOrder.age,
// 			Clidrens: []*TreeNodeMan{},
// 		})
// 	} else {
// 		for _, tn := range t.curOrder.Clidrens {
// 			newT := ThroneInheritance{curOrder: tn}
// 			newT.Birth(parentName, childName)
// 		}
// 	}
// }

// func (t *ThroneInheritance) Death(name string) {
// 	if t.curOrder.name == name {
// 		t.curOrder.death = true
// 	} else {
// 		for _, tn := range t.curOrder.Clidrens {
// 			newT := ThroneInheritance{curOrder: tn}
// 			newT.Death(name)
// 		}
// 	}
// }

// func (t *ThroneInheritance) GetInheritanceOrder() []string {
// 	var answer *[]string = &[]string{}

// 	var sorting func(tn *TreeNodeMan, answer *[]string)
// 	sorting = func(tn *TreeNodeMan, answer *[]string) {
// 		if !tn.death {
// 			*answer = append(*answer, tn.name)
// 		}

// 		for _, tnm := range tn.Clidrens {
// 			sorting(tnm, answer)
// 		}
// 	}

// 	sorting(t.curOrder, answer)

// 	return *answer
// }

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) (t ThroneInheritance) {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (t *ThroneInheritance) Birth(parentName, childName string) {
	t.edges[parentName] = append(t.edges[parentName], childName)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() (ans []string) {
	var preorder func(string)
	preorder = func(name string) {
		if !t.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range t.edges[name] {
			preorder(childName)
		}
	}
	preorder(t.king)
	return
}
