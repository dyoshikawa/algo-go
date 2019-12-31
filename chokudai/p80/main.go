package p80

func Invoke(base int) []int {
	res := make([]int, 0)
	for n := 2; n < base; n++ {
		ok := true
		for i := 0; i < base; i++ {
			for j := 0; j < base; j++ {
				for k := 0; k < base; k++ {
					if (i+j*base+k*base*base)%n == 0 &&
						(i+j+k)%n != 0 {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			res = append(res, n)
		}
	}

	return res
}
