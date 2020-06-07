package main

func wiggleMaxLength(nums []int) (maxLen int) {
	const (
		BEGIN = iota
		UP
		DOWN
	)
	// 初始化状态为BEGIN
	var STATE = BEGIN
	for i := range nums {
		i += 1
		switch STATE {
		case BEGIN:
			if nums[i] > nums[i-1] {
				STATE = UP
				maxLen++
			} else if nums[i] < nums[i-1] {
				STATE = DOWN
				maxLen++
			} else {
				STATE = BEGIN
			}
		case UP:
			if nums[i] >= nums[i-1] {
				continue
			} else {
				STATE = DOWN
				maxLen++
			}
		case DOWN:
			if nums[i] <= nums[i-1] {
				continue
			} else {
				STATE = UP
				maxLen++
			}
		}
	}
	return
}
