<h1 align="center">Guía 2</h1>

<h3 align="center">Pilas y Colas</h2>

<h4>Tipo Abstracto de Datos</h4>

En las carpetas stack y queue se encuentran dos implementaciones básicas de pilas y colas que tienen un problema. No encapsulan los datos y por lo tanto se pueden manipular sabiendo que ambos TAD están implementados sobre arreglos.

1. Se pide modificar los TAD **_Stack_** y **_Queue_** para que no se puedan manipular los datos y que la única forma de acceder a ellos sea a través de los métodos definidos.

2. Implementar nuevamente la cola pero usando internamente dos pilas para almacenar los datos.

    - **_Enqueue:_** agrega el elemento a la primera pila
    - **_Dequeue:_** devuelve el tope de la segunda pila. Si la segunda pila está vacía, entonces desapila uno por uno todos los elementos de la primera pila y los apila en la segunda

Analizar el orden de cada uno de los métodos.

<h4>Usos de pilas y colas</h4>

1. Escribir un programa que lea una secuencia de caracteres y los imprime en orden inverso. Analizar el orden.
2. Escribir un programa que verifique si una palabra es palíndromo (es decir que una cadena es igual a su inversa. Por ejemplo: las cadenas "1456541" y "145541" son palíndromos). Analizar el orden.
3. Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y devuelve true si está bien balanceada y false si está mal balanceada. Por ejemplo: [()]{}{[()()]()} debe devolver true, mientras que [(]) debe devolver false. Analizar el orden.
4. Escribir una función, tal que, dadas dos colas, construya una cola con el resultado de poner una a continuación de la otra. Por ejemplo: por ejemplo si q1 = [1,2,3] (con 1 al frente y 3 al final) y q2 = [5,7], el resultado es [1,2,3,5,7] (con 1 al frente y 7 al final). Analizar el orden.
