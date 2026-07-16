# Technical Requirements Document (TRD)

## Arquitectura y Decisiones de Diseño

### 1. Gestión de Base de Datos y ORM (Object-Relational Mapping)

**Decisión:** No utilizar ORMs completos (como GORM) en favor de utilizar la librería estándar `database/sql` junto con el driver `pgx` y gestión de migraciones con `pressly/goose`.

**Contexto:**
Inicialmente se había contemplado el uso de GORM para agilizar el mapeo objeto-relacional. Sin embargo, para mantener un control estricto sobre las consultas generadas, minimizar la sobrecarga en tiempo de ejecución y seguir una filosofía más alineada con la idiomaticidad de Go (simplicidad y transparencia), se optó por una abstracción menor.

**Detalles Técnicos:**
- **Driver de BD:** Se utiliza `github.com/jackc/pgx/v5/stdlib`, ya que `pgx` es actualmente el driver de PostgreSQL más rápido y robusto para Go.
- **Migraciones:** Se integra `github.com/pressly/goose/v3`. Las migraciones se escriben en puro SQL (`.sql`) dentro del directorio `internal/database/migrations` y se embeben en el binario usando `go:embed`. Esto garantiza que la base de datos se migra automáticamente a su última versión de esquema durante el inicio de la aplicación, sin depender de herramientas externas en los entornos de despliegue.
- **Mapeo:** El mapeo de filas a estructuras (structs) se realizará de forma manual o mediante librerías ligeras como `sqlx` (si fuera necesario en el futuro) para evitar el acoplamiento a patrones de diseño mágicos típicos de los ORMs.

**Consecuencias (Trade-offs):**
- *Pro:* Consultas SQL predecibles y optimizadas manualmente.
- *Pro:* Menos "magia" oculta, lo que facilita el debugging.
- *Contra:* Mayor verbosidad (boilerplate) al escribir las consultas CRUD y escanear (scan) las filas hacia las estructuras de Go.
