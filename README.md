### Логгер
Предоставляет интерфейс для многоуровневого логирования.

#### Уровни логирования:
* FATAL
* ERROR
* WARN
* DEBUG
* INFO

Логгер можно инициализировать как singleton в рамках проекта

Пример:
```go
package main
import (
    log "logger"
    "os"
)

func main() {
    logFile, err := os.Open("error.log")
    if err != nil {...}

    log.Init(
        log.Colorize(true),   // включить цветные сообщения
        log.Level(log.DEBUG), // уровень логирования
        log.Output(logFile),  // перенаправляет лог в файл, по уполчанию stdout
    )
    logger := log.GetLogger()
    logger.Error("some error")
}
```

Можно использовать логгер не как singletone, а в качестве самостоятельного экземпляра

Пример:
```go
package main
import (
    log "logger"
    "os"
)

func main() {
    logger := log.NewLogger(
        log.Colorize(true),
        log.Level(log.DEBUG),
    )
    
    logger.Debug("some debug")
}   
```

Вызов фатальных ошибок отключить нельзя, FATAL это минимальный возможный уровень логирования.
Фатальные ошибки прекращают выполнения программы и возвращают операционной системе код возврата 1.
```go
logger.Fatal("message")

// грубо говоря то же самое что

fmt.Println("message")
os.Exit(1)
```