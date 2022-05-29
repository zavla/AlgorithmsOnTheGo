package main

func sort_bubble(a []int) {
	for i := 1; i < len(a); i++ {
		wasswap := false
		for j := len(a) - 1; j >= i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j] //swap
				wasswap = true
			}
		}
		if !wasswap {
			return
		}
	}
}
