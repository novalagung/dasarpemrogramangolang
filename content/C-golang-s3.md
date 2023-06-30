# C.38. Amazon Simple Storage Service (S3)

Pada bab ini kita akan belajar untuk membuat koneksi ke Amazon S3 menggunakan Golang. Mulai dari cara membuat bucket di S3, melihat semua daftar bucket di S3, melihat semua object/file yang ada di dalam sebuah bucket S3, serta mengupload dan mendownload file dari S3 bucket.

Kita mulai bahasan ini dengan mengenal apa itu Amazon S3 dan beberapa istilah yang berkaitan.

## C.38.1 Apa itu Amazon Simple Storage Service (S3)?

Pada dasarnya Simple Storage Service (S3) adalah layanan penyimpanan file/object yang dimiliki oleh Amazon Web Service (AWS). Layanan ini bisa kita nikmati secara gratis dengan batasan-batasan tertentu. Dengan menggunakan Amazon S3, kita bisa menyimpan dan melindungi object untuk berbagai kebutuhan sistem kita. Ringkasnya, kita bisa menganalogikan Amazon S3 sebagai harddisk/storage online yang bisa kita akses selama kita terhubung dengan internet.

## C.39.2 Beberapa istilah terkait Amazon S3

Beberapa istilah yang biasa kita temukan saat kita bekerja dengan Amazon S3 antara lain:
### 1. Bucket
Bucket adalah wadah untuk object bisa disimpan ke dalam Amazon S3. Kita bisa menganalogikan bucket seperti directory yang ada di harddisk kita, dimana kita bisa membuat folder/path dan menyimpan file di dalamnya. Seperti contoh, misal kita membuat bucket ```padinky-bucket-test``` di region ```ap-southeast-1``` dan mengupload file ```adamstudio.jpg```, maka kita bisa mengakses file tersebut dengan URL ```https://adamstudio-bucket.s3.ap-southeast-1.amazonaws.com/adamstudio.jpg``` (dengan authorisasi tertentu pastinya).

### 2. Object
Object secara singkat bisa kita artikan sebagai file, meskipun pada dasarnya berbeda, karena object juga menyimpan metadata file dan data-data lainnya.  
  
Untuk mempelajari lebih lanjut mengenai definisi dan beberapa istilah lain terkait Amazon S3, silakan merujuk ke https://docs.aws.amazon.com/id_id/AmazonS3/latest/userguide/Welcome.html


## C.39.3 Akses ke bucket S3

Untuk mengakses bucket di S3 melalui aplikasi, secara umum kita memerlukan credential berupa ```aws_access_key_id``` dan ```aws_secret_access_key```. Di pembahasan kali ini, kita tidak mencakup cara setting credential ke bucket S3 kita, untuk mempelajari lebih lanjut tentang memberikan akses ke bucket S3, bisa mengunjungi link berikut: https://docs.aws.amazon.com/id_id/AmazonS3/latest/userguide/about-object-ownership.html  
  
> Disini penulis asumsikan kita sudah memiliki akses berupa ```aws_access_key_id``` dan ```aws_secret_access_key``` untuk digunakan di aplikasi kita dalam membuat koneksi ke bucket S3.

## C.39.3 Koneksi ke S3

*And here we go...*

Pertama, kita akan membuat *S3 Client* untuk mendapatkan *session* dan terhubung dengan S3

```go
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func newS3Client(accessKeyID string, secretAccessKey string, region string) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})

	if err != nil {
		return nil, err
	}

	s3Client := s3.New(sess)

	return s3Client, nil
}
```


## C.39.4 Membuat bucket baru ke S3

## C.39.5 Melihat semua daftar bucket di S3

## C.39.6 Mengupload object ke dalam S3 bucket

## C.39.7 Mendownload object dari S3 bucket

## C.39.8 Menghapus object dari S3 bucket