package summultiples

func SumMultiples(limit int, divisors ...int) (sum int) {
	lists := make([][]int, 0)
	for d := range divisors {
		list := make([]int, 0)
		if d != 0 {
			for mult := 1; (mult * d) < limit; mult++ {
				product := mult * d
				list = append(list, product)
			}
		}
		lists = append(lists, list)
	}
	set := make(map[int]bool)
	for _, list := range lists {
		for _, x := range list {
			if !set[x] {
				set[x] = true
			}
		}
	}
	for key := range set {
		sum += key
	}
	return
}
