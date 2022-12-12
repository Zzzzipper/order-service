# grpc-postgres
Сервис точки продаж: grpc-gateway + postgresql

## Локальное тестирования

Запустить сначала postgres контейнер:

```bash
$ docker run --rm -d --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=mypass -e POSTGRES_DB=postgres postgres:13
d1a2eb0fb44da9c3488184f5296da28d1c7f88bd32bd4ec81fc254f006886b03
```

Затем запустить сервис:

```bash
$ POSTGRES_URL=postgresql://postgres:mypass@localhost:5432/postgres?sslmode=disable go run main.go
...
{"level":"info","msg":"Serving gRPC on [::]:8080"}
{"level":"info","msg":"Serving Web UI on http://0.0.0.0:8080"}
```

Зайти в браузере по ссылке http://0.0.0.0:8080 и используя автоматически сгенерированный web UI для сервиса проверить его работоспособность вместе с  базой данных.

![gRPCUI](./grpcui.png)

## Разработка

### Требования

* `go` > 1.16

### Как сделать изменения в proto

После того, как были сделаны изменения в proto или файле миграций, необходимо сделать регенерацию файлов:

```bash
$ make generate
```

Если есть необходимость изменить схему базы данных, то нужно добавить другой файл миграции.
Для этого используется тот же самый формат наименований, именуя новые файлы префиксом `2_`. Затем нужно выполнить
 `make generate` и инкрементировать счетчик миграция в `users/helpers.go`.
