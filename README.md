# Guía 2

## Pilas y Colas

### Tipo Abstracto de Datos

En las carpetas `/stack` y `/queue` se encuentran dos implementaciones básicas de pilas y colas que tienen un problema: no encapsulan los datos y por lo tanto se pueden manipular sabiendo que ambos TAD están implementados sobre arreglos.

1. Se pide modificar los TAD `Stack` y `Queue` para que no se puedan manipular los datos y que la única forma de acceder a ellos sea a través de los métodos definidos.

2. Implementar nuevamente la cola (`queue`) pero usando internamente dos pilas (`stack`) para almacenar los datos.

   - `Enqueue`: agrega el elemento a la primera pila.
   - `Dequeue`: devuelve el tope de la segunda pila. Si la segunda pila está vacía, entonces desapila uno por uno todos los elementos de la primera pila y los apila en la segunda

Analizar el orden de cada uno de los métodos.

```go
type QueueS //TO-DO

func (q *QueueS) Enqueue(v any) {
	//TO-DO
}

func (q *QueueS) Dequeue() (any, error) {
	//TO-DO
}

func (q *QueueS) IsEmpty() bool {
	//TO-DO
}

func (q *QueueS) Front() (any, error) {
	//TO-DO
}

```

### Usos de pilas y colas

1. Escribir una función que que reciba una cadena de caracteres y devuelva la cadena invertida. Analizar el orden.

```go
func ReverseString(string)(string){
    //TO-DO
}
```

2. Escribir una función que verifique si una palabra es palíndromo (es decir que una cadena es igual a su inversa. Por ejemplo: las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.

```go
func Palindromo(string)(bool){
    //TO-DO
}
```

3. Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y devuelve `true` si está bien balanceada y `false` si está mal balanceada. Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.

```go
func Balanceo(string)(bool){
    //TO-DO
}
```

4. Escribir una función, tal que, dadas dos colas, construya una cola con el resultado de poner una a continuación de la otra. Por ejemplo: si `q1 = [1,2,3]` (con 1 al frente y 3 al final) y `q2 = [5,7]`, el resultado es `[1,2,3,5,7]` (con 1 al frente y 7 al final). Analizar el orden.

```go
func AddQueue(q1, q2 queue.Queue)(queue.Queue){
    //TO-DO
}
```
