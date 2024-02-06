package main

type AnimalShelf struct {
	CatQueue []int
	DogQueue []int
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		CatQueue: []int{},
		DogQueue: []int{},
	}
}

func (a *AnimalShelf) Enqueue(animal []int) {
	bianhao, zhonglei := animal[0], animal[1]
	if zhonglei == 0 {
		a.CatQueue = append(a.CatQueue, bianhao)
	}

	if zhonglei == 1 {
		a.DogQueue = append(a.DogQueue, bianhao)
	}
}

func (a *AnimalShelf) DequeueAny() []int {
	if len(a.CatQueue) == 0 && len(a.DogQueue) == 0 {
		return []int{-1, -1}
	}

	if len(a.CatQueue) == 0 {
		bianhao := a.DogQueue[0]
		a.DogQueue = a.DogQueue[1:]
		return []int{bianhao, 1}
	}

	if len(a.DogQueue) == 0 {
		bianhao := a.CatQueue[0]
		a.CatQueue = a.CatQueue[1:]
		return []int{bianhao, 0}
	}

	if a.CatQueue[0] < a.DogQueue[0] {
		bianhao := a.CatQueue[0]
		a.CatQueue = a.CatQueue[1:]
		return []int{bianhao, 0}
	} else {
		bianhao := a.DogQueue[0]
		a.DogQueue = a.DogQueue[1:]
		return []int{bianhao, 1}
	}
}

func (a *AnimalShelf) DequeueDog() []int {
	if len(a.DogQueue) == 0 {
		return []int{-1, -1}
	} else {
		bianhao := a.DogQueue[0]
		a.DogQueue = a.DogQueue[1:]
		return []int{bianhao, 1}
	}
}

func (a *AnimalShelf) DequeueCat() []int {
	if len(a.CatQueue) == 0 {
		return []int{-1, -1}
	} else {
		bianhao := a.CatQueue[0]
		a.CatQueue = a.CatQueue[1:]
		return []int{bianhao, 0}
	}
}

func main() {
	obj := Constructor()
	obj.Enqueue([]int{1, 0})
	obj.DequeueAny()
	obj.DequeueDog()
	obj.DequeueCat()
}
