package util

// CalcPercentage 计算percentage相对于total的百分比
func CalcPercentage(percentage int, total int) int {
	return int(float64(percentage) / float64(total) * 100)
}
