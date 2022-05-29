//сколько разных стопок можно сложить из n ящиков типов 1 и 2 при том что нельзя ложить три ящика типа 2 подряд.
package main

func numberOfDifferentPermutationsOfBoxes(n int) int {
	a := make([]int, n+1)
	a[0] = 1
	a[1] = 2
	a[2] = 4
	a[3] = 7
	for i := 4; i <= n; i++ {

		a[i] = a[i-1] + a[i-2] + a[i-3]

	}
	return a[n]

	// 1
	// a[3]=7
	// +
	// 2
	// 2
	// 1
	// a[1]=2
	// +
	// 2
	// 1
	// a[2]=4

}
