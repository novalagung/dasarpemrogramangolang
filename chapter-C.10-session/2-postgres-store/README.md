# Contoh Session dengan store Postgres SQL DB

Siapkan terlebih dahulu sebuah postgres server, lalu sesuaikan connection string pada `main.go` line `17`.

```go
connectionString := "postgres://novalagung:@127.0.0.1:5432/novalagung?sslmode=disable"
```

Buat mac user bisa menggunakan [Postgresapp](https://postgresapp.com/) untuk mempermudah set up postgres server.
