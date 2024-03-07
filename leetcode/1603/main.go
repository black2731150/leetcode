package main

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if p.Big > 0 {
			p.Big--
			return true
		} else {
			return false
		}
	case 2:
		if p.Medium > 0 {
			p.Medium--
			return true
		} else {
			return false
		}
	case 3:
		if p.Small > 0 {
			p.Small--
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
