package recursion

import "testing"

func BenchmarkNonsense05(b *testing.B) { benchNonsense(5, b) }
func BenchmarkNonsense10(b *testing.B) { benchNonsense(10, b) }
func BenchmarkNonsense15(b *testing.B) { benchNonsense(15, b) }
func BenchmarkNonsense20(b *testing.B) { benchNonsense(20, b) }
func BenchmarkNonsense25(b *testing.B) { benchNonsense(25, b) }
func BenchmarkNonsense30(b *testing.B) { benchNonsense(30, b) }
func BenchmarkNonsense35(b *testing.B) { benchNonsense(35, b) }
func BenchmarkNonsense40(b *testing.B) { benchNonsense(40, b) }
func BenchmarkNonsense45(b *testing.B) { benchNonsense(45, b) }

func benchNonsense(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Nonsense(1, uint(i))
	}
}

func BenchmarkNonsenseLoop05(b *testing.B) { benchNonsenseLoop(5, b) }
func BenchmarkNonsenseLoop10(b *testing.B) { benchNonsenseLoop(10, b) }
func BenchmarkNonsenseLoop15(b *testing.B) { benchNonsenseLoop(15, b) }
func BenchmarkNonsenseLoop20(b *testing.B) { benchNonsenseLoop(20, b) }
func BenchmarkNonsenseLoop25(b *testing.B) { benchNonsenseLoop(25, b) }
func BenchmarkNonsenseLoop30(b *testing.B) { benchNonsenseLoop(30, b) }
func BenchmarkNonsenseLoop35(b *testing.B) { benchNonsenseLoop(35, b) }
func BenchmarkNonsenseLoop40(b *testing.B) { benchNonsenseLoop(40, b) }
func BenchmarkNonsenseLoop45(b *testing.B) { benchNonsenseLoop(45, b) }

func benchNonsenseLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		NonsenseLoop(1, uint(i))
	}
}
