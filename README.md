# Simulación de concurrencia: Carrera de goroutines
En este programa, se simulan dos patrones de concurrencia
* Select-Multiplexing
* Fan-out/Fan-in

Gráficamente se simula una carrera con 5 corredores, la carrera se puede simular con 2 patrones de concurrencia diferentes, los mencionados anteriormente.

## Instrucciones de uso
Para ejecutar la carrera con el patrón select-multiplexing, se puede ejecutar el programa sin ningún argumento, por defecto se ejecutará este patrón: `go run main.go` o `go run main.go --pattern=select-multiplexing`.

Y para ejecutar el patrón fan-out/fan-in, se puede ejecutar especificando el argumento `pattern`: `go run main.go --pattern=fan-out/fan-in`.

## Explicación
### Simulación en general
La simulación consiste en una carrera, conformada por 5 pistas y corredores, y por cuestiones de indeterminismo, cualquiera de los corredores puede llegar a la meta primero (cada corredor consta de una goroutine propia).
### Select-multiplexing
El patrón select-multiplexing (que se ejecuta por defecto) consiste en el uso de la directiva `select`, el cual espera la señal de varios canales, y reacciona ante la primera señal de un canal que reciba, es decir, este patrón de concurrencia en Go, reacciona ante la primera señal de alguno de los canales que reciba; este patrón es útil para manejar múltiples fuentes de eventos o interrupciones, como en servidores que gestionan diversas conexiones o señales. El programa genera 5 canales (channels), uno por cada goroutine, y esperará al primer corredor que llegue a la meta, imprimiendo el resultado, aún si el resto de corredores (goroutines) siguen corriendo (ejecutandose). El programa cierra a los 5 segundos de que se anuncia el ganador.
### Fan-out/Fan-in
El patrón fan-out/fan-in (usando el flag `--pattern=fan-out/fan-in`) consiste en dividir una tarea en múltiples subtareas, por un lado, la goroutine principal divide las subtareas y las reparte en múltiples goroutines secundarias, unidas por un canal, cuando las goroutines acaban su tarea, la mandan por el mismo canal, y la goroutine principal coordina los resultados y los une; este patrón es útil para distribuir tareas idénticas a múltiples goroutines y luego combinar los resultados, útil para llamadas API paralelas o procesamiento paralelo de archivos. El programa genera un hilo para todos los corredores (goroutines), la goroutine que gestiona el patrón espera a que todos "lleguen a la meta" (es decir, que todas las goroutines acaben su trabajo), conforme vayan mandando sus resultados por el canal, en ese mismo orden llenarán una lista (arreglo) que almacena las IDs de los corredores, e imprimirá una tabla de clasificación (cuando terminen todos). El programa cierra a los 5 segundos de anunciar la tabla de clasificación
### Mecanismos de sincronización
El simulador emplea un par de mecanismos de sincronización:
* **Mutex:** La información de los jugadores (que también implica su ubicación en pantalla) se guarda en un arreglo de jugadores en el atributo `GameManager.players`. En cada worker se va modificando la información que hay en este atributo, y para evitar condiciones de carrera se emplea un mutex (o semáforo binario) donde solo un corredor puede modificar ese atributo a la vez.
* **Context:** Se puede cancelar la simulación al usar la tecla `Esc`, y para cerrar todas las goroutines de manera controlada, se emplea un context. Se inicializa un context y este se pasa al juego como un atributo `GameManager.cancelContext`, y al presionar la tecla `Esc` se usa esta función, que termina con todas las goroutines de manera controlada, si es que no se hayan terminado antes.