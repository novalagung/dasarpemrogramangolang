# C.39 Message Broker Apache Kafka
Pada chapter kali ini kita akan belajar bagaimana cara menggunakan apache kafka di golang secara sederhana, dari membuat koneksi apache kafka, membuat service worker consumer, dan cara komunikasi antara worker dan consumer.
## C.39.1 Apa itu Apache Kafka ?
Apache kafka adalah sebuah platform publish-subscribe streaming yang digunakan untuk mengatasi masalah pengolahan data secara real-time. Pada dasarnya Kafka bertujuan untuk mengirim pesan dari producer ke consumer yang akan diterima pesan dari producer tersebut.

Kelebihanan Kafka daripada platform lain seperti RabbitMQ adalah kecepatannya dalam mengirim data. Di dalam Kafka, setiap pesan hanya bisa ditambahkan saja jadi saat pesan sudah dikirim tidak dapat di edit atau dihapus. Selain itu juga, kelebihan dari Kafka menawarkan durabilitas, dan memiliki fitur toleransi kesalahan. Ini juga membantu untuk memastikan data stream secara skala yang besar dapat di manajemen secara efisien dan dengan latensi yang kecil.

## C.39.2 Persiapan
Kita siapkan dulu apache kafka di lokal komputer kita, untuk instalasi Kafka kalian juga bisa langsung install di sistem operasi komputer masing-masing.Namun dalam praktek ini kita menggunakan docker untuk install apache kafka.
### • Instalasi Apache Kafka
Jalankan command ini untuk mengambil dockerfile Kafka:
```
curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-kafka/master/docker-compose.yml > docker-compose.yml
```
Setelah dockerfile sudah ditambahkan, kita akan mengubah salah satu code ini.

Sebelum:
```
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://:9092
```
Setelah:
```
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
```
Jalankan command ini untuk membuat kontainer baru dan menjalankan apache kafka:
```bash
docker-compose up -d
```

## C.39.3 Praktek
Apache kafka sudah berjalan di komputer kita, setelah itu kita buat struktur folder seperti berikut ini:
```bash
.
│    # Folder connection berisi file konfigurasi dari Apache Kafka.
├── connection
│   └── ...
│    # Folder consumer berisi proses aplikasi digunakan untuk subcription ke satu atau lebih topik dan mengolah data-data dari topik tersebut.
├── consumer
│   └── ...
│    # Folder producer berisi proses atau sistem yang dapat mempublikasikan data ke suatu topik.
├── producer
│   └── ...
│    # File berisi depedencies seperti library kafka.
└── go.mod
```
Setelah struktur sudah dibuat, kita akan mengunduh library Apache Kafka golang.
```
go get github.com/IBM/sarama
```
### C.39.3.1 Koneksi Consumer dan Worker
Kita membuat 1 file bernama ```connection.go```, didalam file ini berisi konfigurasi koneksi untuk apache kafka. Tambahkan 1 fungsi koneksi Consumer di folder ```connection```  yang berfungsi untuk menerima data store(```Topic```) dari worker.
- #### File connection.go
```go
package connection

import "github.com/IBM/sarama"

func ConnectToConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ConnectToProducer(urls []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 6

	conn, err := sarama.NewSyncProducer(urls, config)
	if err != nil {
		return nil, err
	}

	return conn, err
}
```
Seperti yang terlihat di atas, kita menambahkan 1 fungsi lagi untuk koneksi Worker digunakan mengirimkan sebuah ``` Topic ``` yang berisi data ke Consumer.
### C.39.3.2 Mempersiapkan Consumer Kafka
Sekarang kita membuat Consumer, buat 2 file ```consumer.go``` dan ```main.go``` didalam folder ```consumer```.
Tambahkan code seperti dibawah ini:
- #### File consumer.go
```go
package main

import (
	"tutorial-go-kafka/connection"

	"github.com/IBM/sarama"
)

func PullFromProducer(topic string) (sarama.Consumer, sarama.PartitionConsumer) {

	worker, err := connection.ConnectToConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	return worker, consumer
}
```
File ```consumer.go``` secara sederhana berisi code untuk melakukan konsumsi data Message dari Worker. Di file ini juga kita menambahkan Topic yang bernama ```Pizza```, Topic ini nanti akan kita gunakan untuk menerima message berupa data.
pada code dibawah ini, berfungsi untuk meng-konsumsi data message dengan parameter nama Topic, partisi, dan offset consumer.
```go
consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
```
- #### File main.go
Buat main program untuk Consumer seperti dibawah ini:
```go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Nama Topic
	topic := "Pizza"

	worker, consumer := PullFromProducer(topic)

	log.Println("Consumer started")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	messageCount := 0

	doneChan := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				messageCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", messageCount, string(msg.Topic), string(msg.Value))
			case <-sigChan:
				fmt.Println("Interrupt is detected")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	fmt.Println("Processed", messageCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
```
Pada file ini berisi main program untuk menjalankan Consumer. Nantinya apabila ada data Message dikirim, maka code dibawah ini akan tampil log berisi total data Message dan data seperti apa yang dikirim.
```go
fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", messageCount, string(msg.Topic), string(msg.Value))
 ```

