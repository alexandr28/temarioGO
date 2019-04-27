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

```
name := "Mariluz"
lastname := "Garro"
fmt.Printf("Hola %s %s como estas?\n", name, lastname)
fmt.Printf("lastname es de tipo: %T\\n", lastname)
fmt.Printf("C es de tipo: %T", c)
```
### Condicionales
Existen dos formas de declarar las condiciones y son en simples y anidadas
ejemplo de Simples
```
if 5 > 10 {
    fmt.Println("es verdadero")
}
fmt.Println("fin del ciclo")
```
ejemplo anidada
```
if edad, nombre := 17, "Alexandr"; edad < 14 {
    fmt.Println(nombre, "es un niño")
} else if edad < 18 {
    fmt.Println(nombre, "es un adolecente")
} else {
    fmt.Println(nombre, "es un adulto")
}
```
### Switch
Switch en go tiene una forma particular de trabajar, que veremos acontinuación en el siguiente ejemplo
```
var a uint8 = 4
switch a {
case 1:
    // fallthrough -> sigue evaluando, la siguiente expresión
    fallthrough
case 2:
    fmt.Println("ganaste un viaje")
case 3:
    fallthrough
case 4:
    fmt.Println("ganas una camioneta 4x4")
default:
    fmt.Println("sigue intentando")
}
```
 ejemplo de switch usando condicionales en cada Case
 ```
switch x := 7; {
case x >= 0 && x <= 5:
    fmt.Println("estas entre semana")
case x >= 6 && x <= 7:
    fmt.Println("estas en fin de semana")
default:
    fmt.Println("no es un dia valido")
} 
 ```