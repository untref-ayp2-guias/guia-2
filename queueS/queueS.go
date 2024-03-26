package queueS

// Implementación de una cola genérica utilizando dos pilas
type QueueS[T any] struct {
	// Implementar
}

// Crea una nueva cola vacía.
func NewQueueS[T any]() *QueueS[T] {
	// Implementar
	return nil
}

// Agrega un elemento a la cola.
func (q *QueueS[T]) Enqueue(v T) {
	// Implementar
}

// Elimina y devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Dequeue() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// Devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Front() (T, error) {
	// Implementar
	var x T
	return x, nil
}

// Devuelve true si la cola está vacía.
func (q *QueueS[T]) IsEmpty() bool {
	// Implementar
	return false
}
