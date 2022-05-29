//найти кратчайшие пути из одной вершины во все остальные.
//Алгоритм Дейкстры.

//D - вспомогат массив длин кратчайших путей до вершин на данном шаге
//инициализируем его макс целым.

//S - пройденные вершины для которых мы точно знаем длину кратчайшего пути.

//for{
//Выбираем из D вершину w с кратчайшим путем, такую которую еще не брали(которая не в S)
//Если ничего не выбрали - то конец.
//Обновляем D по формуле D[k]=min(D[k], D[w]+V[w][k]) может путь в k через выбранную вершину стал короче.
//к S добавляем w

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Graph map[rune]maprune
type maprune map[rune]int

func (o maprune) String() string {
	keys := make([]int, 0, len(o))
	for k := range o {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ret := strings.Builder{}
	ret.WriteString("map{")
	for _, v := range keys {
		ret.WriteRune(rune(int32(v)))
		ret.WriteString(":")
		ret.WriteString(fmt.Sprintf("%v, ", o[rune(v)]))

	}
	ret.WriteString("}")
	s := ret.String()
	return s
}

func (G Graph) String() string {
	keys := make([]int, 0, len(G))
	for k := range G {
		keys = append(keys, int(k))

	}
	sort.Ints(keys)

	ret := strings.Builder{}
	ret.WriteString("{")
	for _, v := range keys {
		ret.WriteRune(rune(v))
		ret.WriteString(": ")
		ret.WriteString(fmt.Sprintf("%v,\n", G[rune(v)]))

	}
	ret.WriteString("}")
	return ret.String()
}

func (G *Graph) load(from []struct {
	s, e rune
	cost int
}) {
	for k := range from {
		E := from[k]
		if _, ok := (*G)[E.s]; !ok {
			(*G)[E.s] = make(map[rune]int)
		}
		(*G)[E.s][E.e] = E.cost
	}
}
func (G Graph) Edge(s, e rune) int {

	ret := math.MaxInt64
	if _, ok := G[s][e]; ok {
		return G[s][e]
	}
	return ret
}
func (G *Graph) getminvert(D map[rune]int, notIn map[rune]bool) (rune, bool) {

	mincost := math.MaxInt64 //max int
	minvert := rune(0)
	for k := range D {
		if mincost > D[k] {
			_, ok := notIn[k] //в выборе минимума не участвуют пути к вершинам до которых мы уже знаем кратчайший путь
			if !ok {

				mincost = D[k]
				minvert = k
			}
		}
	}
	if minvert == rune(0) {
		return rune(0), true
	}
	return minvert, false
}

func (G *Graph) deikstra_shortest_path(start rune) map[rune]int {
	//инициализация D, это список _всех_ вершин графа
	D := make(maprune) //начальные значение в D это MaxInt64, что означает "нет пути"
	for k := range *G {
		//G это список списков, одна вершина может иметь дуги к некоторым вершинам.
		for k2 := range (*G)[k] {
			if _, ok := D[k2]; ok {
				continue
			}
			D[k2] = math.MaxInt64
		}
		D[k] = math.MaxInt64
	}
	D[start] = 0
	//помещаем дуги начальной вершины в D
	for k, kcost := range (*G)[start] {
		D[k] = kcost
	}

	S := make(map[rune]bool) //список вершин про которые мы точно знаем кратчайший путь к ним

	k := start //начинаем с вершины , она в цикле меняется
	S[k] = true
	for {
		minVersh, none := G.getminvert(D, S) //находим минимум из длин дуг, вернет вершину к которой эта дуга
		if none {
			break
		}
		for k := range D { //обновляем длины кратчайших путей через вершину minVersh
			if _, ok := S[k]; ok && k != minVersh {
				//если k уже в S то не обновляем длинну пути к ней, т.к. к ней уже не будет более короткого пути
				continue
			}
			//если дуги нет то G.Edge вернет MaxInt64
			if G.Edge(minVersh, k) != math.MaxInt64 && //нет пути из minVersh ->k
				D[k] > D[minVersh]+G.Edge(minVersh, k) {
				D[k] = D[minVersh] + G.Edge(minVersh, k)
			}
		}
		S[minVersh] = true //вершину minVersh помещаем в список просмотренных
	}
	return D //вернет кратчайшие пути до каждой вершины
}
