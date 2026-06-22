# Go Code Conventions & Style Guide

Este documento establece las convenciones de código y buenas prácticas para el desarrollo en **Go (Golang)** dentro de nuestro equipo. El objetivo es mantener una base de código limpia, legible, consistente y eficiente.

Nos basamos en los estándares oficiales de la comunidad de Go:
* [Effective Go](https://go.dev/doc/effective_go)
* [Go Style Decisions](https://go.dev/wiki/CodeReviewComments)

---

## 1. Formateo y Estilo de Código

### Herramientas Obligatorias
No se discute el formateo manual. Todo el código debe pasar por las siguientes herramientas antes de ser enviado a revisión:
* **`gofmt` / `goimports`**: Para el espaciado, alineación e importaciones automáticas.
* **`golangci-lint`**: El linter oficial del proyecto. Asegúrate de ejecutarlo localmente.

### Organización de Importaciones
Las importaciones deben agruparse en tres bloques separados por una línea en blanco, en el siguiente orden:
1. Paquetes de la librería estándar de Go.
2. Paquetes de terceros (external dependencies).
3. Paquetes internos del proyecto.

```go
package main

import (
	"fmt"
	"net/http"

	"[github.com/google/uuid](https://github.com/google/uuid)"
	"go.uber.org/zap"

	"mi-proyecto/internal/config"
	"mi-proyecto/internal/user"
)

```

---

## 2. Convenciones de Nombres (Naming)

* **CamelCase**: En Go se usa `mixedCaps` o `CamelCase`, nunca `snake_case`.
* **Visibilidad**:
* Primera letra **Mayúscula**: Exportado (Público).
* Primera letra **Minúscula**: No exportado (Privado del paquete).


* **Acrónimos**: Los acrónimos deben mantener la misma capitalización (`userID` o `URLPath`, **no** `userId` ni `UrlPath`).
* **Variables cortas**: Para scopes pequeños, usa nombres cortos (ej. `i` para índices, `r` para un reader, `w` para un writer). Si la variable se usa en un scope amplio, sé más descriptivo.

---

## 3. Manejo de Errores

El manejo de errores en Go es explícito. No uses `panic` para el flujo de control normal; `panic` es solo para errores irrecuperables del sistema.

### Buenas Prácticas

* **Retorna el error siempre al final** de la firma de la función.
* **Envuelve los errores (Error Wrapping)** para dar contexto usando `%w`:
```go
if err != nil {
    return fmt.Errorf("user repository: failed to save user: %w", err)
}

```


* **Primero la salida rápida (Happy Path)**: Evita la anidación excesiva de `if`. Trata el error y sal de la función temprano.

```go
// BIEN: Estilo limpio e idiomático
user, err := repo.Find(id)
if err != nil {
    return nil, err
}
return user, nil

// MAL: Anidación innecesaria
user, err := repo.Find(id)
if err == nil {
    return user, nil
} else {
    return nil, err
}

```

---

## 4. Contexto (`context.Context`)

* El `ctx context.Context` debe ser **siempre el primer argumento** de una función que realice operaciones de I/O (base de datos, llamadas HTTP, colas, etc.).
* No guardes contextos dentro de una estructura (struct); pásalo explícitamente a los métodos.

```go
func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    // ...
}

```

---

## 5. Concurrencia

La concurrencia en Go es poderosa pero peligrosa. Sigue estas reglas:

* **No expongas canales** en las APIs públicas de tus paquetes a menos que sea estrictamente necesario. Prefiere funciones síncronas y deja que el llamador decida si usar concurrencia.
* **Evita fugas de goroutines (Goroutine leaks)**: Asegúrate siempre de saber cómo y cuándo va a terminar una goroutine.
* Usa el paquete `sync` (`sync.Mutex`, `sync.WaitGroup`) para sincronización simple y canales para orquestación de datos.
* Ejecuta el detector de condiciones de carrera en tus pruebas: `go test -race ./...`.

---

## 6. Estructura y Declaración de Variables

### Inicialización de Slices y Maps

* Para slices vacíos que se enviarán a un cliente (JSON), inicializa con `make` o un literal vacío para evitar que se renderice como `null`.
* Usa la asignación corta `:=` solo dentro de funciones.

```go
// Slice vacío listo para JSON (evita null en JSON)
users := []User{} 

// Slice optimizado si conoces el tamaño de antemano
items := make([]Item, 0, expectedSize)

```

### Punteros vs. Valores

* Usa receptores de puntero (`func (s *Struct) Method()`) si el método necesita modificar el estado de la estructura, o si la estructura es muy grande en memoria.
* Usa receptores de valor (`func (s Struct) Method()`) para estructuras pequeñas e inmutables.

---

## 7. Pruebas (Testing)

* Las pruebas deben residir en el mismo paquete que el código que prueban, usando el sufijo `_test.go`.
* Se prefiere el uso de **Table-Driven Tests** para probar múltiples casos con la misma lógica.

```go
func TestSuma(t *testing.T) {
    cases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positivo y positivo", 2, 3, 5},
        {"negativo y positivo", -1, 1, 0},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            res := Suma(tc.a, tc.b)
            if res != tc.expected {
                t.Errorf("Suma(%d, %d) = %d; se esperaba %d", tc.a, tc.b, res, tc.expected)
            }
        })
    }
}

```

---

## 8. Comentarios y Documentación

* Todo paquete y toda función pública (exportada) **debe** tener un comentario de documentación.
* El comentario debe empezar con el nombre del elemento que documenta.

```go
// UserService define las operaciones de negocio para los usuarios.
type UserService struct {
    // ...
}

```



---

### 💡 Un par de consejos adicionales para tu equipo:
1. **Automatiza el proceso:** No dependas de que la gente recuerde este archivo. Configura un workflow de GitHub Actions (o tu CI/CD) que corra `golangci-lint` en cada Pull Request. Si el linter falla, no se aprueba el código.
2. **Revisión en pares:** Usa este archivo como "árbitro" en los Code Reviews. Si alguien hace una observación sobre el formato o el estilo, la justificación debe estar respaldada por este documento o por *Effective Go*.

```