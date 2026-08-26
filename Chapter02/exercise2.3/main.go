// Перепишите PopCount так, чтобы использовать цикл вместо одного выражения.
// Сравните производительность двух версий.
// (В Разделе 11.4 показано, как систематически сравнивать производительность различных реализаций.)
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var result int
	for i := range 8 {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}
