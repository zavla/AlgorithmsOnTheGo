package main

//шелл сортировка начинается как сортировка вставками.
//но шаг не 1 а переменный: от 4 до 1. k=k/2
//т.е. почти что единичку заменяем на k
func sort_shell(a []int) {

	for k := 4; k >= 1; k = k / 2 {

		for i := k; i < len(a); i++ {
			t := a[i]
			j := i
			for j-k >= 0 && a[j-k] > t {
				a[j] = a[j-k]
				j = j - k
			}
			a[j] = t
		}
	}

}
