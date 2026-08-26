package popcount

import "testing"

func TestPopCount(t *testing.T) {
	if got := PopCount(1488); got != 5 {
		t.Fatalf("неверный результат: got %d, want %d", got, 5)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1488)
	}
}
