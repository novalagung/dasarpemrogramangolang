# Contoh Session dengan store Mongo DB

Siapkan terlebih dahulu mongo db server, lalu sesuaikan connection string pada `main.go` line `18`.

```go
session, err := mgo.Dial("localhost:27123")
```
