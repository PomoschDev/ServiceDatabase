# Микросервис для работы с базой данных
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

# Proto
Proto файл находится в папке ```.\protos\proto```
Для кодогенерации есть **Makefile** файл, переходим в папку ```.\protos\proto``` и в терминале пишем ```make 
generate_go``` - для генерации сервера и клиента на языке **Golang** или ```make generate_python``` - для генерации 
сервера и клиента на языке **Python**, 
сгенерированные файлы сами скопируются в необходимую папку (за исключением **generate_python**). На данный момент 
только для 
<b>Windows</b>

# Документация
С документацией по API можно ознакомиться [тут](https://github.com/PomoschDev/ServiceDatabase/blob/main/doc/docs.md).