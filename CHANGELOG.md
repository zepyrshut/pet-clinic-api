# Lista de cambios
Todos los cambios realizados en este proyecto irán aquí.

[Versionado semántico](https://semver.org/spec/v2.0.0.html), [Mayor.Menor.Parche-Status.Commit]

## [0.1.3-beta.5](#) - 2022-05-02
### Añadido
- Obtener todas las mascotas en la agenda.

## [0.1.2-beta.4](https://github.com/zepyrshut/pet-clinic/commit/368d1f895c92792e7d9a5c5fd60cf9b4da46cdee) - 2022-04-28
### Cambiado
- Directorio _driver_ a _database_.

## [0.1.1-beta.3](https://github.com/zepyrshut/pet-clinic/commit/7c468b91320e028172e1a942071e8d6db5672b3f) - 2022-04-27
### Añadido
- Servicio personas (dueños de las mascotas).
- Servicio mascotas (vinculación con el propietario).
- Validación en todas las entidades.
- Internacionalización básica. Funciona, pero no tiene funcionalidad.
- Capa interceptora con funciones de sesión, CORS e internacionalización.
- DTO para persona y mascota.
- Repositorio persona. Funciones nueva persona, una persona y vincular mascota a la persona.
- Rutas a los nuevos servicios.
- Utilidades de pruebas.

### Cambiado
- Nombre del proyecto, ahora se llama _pet-clinic_.

## [0.0.1-beta.1](https://github.com/zepyrshut/pet-clinic/commit/f1c0374bd50d509adc3c4406dd19eee76ac1161c) - 2022-04-26
### Añadido
- Primera versión.