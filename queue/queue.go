package queue

import "errors"

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	data []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	q.data = append(q.data, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if len(q.data) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := (q.data)[0]
	q.data = (q.data)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if len(q.data) == 0 {
		return nil, errors.New("queue is empty")
	}
	return (q.data)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

/* type QueueS
//Todo

func (q *QueueS) Enqueue(v any) {
	//TODO
}

func (q *QueueS) Dequeue() (any, error) {
    //TODO
}

func (q *QueueS) IsEmpty() bool {
    //TODO
}

func (q *QueueS) Front() (any, error) {
    //TODO
} */
