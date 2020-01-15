package recursion

import "testing"

func BenchmarkFib05(b *testing.B) { benchFib(5, b) }
func BenchmarkFib10(b *testing.B) { benchFib(10, b) }
func BenchmarkFib15(b *testing.B) { benchFib(15, b) }
func BenchmarkFib20(b *testing.B) { benchFib(20, b) }
func BenchmarkFib25(b *testing.B) { benchFib(25, b) }
func BenchmarkFib30(b *testing.B) { benchFib(30, b) }
func BenchmarkFib35(b *testing.B) { benchFib(35, b) }
func BenchmarkFib40(b *testing.B) { benchFib(40, b) }
func BenchmarkFib45(b *testing.B) { benchFib(45, b) }

func benchFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(uint(i))
	}
}

func BenchmarkFibLoop05(b *testing.B) { benchFibLoop(5, b) }
func BenchmarkFibLoop10(b *testing.B) { benchFibLoop(10, b) }
func BenchmarkFibLoop15(b *testing.B) { benchFibLoop(15, b) }
func BenchmarkFibLoop20(b *testing.B) { benchFibLoop(20, b) }
func BenchmarkFibLoop25(b *testing.B) { benchFibLoop(25, b) }
func BenchmarkFibLoop30(b *testing.B) { benchFibLoop(30, b) }
func BenchmarkFibLoop35(b *testing.B) { benchFibLoop(35, b) }
func BenchmarkFibLoop40(b *testing.B) { benchFibLoop(40, b) }
func BenchmarkFibLoop45(b *testing.B) { benchFibLoop(45, b) }

func benchFibLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibLoop(uint(i))
	}
}
