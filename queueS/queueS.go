package queueS

import(
    "github.com/untref-ayp2/stack"
)

// Implementación de una cola genérica utilizando dos pilas
type QueueS[T any] struct {
   // Implementar
}

// Crea una nueva cola vacía. 
func NewQueueS[T any]() *QueueS[T] {
    // Implementar
}

// Agrega un elemento a la cola. 
func (q *QueueS[T]) Enqueue(v T) {
   // Implementar
}

// Elimina y devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Dequeue() (any, error) {
   // Implementar
}

// Devuelve el elemento al frente de la cola.
func (q *QueueS[T]) Front() (any, error){
    // Implementar
}

// Devuelve true si la cola está vacía.
func (q *QueueS[T]) IsEmpty() bool{
   // Implementar
}

