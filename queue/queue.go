package queue

import (
	"errors"
)

type Queue []any

func (q *Queue) Enqueue(v any) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head, nil
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Front() (any, error) {
	if len(*q) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (*q)[0], nil
}
