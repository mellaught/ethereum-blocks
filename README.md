# Тестовое задание для EXMO

## Задача: Релизовать микросервис на golang со следующими функциями
1. Подписаться по сокетам к ноде Ethereum
2. Получать новые блоки с транзакциями и сохранять в памяти последнии 100 блоков
3. Сделать endpoint (HTTP/gRPC) для получения по диапазону блоков (N-M) или transaction id
4. Реализовать graceful shutdown сервиса
*Ограничения:* Минимизировать количество зависимостей
Примечания: логи, тесты, обработка ошибок и т.п. на усмотрение разработчика


# TORUN

- `go get github.com/mellaught/ethereum-blocks`
- `cd $GOPATH:/src/github.com/mellaught/ethereum-blocks/src`
- **update config.json.** Add your appID: wss://ropsten.infura.io/ws/v3/ + *{YOUR_APP_ID}*.
- `go run main.go`

## Usage

**1. By blocks range:**
- `http://localhost:8000/blocks/8976341-8976344`

**2. By transaction id:**
- `http://localhost:8000/blocks/0x8488f4d75348439b9fc23cd2066edd7e9b8e79516d3aa01f0676414966632b31`