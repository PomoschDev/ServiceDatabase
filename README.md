# Микросервис (gRPC) для работы с базой данных Postgres
Конфигурация находится в ```./config/prod.yaml```
```yaml
    env: "prod" #Прод конфиг
    grpc: #Конфигурация сервера
      port: 44044 #Порт который будет прослушивать gRPC сервер
      timeout: 5s #Таймаут на запрос
    database: #Конфигурация базы данных
      host: localhost #Хост
      user: postgres #Пользователь БД
      password: postgres #Пароль БД
      name: postgres #Имя базы данных
      port: 3050 #Порт базы данных
      sslmode: disable #Использовать ли SSL сертификат (enable - включить)
      migrations: true #Применять миграции при запуске
```

## Запуск
Присутствует запуск с аргументами: ```server --config=./config/prod.yaml```

## Docker
В корне присутствует **Dockerfile** и **docker-compose.yml**. В случае клонирования репозитория, достаточно 
запустить команду ```docker compose up -d``` для разворачивания проекта. По умолчанию будет скопирована конфигурация 
из папки ```./config/prod.yaml``` и скопирована в контейнер как ```local.yaml```, сервис запустится с аргументом 
```--config=local.yaml```
### Docker без клонирования репозитория
В папке **```docker```**, так же присутствует **Dockerfile** и **docker-compose.yml**, данные файлы предназначены 
для того, что бы не клонировать весь репозиторий, сервис склонируется прямиков в контейнере, рядом с docker-файлами 
обязательно должен лежать файл конфигурации ```prod.yaml```, в следствии он будет скопирован в контейнере под именем 
```local.yaml``` 
для 
запуска сервиса с аргументом ```--config=local.yaml```. **Рекомендуется использовать именно этот вариант 
развертывания сервиса, так при развертывании мы всегда будем получать актуальную версию сервиса.**

> [!IMPORTANT]
> Если не указать аргументы запуска, по умолчанию будет загружен конфиг по пути ```./config/local.yaml```, если 
> конфигурация по данному пути не будет найдена, то сервис попытается загрузить конфигурацию по пути ```./local.
> yaml```, в случае если конфигурация так и не будет найдена, сервис завершится с ошибкой.

# Protocol Buffers
Proto файл находится в папке ```./protos/proto```
Для кодогенерации есть **Makefile** файл, переходим в папку ```./protos/proto``` и в терминале пишем ```make 
gen_go``` *(для систем Linux: ```make linux_gen_go```)* - для генерации сервера и клиента на языке **Golang** или ```make 
gen_python``` - 
для 
генерации 
сервера и клиента на языке **Python**, 
сгенерированные файлы сами скопируются в необходимую папку (за исключением **gen_python**).

> [!IMPORTANT]
> Для использования моего **makefile** обязательно необходима утилита **[make](https://www.make.com/en)** и установленные версии **protoc** 
> под 
> необходимый язык (В моем случае это **[grpcio, grpcio-tools](https://grpc.io/docs/languages/python/quickstart/)** для языка **Python** и **[protoc-gen-go, 
> protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/)** для языка 
> **Golang**). 
> Необходимые инструменты берем с официальной документации **[gRPC](https://grpc.io/)**

# Документация
С документацией по API можно ознакомиться [тут](https://github.com/PomoschDev/ServiceDatabase/blob/main/doc/docs.md).