### C.39.3.3 Mempersiapkan Worker Kafka
Kita buat 2 file ```producer.go``` dan ```main.go``` didalam folder ```producer```. Nantinya pada main program Worker akan bertindak sebagai pengirim data Message.
- #### File producer.go
```go
package main

import (
	"fmt"
	"tutorial-go-kafka/connection"

	"github.com/IBM/sarama"
)

func PushPizzaQueue(topic string, message []byte) error {
	urls := []string{"localhost:9092"}
	producer, err := connection.ConnectToProducer(urls)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
```
Pada file diatas code hampir sama dengan file ```consumer.go``` namun hanya saja disini berisi fungsi untuk mengirimkan data Message.
```go
msg := sarama.ProducerMessage{
	Topic: topic,
	Value: sarama.StringEncoder(message),
}

partition, offset, err := producer.SendMessage(&msg)
if err != nil {
	return err
}
```
Barisan code diatas berisi mengirimkan data dan juga mengirimkan sebuah Topic yang bernama ```Pizza```. Parameter message data menggunakan tipe data ```[]byte```. Pada praktek ini kita akan mengirimkan data berupa json.
- #### File main.go
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pizza struct {
	Slice int `form:"slice" json:"slice"`
	Price int `form:"price" json:"price"`
}

func main() {
	http.HandleFunc("/sendpizza", func(w http.ResponseWriter, r *http.Request) {
		pizza := new(Pizza)
		dec := json.NewDecoder(r.Body)

		err := dec.Decode(&pizza)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("[ERROR]:" + err.Error())
		}
		pizzaBytes, err := json.Marshal(pizza)
		err = PushPizzaQueue("Pizza", pizzaBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("[ERROR]:" + err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(pizza)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
```
Pada main program Worker kita membuat sebuah Http Server dengan url ```/sendpizza```. Url ini berisi request data berupa json dan dari data tersebut akan dikirimkan ke server apache Kafka local lalu akan diterima oleh Consumer.
## C.39.4 Testing Aplikasi
Jalankan service Consumer dan Worker di folder ```consumer``` dan ```producer``` sudah dibuat dengan command berikut:
```
go run .
```
Kita akan melakukan testing mengirim 1 data Message dari Worker dengan menembak api ```/sendpizza```. Buat 1 request payload dengan text dan value seperti dibawah ini:
```json
{
    "slice": 12,
    "price": 1000
}
```
Maka apabila berhasil maka service Producer akan muncul log seperti ini:
![](images/producer.png)
Gambar diatas menandakan bahwa message sudah dikirim dengan 3 kali pengiriman karena sebelumnya penulis sudah testing dengan mengirim 2 buah message.

Seperti yang kita lihat gambar dibawah ini adalah log dari service Consumer. Data message dari service Producer telah diterima oleh service Consumer.
![](images/consumer.png)

---
 - [IBM Sarama](https://github.com/IBM/sarama), MIT license
---
<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.39-apache-kafka">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.9...</a>
</div>
---
<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>