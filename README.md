# Микросервис (gRPC) для работы с базой данных Postgres
Конфигурация находится в ```.\config\prod.yaml```
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

> [!IMPORTANT]
> Если не указать аргументы запуска, по умолчанию будет загружен конфиг по пути ```.\config\local.yaml```

# Protocol Buffers
Proto файл находится в папке ```.\protos\proto```
Для кодогенерации есть **Makefile** файл, переходим в папку ```.\protos\proto``` и в терминале пишем ```make 
generate_go``` - для генерации сервера и клиента на языке **Golang** или ```make generate_python``` - для генерации 
сервера и клиента на языке **Python**, 
сгенерированные файлы сами скопируются в необходимую папку (за исключением **generate_python**). На данный момент 
только для 
<b>Windows</b>

> [!IMPORTANT]
> Для использования моего **makefile** обязательно необходима утилита **[make](https://www.make.com/en)** и установленные версии **protoc** 
> под 
> необходимый язык (В моем случае это **[grpcio, grpcio-tools](https://grpc.io/docs/languages/python/quickstart/)** для языка **Python** и **[protoc-gen-go, 
> protoc-gen-go-grpc](https://grpc.io/docs/languages/go/quickstart/)** для языка 
> **Golang**). 
> Необходимые инструменты берем с официальной документации **[gRPC](https://grpc.io/)**

# Документация
С документацией по API можно ознакомиться [тут](https://github.com/PomoschDev/ServiceDatabase/blob/main/doc/docs.md).