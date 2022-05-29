package main

func merge(a []int) (ret []int) {

	if len(a) == 1 {
		return a
	}
	ret = make([]int, len(a))

	m := len(a) / 2

	b := merge(a[:m])
	c := merge(a[m:])

	ib := 0 // index for b
	ic := 0 //
	r := 0  //

	//пока оба индекса меньше соотв массивов копируем в ret оба массива, берем меньший элемент
	for ib < len(b) && ic < len(c) {
		//индекс для б меньше и индекс для ц меньше длины
		if b[ib] < c[ic] {
			ret[r] = b[ib]
			ib++

		} else {
			ret[r] = c[ic]
			ic++
		}
		r++
	}
	copy(ret[r:], b[ib:])
	copy(ret[r:], c[ic:])
	// for ib < len(b) {
	// 	ret[r] = b[ib]
	// 	r++
	// 	ib++
	// }
	// for ic < len(c) {
	// 	ret[r] = c[ic]
	// 	r++
	// 	ic++
	// }
	return ret
}
