package utils

func IntMin(i1 int64, i2 int64) int64 {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func IntMax(i1 int64, i2 int64) int64 {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}

func IntAbs(i int64) int64 {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}
