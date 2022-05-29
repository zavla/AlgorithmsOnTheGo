package main

func sort_insertion(a []int) {
	for i := 0; i < len(a)-1; i++ {
		t := a[i]
		j := i
		for j > 0 && a[j-1] > t {
			a[j] = a[j-1]
			j--
		}
		a[j] = t
	}
}
