package main

//Куча это структура которая бысто возвращает максимальный элемент.
//Этот файл сортировка кучей, но куча на основании связного списка.
//Особенность в том что надо хранить указатель на последний узел кучи, т.к. к нему есть частый доступ.
//

type treenode[T ~int | ~float64] struct {
	left, right *treenode[T]
	up          *treenode[T] //указатель на родителя, нужен чтобы понимать какой узел потомок: левый или правый.
	value       T            // data by value

}
type heap[T ~int] struct {
	root *treenode[T]
	last *treenode[T] //способность быстро обратиться к последнему узлу, иначе последний узел постоянно
	//приходится искать.
}

func (t *heap[T]) GetMax() T {
	return t.root.value
}

//для соблюдения свойства кучи как полного дерева будем добавлять в кучу на последнем уровне дерева
//заполняя его слева направо.
func (t *heap[T]) findNodeAtWhichToAdd() *treenode[T] {

	//используем BFS т.к. он основан на очереди/списке с началом и концом и
	//обходит дерево по уровням. Первый извлеченным из очереди который без одного или без двух потомком
	//и есть тот узел к которому добавить последний елемент.
	todo := make([]*treenode[T], 1)
	todo[0] = t.root
	if t.root == nil {

		return nil
	}
	v := t.root
	for i := 0; i < len(todo); i++ {
		v = todo[i]
		if v.left == nil || v.right == nil {
			return v
		}
		todo = append(todo, v.left, v.right)

	}
	return v

}

func (t *heap[T]) Add(el T) {
	//добавляет значение в конец кучи
	p := t.findNodeAtWhichToAdd()
	if p == nil {
		t.root = &treenode[T]{value: el}
		t.last = t.root
		return
	}
	if p.left == nil {
		p.left = &treenode[T]{value: el, up: p}
		t.last = p.left
		t.moveUp(p.left)
	} else {
		p.right = &treenode[T]{value: el, up: p}
		t.last = p.right
		t.moveUp(p.right)
	}

}

func (t *heap[T]) moveUp(n *treenode[T]) {
	if n == nil {
		return
	}
	p := n.up //p = parent
	if p == nil {
		return
	}
	for p.value < n.value {

		if n == t.last { //т.е. из позиции "последний элемент" двигается вверх
			t.last = p //новый последний элемент
		}
		t.swapPositions(p, n)
		p = n.up
		if p == nil {
			//this is the root
			t.root = n
			break
		}
	}
}

func (t *heap[T]) swapPositions(above, under *treenode[T]) {

	//начинаем с того узла что над верхним

	//потом два узла что под нижним

	//потом узел что под верхним на уровне нижнего

	//потом сами два узла которые меняем местами
	//всего 10 указателей

	//1-й указатель
	if above.up != nil {
		if above.up.left == above {
			above.up.left = under
		} else {
			above.up.right = under
		}
	}
	//2-й
	if under.left != nil {
		under.left.up = above
	}
	//3-й
	if under.right != nil {
		under.right.up = above
	}
	//4-й это тот что либо справа либо слева от above, такой же как under
	if above.right == under && above.left != nil {
		above.left.up = under
	}
	if above.left == under && above.right != nil {
		above.right.up = under
	}
	//5,6
	under.up = above.up
	above.up = under
	//7,8
	l, r := under.left, under.right //temp l r
	if above.right == under {       //under был справа у above
		under.right = above
		under.left = above.left
	} else {
		under.left = above
		under.right = above.right
	}
	//9,10
	above.left = l
	above.right = r

}

func (t *heap[T]) swapPointers(p *treenode[T], n *treenode[T]) {
	up := p.up
	left := p.left
	right := p.right
	p.up = n
	p.left = n.left
	p.right = n.right
	n.up = up
	n.left = left
	n.right = right
}

func (t *heap[T]) findPenultimate() *treenode[T] { //Penultimate = penult = last but one
	if t.last == nil || t.last.up == nil {
		return nil
	}
	if t.last == t.last.up.left {
		//если это левый потомок
		//то предпоследний узел найти сложно.
		//подняться до первого узла с левым потомком который не указывает на тот узел откуда мы пришли.
		wefrom := t.last
		for wefrom.up != nil && wefrom.up.left == wefrom {
			wefrom = wefrom.up
		}
		if wefrom.up != nil {
			wefrom = wefrom.up.left

		}
		stepdown := wefrom
		for stepdown.left != nil || stepdown.right != nil {
			if stepdown.right != nil { //приоритет идти вправо
				stepdown = stepdown.right
				continue
			}
			stepdown = stepdown.left //если не вправо то влево
		}
		return stepdown

	} else {
		//просто частынй случай
		//последний узел это правый потомок, а левый потомок это предпоследний
		return t.last.up.left
	}
}

func (t *heap[T]) RemoveMax() {
	//перенести в вершину последний узел, потом сдвинуть его вниз до его нужного места.
	if t.root == nil || t.last == nil {
		return //no-op
	}
	if t.last.up == nil {
		//если предыдущего нет
		//только когда root указывает на последний элемент
		if t.root != nil {
			t.root.left, t.root.right = nil, nil
		}
		t.last.up = nil
		t.last = nil
		t.root = nil
		return
	}
	newlast := t.findPenultimate()
	if t.last.up.right == t.last {
		//отключить родителя от последнего
		t.last.up.right = nil
	} else {
		t.last.up.left = nil
	}
	t.root.left.up = t.last
	t.root.right.up = t.last
	t.last.left = t.root.left
	t.last.right = t.root.right
	t.root = t.last

	t.moveDown(t.root)

	t.last = newlast
}

// moveDown - продвигает вниз узел пока значение в енм меньше большего из его потомков
func (t *heap[T]) moveDown(v *treenode[T]) {
	//выбрать того из двух с кем меняем местами v
	var under *treenode[T]
	for {
		if v.left != nil && v.right != nil {
			if v.left.value > v.right.value {
				under = v.left
			} else {
				under = v.right
			}

		} else if v.left != nil {
			under = v.left
		} else if v.right != nil {
			under = v.right
		} else { //дошли до листа
			break
		}
		if v.value > under.value {
			t.swapPositions(v, under)
			v = under
			continue
		}
		break
	}

	return
}
