# Temario GO 2019

## Introdución

Temario practico creado con fines didacticos para aprender GO. 
Author:[Alexandr Campos --GitHub](https://github.com/alexandr28?tab=repositories)

![Golang](https://cdn-images-1.medium.com/max/2400/1*30aoNxlSnaYrLhBT0O1lzw.png)

## Primera Clase GO

### Variables

ejemplo de declaración de variables  de forma tradicional

```
package main
import "fmt"
func main() {
var nombre, apellido string
nombre = "Alexandr"
apellido = "Campos"
fmt.Println("usuario", nombre, apellido)
}
```

otra forma de declarar variables
```
func main() {
nombre, apellido:= "alexandr" "campos"
fmt.Println("usuario", nombre, apellido)
}
```

declarar constantes ejemplo

```
func main() {
const usuario = "Alexandr Campos"
fmt.Println("usuario:", usuario)
}
```

### Tipos de Datos

No se puede realizar operaciones artimeticas con diferentes tipos de datos excepto que se haga casting

```

var a int64 = 12121212
var b int8 = 88
//casting
c := a + int64(b)
fmt.Println("el resultado es:", c)
```

uso de comodinemos en Go es practico, asi lo demos ver en este ejemplo con %s que representa a una variable de tipo string para reemplazar, asi como identificar el tipo de variable por ejemplo

````
name := "Mariluz"
lastname := "Garro"
fmt.Printf("Hola %s %s como estas?\n", name, lastname)
fmt.Printf("lastname es de tipo: %T\\n", lastname)
fmt.Printf("C es de tipo: %T", c)
```