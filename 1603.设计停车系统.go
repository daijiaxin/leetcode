package leetcode

/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */

// @lc code=start
type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big: big, medium: medium, small: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.big > 0 {
		this.big--
		return true
	}
	if carType == 2 && this.medium > 0 {
		this.medium--
		return true
	}
	if carType == 3 && this.small > 0 {
		this.small--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
