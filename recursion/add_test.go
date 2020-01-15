package recursion

import "testing"

func Test(t *testing.T) {
	expected := uint(4)

	actuals := []uint{AddOne(1, 3), AddOneLoop(1, 3)}

	for _, actual := range actuals {
		if actual != expected {
			t.Errorf("expected: %d <=> actual: %d", expected, actual)
		}
	}
}

func BenchmarkAddOne05(b *testing.B) { benchAddOne(5, b) }
func BenchmarkAddOne10(b *testing.B) { benchAddOne(10, b) }
func BenchmarkAddOne15(b *testing.B) { benchAddOne(15, b) }
func BenchmarkAddOne20(b *testing.B) { benchAddOne(20, b) }
func BenchmarkAddOne25(b *testing.B) { benchAddOne(25, b) }
func BenchmarkAddOne30(b *testing.B) { benchAddOne(30, b) }
func BenchmarkAddOne35(b *testing.B) { benchAddOne(35, b) }
func BenchmarkAddOne40(b *testing.B) { benchAddOne(40, b) }
func BenchmarkAddOne45(b *testing.B) { benchAddOne(45, b) }

func benchAddOne(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddOne(1, uint(i))
	}
}

func BenchmarkAddOneLoop05(b *testing.B) { benchAddOneLoop(5, b) }
func BenchmarkAddOneLoop10(b *testing.B) { benchAddOneLoop(10, b) }
func BenchmarkAddOneLoop15(b *testing.B) { benchAddOneLoop(15, b) }
func BenchmarkAddOneLoop20(b *testing.B) { benchAddOneLoop(20, b) }
func BenchmarkAddOneLoop25(b *testing.B) { benchAddOneLoop(25, b) }
func BenchmarkAddOneLoop30(b *testing.B) { benchAddOneLoop(30, b) }
func BenchmarkAddOneLoop35(b *testing.B) { benchAddOneLoop(35, b) }
func BenchmarkAddOneLoop40(b *testing.B) { benchAddOneLoop(40, b) }
func BenchmarkAddOneLoop45(b *testing.B) { benchAddOneLoop(45, b) }

func benchAddOneLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		AddOneLoop(1, uint(i))
	}
}
