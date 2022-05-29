package main

//возвращает цифру n разряда, для несуществующих разрядов возвращает 0
func getn[T ~int](N T, n int) int {
	rest := 0
	for i := 1; i <= n; i++ {

		rest = int(N % 10)
		N = N / 10
	}
	return rest
}

func sort_podschet10[T ~int](a []T, n int) []T {

	//вспомогательный массив со счетчиками показывающими с какого индекса начнутся например тройки
	c := [10]int{}

	ret := make([]T, len(a))

	for i := range a {
		//получаем значение разряда, счетчик по индексу полученного значения разряда ++
		r := getn(a[i], n)
		c[r]++
	}
	//превращаем массив счетчиков в массив смещений (накопленных итогов)
	total := 0
	for i := range c {
		t := c[i]
		c[i] = total
		total += t
	}
	//проходим еще раз по a и разбрасываем его по местам в соотв с с
	for i := range a {
		r := getn(a[i], n)
		ret[c[r]] = a[i]
		c[r]++

	}
	return ret
}

func sort_porazrad[T ~int](a []T) []T {

	//сначала найдем максимум чтобы знать сколько разрядов у самого большого числа
	if len(a) == 0 {
		return a
	}
	maxN := a[0]
	for i := 1; i < len(a); i++ {
		if maxN < a[i] {
			maxN = a[i]
		}
	}
	maxRazrad := getRazradCount(int(maxN))
	//будем сортировать по каждому разряду
	ret := sort_podschet10(a, 1)
	for i := 2; i <= maxRazrad; i++ {
		ret = sort_podschet10(ret, i)
	}
	return ret
}

func getRazradCount(N int) int {
	r := 0
	floor := N
	for floor != 0 {
		floor /= 10
		r++
	}
	return r
}
