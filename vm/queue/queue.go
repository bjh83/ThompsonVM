package queue

import(
	"container/list"
)

type Queue struct {
	List *list.List
	Map map[int]bool
}

func New() *Queue {
	queue := &Queue{}
	queue.Map = make(map[int]bool)
	queue.List = list.New()
	queue.List.Init()
	return queue
}

func (queue *Queue) Push(pc int) {
	if !queue.Map[pc] {
		queue.Map[pc] = true
		queue.List.PushBack(pc)
	}
}

func (queue *Queue) Pop() int {
	ret := queue.List.Remove(queue.List.Front()).(int)
	delete(queue.Map, ret)
	return ret
}

func (queue *Queue) Len() int {
	return queue.List.Len()
}

