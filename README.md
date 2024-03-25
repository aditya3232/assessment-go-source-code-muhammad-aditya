# Description

This is golang clean architecture.

## Architecture

![Clean Architecture](architecture.png)

1. Sistem eksternal melakukan permintaan (HTTP, gRPC, Messaging, dll) ke delivery, contoh permintaan dari sistem eksternal adalah request HTTP dari client, request gRPC dari client, atau baca pesan dari message broker
2. delivery akan mengakses model request untuk setiap request yang masuk, dan memanggil model response, untuk mengembalikan response data dari use case ke sistem eksternal
3. delivery akan memanggil use case, didalam use case berisi bisnis logic (transaksi, validasi, dll) 
4. use case membuat instance baru dari entity dan mengisi data entity dari request model
5. use case memanggil repository, untuk menyimpan data instance entity yg telah dibuat ke database
6. repository menggunakan data entity untuk melakukan operasi database
7. repository melakukan operasi basis data ke database
8. use case memberikan response yang sesuai berdasarkan model response
9. lalu jika ada data yang perlu dikirim ke sistem eksternal, maka use case akan memanggil gateway, dan gateway akan mengirimkan data ke sistem eksternal
10. gateway juga akan mengaksesl model untuk data yang dikirim ke sistem eksternal
11. gateway akan melakukan permintaan mengirimkan data ke sistem eksternal

## Tech Stack

- Golang : https://github.com/golang/go
- MySQL (Database) : https://github.com/mysql/mysql-server

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus
- Confluent Kafka Golang : https://github.com/segmentio/kafka-go

## Configuration

All configuration is in `config.json` file.

## API Spec

All API Spec or postman collection is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://root:root@tcp(localhost:3306)/assessment_go_source_code_muhammad_aditya?charset=utf8mb4&parseTime=True&loc=Local" -path db/migrations up
```

## Run Application

### Run unit test

```bash
go test -v ./test/
```

### Run web server

```bash
go run cmd/web/main.go
```
