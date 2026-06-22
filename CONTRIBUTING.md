# Contributing to [Nombre del Proyecto]

¡Gracias por tu interés en colaborar! Este proyecto se mantiene gracias al esfuerzo de personas como tú. Para asegurar una experiencia de colaboración fluida y eficiente, por favor lee y sigue estas directrices antes de enviar tu contribución.

## 1. Código de Conducta

Al participar en este proyecto, te comprometes a mantener un ambiente respetuoso, inclusivo y libre de acoso.


## 2. ¿Cómo puedo contribuir?

Puedes colaborar de muchas formas, no solo escribiendo código:
* **Reportar Bugs:** Si encuentras un error, avísanos.
* **Sugerir Características:** Propón nuevas ideas o mejoras.
* **Mejorar la Documentación:** Corrección de textos, ejemplos claros o traducciones.
* **Revisar Pull Requests:** Ayuda a otros aprobando o comentando en sus propuestas.


## 3. Reportar un Bug o Sugerir una Mejora

Antes de abrir un nuevo *Issue*, por favor **busca en los existentes** para asegurarte de que nadie haya reportado lo mismo.

Si vas a abrir un Issue, asegúrate de incluir:
* **Un título claro y descriptivo.**
* **Pasos para reproducir** el problema (si es un bug).
* **Comportamiento esperado vs. Comportamiento real.**
* **Entorno:** Sistema operativo, versión de lenguaje/herramientas, etc.
* **Capturas de pantalla o logs** si aplica.



## 4. El Proceso de Desarrollo (Workflow)

Si quieres agregar código o arreglar un bug, sigue este flujo de trabajo para evitar que tu trabajo choque con el de otros:

1. **Busca o crea un Issue:** Asegúrate de que el trabajo esté aprobado o discutido en un Issue antes de programar.
2. **Haz un Fork** del repositorio (si es un proyecto público) o clónalo localmente.
3. **Crea una rama (branch) descriptiva** desde la rama principal (`main` o `develop`):
   * Estructura recomendada: `feature/nombre-de-la-mejora` o `fix/nombre-del-bug`.
4. **Escribe tu código** siguiendo las guías de estilo del proyecto (`CODE_CONVENTIONS.md`).
5. **Asegúrate de que las pruebas pasen** localmente antes de subir tus cambios.


## 5. Convención de Commits

Para mantener un historial limpio y legible, utilizamos la convención de [Conventional Commits](https://www.conventionalcommits.org/). Tus mensajes de commit deben seguir esta estructura:

`tipo(alcance opcional): descripción corta en minúsculas`

### Tipos permitidos:
* **`feat`**: Una nueva característica para el usuario.
* **`fix`**: Corrección de un bug.
* **`docs`**: Cambios exclusivamente en la documentación.
* **`style`**: Cambios que no afectan el significado del código (espacios, formateo, punto y coma perdidos).
* **`refactor`**: Un cambio de código que no corrige un error ni añade una característica.
* **`test`**: Añadir o corregir pruebas existentes.

*Ejemplo:* `feat(auth): add google oauth2 login integration`



## 6. Envío de Pull Requests (PR)

Cuando tu rama esté lista, abre un Pull Request apuntando a la rama `main` (o la definida por el equipo) y asegúrate de cumplir con el siguiente Checklist:

* [ ] El PR está vinculado al Issue correspondiente (ej. *Closes #123*).
* [ ] El código pasa los linters y formateadores automáticos.
* [ ] Se han añadido o actualizado las pruebas unitarias correspondientes.
* [ ] La documentación ha sido actualizada si es necesario.
* [ ] El título del PR es descriptivo y sigue la convención de commits.

### Proceso de Revisión
* Al menos **[1 o 2]** mantenedor(es) deben revisar y aprobar tu código.
* Si el CI (Integración Continua) falla, el PR no se fusionará hasta que se resuelva el error.

¡Gracias de nuevo por hacer que este proyecto sea mejor para todos! 🚀