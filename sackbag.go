//Задача о рюкзаке.
package main

//какую максимальную стоимость можно вместить в рюкзак массой М если есть предметы массими m[i],i<=n, и стоимостью c[i] каждый.
//№грузика
//стоимость 	c 1 2 3 5
//масса 		m 1 4 6 3
// ---------
//на входе массив весов предметов m и массив с стоимости предметов.
//W - размер рюкзака в кг.
//W,m[i] используются как индексы к массиву выгод, поэтому они должны быть целыми.
//Если веса, или стоимости не целые то их надо самому домножить все на 10 например, чтобы вместо 23.7кг получить 237 кг.
//Т.о. при дробных весах сначала использовать scaleToInts(m, 10), W*=10
func sackbag[T ~int](W T, m []T, c []T) (T, []T) {
	//решим задачу сначала для одного грузика со всеми остатками свободного места. остатки свободного места в кг.
	//заполняем таблицу выгод от размещения первого грузика(грузиков) в остатках свободного места.
	//как выгоднее заполнить 1 кг свободного места если есть только один грузик. два грузика.
	//На десятом грузике мы уже в таблице видим что если мы его берем то заполненность рюкзака с остатком места под десятый
	// грузик чтобы его взять оптимальная такая-то.
	// Добавим десятый предмет - получим выгоду такую-то. Проверим что она не меньше
	// если десятый предмет не брать.
	// 			номера грузиков\предметов
	// 	табА		1	2	3	4	5	...
	// 		0кг
	// с  	1кг
	// в	2кг
	// о	3кг
	// б	4кг
	// о	5кг
	// д	.
	// н	.
	// о	.
	// е
	// место
	N := len(m) //для краткости записи это число предметов

	//создаем А=таблица выгод, в строках килограммы, в столбцах номера\индексы грузиков
	//шаг измерения массы 1 кг

	A := make([][]T, W+1) //таблица выгод. есть W строк 1кг, 2кг, 3кг, индекс будет сравниваться с кг.
	for i := 0; i < len(A); i++ {
		A[i] = make([]T, N) //создаем А, в строках килограммы=размеры мешков, в столбцах номера\индексы грузиков
	}

	//т.к. индекс kg сравнивается с реальным весом в кг то его надо его начать с 1, а не 0.
	for kg := T(1); kg <= W; kg++ { //для каждого кг, т.е. для каждого свободного места в мешке
		for j := 0; j < N; j++ { //для каждого грузика, j это номер грузика
			profit1 := T(0) //когда не взяли текущий грузик
			profit2 := T(0) //прибыль когда взяли текущий грузик
			if j > 0 {
				//возможна прибыль без этого товара, т.е. другими товарами заполнили эти килограммы
				profit1 = A[kg][j-1]
			}
			if kg >= m[j] { //если рассматриваем мешок больший чем вес грузика, то возможно поместить грузик
				//возможна прибыль с этим товаром но по этому же принципу заполненны килограммы без него.
				takeweight := m[j]
				profitWithoutWhatWeTake := T(0)
				if j > 0 {
					profitWithoutWhatWeTake = A[kg-takeweight][j-1] //профит меньшего мешка
				}
				profit2 = profitWithoutWhatWeTake + T(c[j]) //j-1 значит попробую взять меньший мешок без груза j плюс цену груза j
			}
			//из двух вариантов прибылей выбрать большую
			A[kg][j] = max(profit1, profit2) //либо без этого товара такое же число кг, либо с этим товаром но оптимально забитые предыдущие килограммы
		}
	}

	ret := A[W][N-1] //строки нумеруються до W включительно, а колонки до N-1

	//знаем макс выгоду, найдем обратным ходом то из каких грузиков она получилась
	items := []T{}
	kg := W
	j := N - 1
	for {
		profit := A[kg][j]
		if j == 0 {
			break
		}
		if profit == A[kg][j-1] { //значит грузик j не брался
			j--
			continue
		}

		items = append(items, m[j])
		kg = kg - m[j]
		j-- //если j уже взялся второй раз его не берем

	}
	//reverse items
	for i := 0; i < len(items)/2; i++ {
		items[i], items[len(items)-i-1] = items[len(items)-i-1], items[i]
	}
	return ret, items
}

func scaleToInts[T ~float64 | ~int](m []T, factor int) []int {
	//шаг измерения массы 1 кг
	if factor < 1 || factor > 100 {
		panic("bad scale factor")
	}
	ret := make([]int, len(m))
	//m[i] приводим к целым числам
	for i := range m {
		ret[i] = int(T(factor) * m[i]) // все массы превращаем в целые т.к. используем массу как индекс в таблице выгод.
		//в таблице выгод.
	}
	//например было 23.7 кг, станет 237 кг
	return ret

}

func scaleToFloats(m []int, factor float64) []float64 {
	ret := make([]float64, len(m))
	for i := range m {
		ret[i] = float64(m[i]) * factor
	}
	return ret
}
