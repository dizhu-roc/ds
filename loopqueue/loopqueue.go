package loopqueue

import (
	"bytes"
	"errors"
	"fmt"
)

//循环队列

//关键点: 取模算法
//定义 front 指向队列的第一个元素
//定义 rear  指向队列的最后一个元素的后一个位置
//多余空闲一个空间，用来判断队列的临界点

//关键指标
//队空: rear == front
//队满: (rear+1)%maxSize = front
//判断有效的数据个数: (rear+maxSize - front)%maxSize

type loopQueue struct {
	data        []interface{}
	size        int
	front, rear int
}

func New(capacity int) *loopQueue {
	return &loopQueue{
		data:  make([]interface{}, capacity),
		size:  0,
		front: 0,
		rear:  0,
	}
}

func (q *loopQueue) GetSize() int {
	return q.size
}

func (q *loopQueue) getCapacity() int {
	return cap(q.data)
}

func (q *loopQueue) isEmpty() bool {
	return q.size == 0
}

func (q *loopQueue) isFull() bool {
	return (q.rear+1)%q.getCapacity() == q.front
}

func (q *loopQueue) EnQueue(v interface{}) (int, error) {
	if q.isFull() {
		return 0, errors.New("队列已满,入队失败")
	}

	q.data[q.rear] = v
	q.rear = (q.rear + 1) % q.getCapacity()
	q.size++

	return q.size, nil
}

func (q *loopQueue) DeQueue() (interface{}, error) {
	if q.isEmpty() {
		return 0, errors.New("队列为空，出队失败")
	}

	v := q.data[q.front]
	q.front = (q.front + 1) % q.getCapacity()
	q.size--

	return v, nil
}

func (q *loopQueue) GetFront() (interface{}, error) {
	if q.isEmpty() {
		return 0, errors.New("队列为空，出队失败")
	}

	return q.data[q.front], nil
}

func (q *loopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("LoopQueue: ")
	buffer.WriteString("Front [")
	for i := 0; i < q.size; i++ {
		buffer.WriteString(fmt.Sprint(q.data[(q.front+i)%q.GetSize()]))
		if (q.front+i)%q.getCapacity() != q.rear {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
