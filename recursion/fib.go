package recursion

func fib(n uint) uint {
	if n <= 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fibLoop(n uint) uint {
	f1 := uint(0)
	if n <= 1 {
		return f1
	}

	f2 := uint(1)
	if n == 2 {
		return f2
	}

	f3 := f2 + f1
	for i := uint(3); i < n; i++ {
		f1 = f2
		f2 = f3
		f3 = f2 + f1
	}

	return f3
}
