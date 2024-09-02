# Проект предоставляет набор задач для сборки и запуска сервера и клиентского приложения с использованием Docker.
## Задачи определены в файле `Taskfile.yml` и могут быть выполнены с помощью команды `task`.
### Предварительные требования
1. Docker установлен на вашем компьютере
2. Установленный инструмент командной строки task (вы можете установить его отсюда [ссылка](https://taskfile.dev/#/installation)
## Команды Taskfile
### Порядок запуска с помощью task: !!!Важно!!! Перед запуском приложений  необходимо создать сеть docker(Настройка сети).
#### Сборка приложения
1. Собрать сервер:
```shell
task build-server
```
2. Собрать клиента:
```shell
task build-client
```
3. Собрать и сервер, клиента:
```shell
task build-all
```
#### Настройка сети
1. Создать сеть Docker:
```shell
task setup-network
```
#### Запуск приложения
1. Запустить сервер:
```shell
task run-server
```
2. Запустить клиента:
```shell
task run-client
```
3. Запустить и сервер, клиента:
```shell
task run-all
```
#### Очистка
1. Удалить Docker-образы и сеть:
```shell 
task cleanup
```

## Порядок запуска с помощью docker
1. Создаем docker-сеть:
```shell
       docker network create wow-network
```
2. Собираем сервер:
```shell 
      docker build -t wow-server -f Dockerfile.server .
```
3. Запускаем сервер:
```shell
      docker run -d --name wow-server --network wow-network -p 12345:12345 wow-server:latest
```
4. Собираем клиента:
```shell 
      docker build -t wow-client -f Dockerfile.client .
```
5. Запускаем клиента:
```shell
      docker run --network wow-network wow-client:latest
```