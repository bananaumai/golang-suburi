package recursion

func Nonsense(n uint, times uint) {
	if times == 0 {
		_ = n
		return
	}
	Nonsense(n+1, times-1)
}

func NonsenseLoop(n uint, times uint) {
	for i := uint(0); i < times; i++ {
	}
	_ = n
}
