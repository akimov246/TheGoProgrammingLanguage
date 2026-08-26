// Напишите версию PopCount, которая считает биты путем сдвига своего аргумента через 64 битовые позиции,
// проверяя самый правый бит каждый раз. Сравните ее производительность с табличной версией.
package popcount

func PopCount(x uint64) int {
	var result uint64
	for range 64 {
		result += x & 1
		x >>= 1
	}
	return int(result)
}
