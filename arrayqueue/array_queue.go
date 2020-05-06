package arrayqueue

import (
	"bytes"
	"errors"
	"fmt"
)

//用数组实现队列
//队列是先进先出的一种线性结构 FIFO
//进的为队尾
//出的为队首

//主要操作:
//入队
//出队
//查看队首元素
//取队长

type arrayQueue struct {
	data []interface{}
	size int
}

func New(capacity int) *arrayQueue {
	return &arrayQueue{
		data: make([]interface{}, capacity),
		size: 0,
	}
}

func (q *arrayQueue) GetSize() int {
	return q.size
}

func (q *arrayQueue) getCapacity() int {
	return cap(q.data)
}

func (q *arrayQueue) isEmpty() bool {
	return q.size == 0
}

func (q *arrayQueue) isFull() bool {
	return q.size == q.getCapacity()
}

//入队
func (q *arrayQueue) EnQueue(v interface{}) (int, error) {
	if q.isFull() {
		return 0, errors.New("队列已满，入队失败")
	}

	q.data[q.size] = v
	q.size++

	return q.size, nil
}

//出队
func (q *arrayQueue) DeQueue() (interface{}, error) {
	if q.isEmpty() {
		return 0, errors.New("队列为空，出队失败")
	}

	v := q.data[0]
	for i := 0; i < q.size-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.size--

	return v, nil
}

//查看队首元素
func (q *arrayQueue) GetFront() (interface{}, error) {
	if q.isEmpty() {
		return 0, errors.New("队列为空，没有头元素")
	}

	return q.data[0], nil
}

func (q *arrayQueue) String() string {
	var buf bytes.Buffer

	buf.WriteString("Queue: ")
	buf.WriteString("Front [")

	for i := 0; i < q.size; i++ {
		buf.WriteString(fmt.Sprint(q.data[i]))
		if i != q.size-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")

	return buf.String()
}
