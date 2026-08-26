// Выражение x&(x-1) сбрасывает самый правый ненулевой бит в x.
// Напишите версию PopCount, использующую этот факт, и оцените ее производительность.
package popcount

func PopCount(x uint64) int {
	var result uint64
	for x != 0 {
		result++
		x &= x - 1
	}
	return int(result)
}
