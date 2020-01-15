package recursion

func AddOne(n uint, time uint) uint {
	if time == 0 {
		return n
	}
	return AddOne(n+1, time-1)
}

func AddOneLoop(n uint, time uint) uint {
	for i := uint(0); i < time; i++ {
		n += 1
	}
	return n
}
