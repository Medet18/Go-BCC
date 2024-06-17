package problem1

func isChangeEnough(change []int, totalDue float32) bool {

	if len(change) != 4 {
		return false
	}

	totalAmount := float32(change[0])*0.25 + float32(change[1])*0.10 + float32(change[2])*0.05 + float32(change[3])*0.01

	return totalAmount >= totalDue
}
