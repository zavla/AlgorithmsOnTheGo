package main

import "math"

//BruteForceSackbag находит набор предметов из m так чтобы их масса была меньше limit,
//а стоимость максимально возможной.
//m - массы предметов, могут быть дробными, потому что решаем рюкзак перебором.
//c - стоиомости предметов
//limit - предел ограничение общей массы набора предметов.
//Решает задачу Рюкзака перебором.
func BruteForceSackbag[T ~float64](m []T, c []int, limit T) (T, []T) {
	n := len(m)
	_2n := int(math.Pow(2, float64(n)))
	maxS := T(0) //стоимость максимального набора
	maxNabor := 0
	//1010101 - взять\не взять. Всего возможных наборов 2**n.
	//получаем номер набора, из номера набора используем биты.
	for nabor := 1; nabor < _2n; nabor++ {
		s := T(0) //сумма стоимостей каждого набора
		w := T(0) //вес каждого набора
		//получаем бит для каждого предмета, 1й бит, 2й бит\предмет
		for item := 0; item < n; item++ {
			mask := 1 << item
			bit := nabor & mask
			if bit > 0 {
				//предмет в наборе выбран
				//добавим его значение к общ сумме набора
				w += m[item] //пусть нулевой бит означает берем ли нулевой элемент
				s += T(c[item])
				if w > limit {
					break //набор уже превысил лимит
				}
			}
		}
		if w <= limit && maxS < s {
			maxS = s
			maxNabor = nabor
		}
	}
	return maxS, getItemsFromNabor(maxNabor, m)
}
func getItemsFromNabor[T ~float64](nabor int, m []T) []T {
	ret := make([]T, 0)
	for i := 0; i < len(m); i++ {
		mask := 1 << i //нулевой бит это нулевой элемент
		if mask&nabor != 0 {
			ret = append(ret, m[i])
		}

	}
	return ret
}
