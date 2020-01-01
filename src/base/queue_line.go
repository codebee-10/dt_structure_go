//队列设计排队系统
package main

import (
	q "../structs/queue"
	"fmt"
)

type User struct {
	id       int
	name     string
	position int
}

type LoginQueue struct {
	queue *q.SliceQueue
}

//新增队列元素
func (p *LoginQueue) AddQueueItem(u *User) {
	p.queue.Push(u)
	u.position = p.queue.Size()
}

//出队列
func (p *LoginQueue) PopQueue() {
	p.queue.Pop()
	p.UpdateItemPos()
}

//更新队列中元素位置信息
func (p *LoginQueue) UpdateItemPos() {
	for k, v := range p.queue.Slice {
		v.(*User).position = k + 1
	}
}

func (p *LoginQueue) DelQueueItem(u *User) {
	for k, v := range p.queue.Slice {
		if v.(*User).id == u.id {
			p.queue.Slice = append(p.queue.Slice[:k], p.queue.Slice[k+1:]...)
		}
	}
	p.UpdateItemPos()
}

func main() {
	sliceQueue := q.SliceQueue{Slice: make([]interface{}, 0)}
	que := &LoginQueue{queue: &sliceQueue}
	que.AddQueueItem(&User{
		id:       1,
		name:     "Ethan",
		position: 1,
	})

	que.AddQueueItem(&User{
		id:       2,
		name:     "Alicia",
		position: 2,
	})

	que.AddQueueItem(&User{
		id:       3,
		name:     "Luo",
		position: 3,
	})

	que.AddQueueItem(&User{
		id:       4,
		name:     "Luo Cheng",
		position: 4,
	})

	fmt.Println("原队列:")
	que.queue.PrintQueue(que.queue)
	fmt.Println("排队后队列:")
	//que.PopQueue()
	que.DelQueueItem(&User{
		id:       2,
		name:     "Alicia",
		position: 2,
	})
	que.queue.PrintQueue(que.queue)
}
