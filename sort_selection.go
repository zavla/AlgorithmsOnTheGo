// сортировка выбором.
// Selection sort.
package main

//слева от индекса i отсортированная часть, справа не отсортированная
func sort_vibor(a []string) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}
