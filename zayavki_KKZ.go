//Задача "Заявки в киноконцертный зал".
//Сколько макс заявок мы можем расположить в рассписании зала.

package main

import (
	"fmt"
	"sort"
	"time"
)

type beginend struct {
	starts time.Time
	ends   time.Time
	err    error
}

func (v beginend) String() string {
	return fmt.Sprintf("%s - %s", v.starts, v.ends)
}
func NewBeginend(sb, se string) beginend {
	t1, err := time.Parse(time.Kitchen, sb)
	if err != nil {
		return beginend{}
	}
	t2, err := time.Parse(time.Kitchen, se)
	if err != nil {
		return beginend{}
	}
	return beginend{
		starts: t1,
		ends:   t2,
	}

}

type mostRecentEnd []beginend

func (v mostRecentEnd) Len() int {
	return len(v)
}
func (v mostRecentEnd) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v mostRecentEnd) Less(i, j int) bool {
	if v[i].ends.Before(v[j].ends) {
		return true
	}
	return false
}

func eventOK(zal, e beginend) bool {
	if zal.starts.After(e.starts) || zal.ends.Before(e.ends) {
		return false
	}
	return true

}

func maxZayavok(zal beginend, actions []beginend) (ret []beginend) {
	//отсортировать акции по дате окончания, удалить начинающиеся ранее Время1, выбрать следующую,
	//снова убрать все начинающиеся ранее, взять следующую самую ранее заканчивающуюся.
	sort.Sort(mostRecentEnd(actions))
	t1 := actions[0].ends
	iok := 0
	ret = []beginend{}

	for i := range actions {

		if eventOK(zal, actions[i]) {
			iok = i
			t1 = actions[i].ends
			ret = append(ret, actions[i])
			break
		}
		//событие не подходящее под рабочее время зала, из отсортированных

	}

	for i := iok; i < len(actions); i++ {
		if actions[i].starts.Before(t1) || !eventOK(zal, actions[i]) {
			continue //началось до события t1 пропускаем
		}
		t1 = actions[i].ends
		ret = append(ret, actions[i])
	}
	return
}
