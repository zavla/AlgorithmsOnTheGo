package main

//пока i<j пропустить явно меньшие слева i++
//пока i<j пропустить явно большие справа j--
//обменять
//пока j<j пропустить равные опорному, i++
//вурнуть j
func sort_quick(a []int) {

	if len(a) <= 1 {
		return
	}
	m := partition(a)
	sort_quick(a[:m])
	sort_quick(a[m:])

}

func partition(a []int) int {
	m := len(a) / 2
	cent := a[m]
	i := 0
	j := len(a) - 1
	for i < j {
		for i < j && a[i] < cent {
			i++
		}
		for i < j && a[j] > cent {
			j--
		}
		a[i], a[j] = a[j], a[i]
		for i < j && a[i] == cent {
			i++
		}
		m = j
	}
	return m
}
