# C.38. Amazon S3 (Simple Storage Service)

Pada bab ini kita akan belajar untuk membuat koneksi ke Amazon S3 menggunakan Golang. Mulai dari cara membuat bucket di S3, melihat semua daftar bucket di S3, melihat semua object/file yang ada di dalam sebuah bucket S3, serta mengupload dan mendownload file dari S3 bucket.

Kita mulai bahasan ini dengan mengenal apa itu Amazon S3 dan beberapa istilah yang berkaitan.

## C.38.1 Apa itu Amazon Simple Storage Service (S3)?

Pada dasarnya Simple Storage Service (S3) adalah layanan penyimpanan file/object yang dimiliki oleh Amazon Web Service (AWS). Dengan menggunakan Amazon S3, kita bisa menyimpan dan melindungi object untuk berbagai kebutuhan sistem kita. Ringkasnya, kita bisa menganalogikan Amazon S3 sebagai harddisk/storage online yang bisa kita akses selama kita terhubung dengan internet.

> AWS menyediakan layanan [Free Tier](https://aws.amazon.com/free/), dengan kita bisa memanfaatkan service S3 secara gratis selama 12 bulan.

## C.39.2 Beberapa istilah terkait Amazon S3

Beberapa istilah yang biasa kita temukan saat kita bekerja dengan Amazon S3 antara lain:

#### • Bucket

Bucket adalah wadah untuk object bisa disimpan ke dalam Amazon S3. Kita bisa menganalogikan bucket seperti directory yang ada di harddisk kita, dimana kita bisa membuat folder/path dan menyimpan file di dalamnya. Seperti contoh, misal kita membuat bucket `adamstudio-bucket` di region `ap-southeast-1` dan mengupload file `adamstudio.jpg`, maka kita bisa mengakses file tersebut dengan URL `https://adamstudio-bucket.s3.ap-southeast-1.amazonaws.com/adamstudio.jpg` (dengan authorisasi tertentu pastinya).

#### • Object

Object secara singkat bisa kita artikan sebagai file, meskipun pada dasarnya berbeda, karena object juga menyimpan metadata file dan data-data lainnya.

> Untuk mempelajari lebih lanjut mengenai definisi dan beberapa istilah lain terkait Amazon S3, silakan cek https://docs.aws.amazon.com/id_id/AmazonS3/latest/userguide/Welcome.html

## C.39.3 *Authentication* dan *authorization* bucket S3

AWS S3 menyediakan beberapa jenis metode otentikasi untuk pengaksesan konten S3 via aplikasi, salah satunya menggunakan access keys (`aws_access_key_id` dan `aws_secret_access_key`).

Pada chapter ini kita akan menerapkan metode otentikasi tersebut, dengan asumsi *bucket policy* yang digunakan adalah *defalt*, dimana semua konten dalam bucket bisa diakses dan dimodifikasi.

> Lebih lanjut mengenai *bucket policy* bisa mengunjungi link berikut: https://docs.aws.amazon.com/id_id/AmazonS3/latest/userguide/about-object-ownership.html

> Disini penulis asumsikan kita sudah memiliki akses berupa `aws_access_key_id` dan `aws_secret_access_key` untuk digunakan di aplikasi kita dalam membuat koneksi ke bucket S3.

## C.39.3 Koneksi ke S3

*And here we go...*

Ok, seperti biasa buat proyek baru, kemudian buat 1 file bernama `main.go`. Definisikan package dan import beberapa dependensi yang dibutuhkan. Disini kita akan menggunaakan [official SDK dari AWS untuk Golang](https://github.com/aws/aws-sdk-go).

> Untuk dokumentasi detailnya bisa dibaca disini: https://aws.amazon.com/id/sdk-for-go/

Definisikan konstanta untuk menampung nilai access key, dan juga region bucket.

```go
package main

import (
    "context"
    "fmt"
    "os"
    "path/filepath"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
    AWS_ACCESS_KEY_ID     string = "AKIA**********"
    AWS_SECRET_ACCESS_KEY string = "LReO***********"
    AWS_REGION            string = "ap-southeast-1"
)
```

> Jangan lupa untuk `go get -u github.com/aws/aws-sdk-go` dependensi aws-sdk-go

Selanjutnya, kita siapkan fungsi untuk mendapatkaan objek `*session.Session` dari AWS. Object ini berisi informasi session pengaksesan *resource* AWS, yang pada konteks ini adalah AWS S3.

```go
func newSession() (*session.Session, error) {
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(AWS_REGION),
        Credentials: credentials.NewStaticCredentials(
            AWS_ACCESS_KEY_ID,
            AWS_SECRET_ACCESS_KEY,
            "",
        ),
    })

    if err != nil {
        return nil, err
    }

    return sess, nil
}
```

Panggil fungsi `newSession()` di atas, lalu bungkus menggunakan fungsi `s3.New()` untuk mendapatkan object session `*s3.S3`, yang pada kode berikut direpresentasikan oleh variabel `s3Client`. Lewat object tersebut nantinya operasi pengaksesan dan manipulasi konten S3 dilakukan.

```go
func main() {
    sess, err := newSession()
    if err != nil {
        fmt.Println("Failed to create AWS session:", err)
        return
    }

    s3Client := s3.New(sess)
    fmt.Println("S3 session & client initialized")

    // ...
}
```

## C.39.4 Membuat bucket baru ke S3

Gunakan method `CreateBucket()` milik object `client` untuk membuat bucket baru, siapkan statement dalam fungsi bernama `createBucket()`. Tulis nama bucket pada parameter pemanggilan fungsi dengan skema penulisan bisa dilihat di bawah ini:

```go
func createBucket(client *s3.S3, bucketName string) error {
    _, err := client.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucketName),
    })

    return err
}
```

Navigasi ke fungsi `main()`, panggil fungsi `createBucket()` di atas. Pada contoh berikut, bucket name yang dipilih adalah `bucketName`.

```go
func main() {
    // ...

    bucketName := "adamstudio-new-bucket"

    // =============== create bucket ===============
    err = createBucket(s3Client, bucketName)
    if err != nil {
        fmt.Printf("Couldn't create new bucket: %v", err)
        return
    }

    fmt.Println("New bucket successfully created")

    // ...
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_1.png)

## C.39.5 Melihat semua daftar bucket di S3

Sekarang siapkan fungsi baru untuk menampilkan daftar bucket. Gunakan method `ListBuckets()`. Pada contoh berikut method dipanggil dengan parameter `nil` karena kita tidak menggunakan filter.

```go
func listBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {
    res, err := client.ListBuckets(nil)
    if err != nil {
        return nil, err
    }

    return res, nil
}
```

Modifikasi fungsi `main()`, kemudian tambahkan baris kode berikut untuk memanggil fungsi `listBuckets()` di atas. Lewat nilai balik fungsi tersebut, akses property `.Buckets` untuk mendapatkan list buckets.

```go
func main() {
    // ...

    // =============== list all buckets ===============
    buckets, err := listBuckets(s3Client)
    if err != nil {
        fmt.Printf("Couldn't list buckets: %v", err)
        return
    }

    for _, bucket := range buckets.Buckets {
        fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
    }

    // ...
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_2.png)

## C.39.6 Mengupload object ke dalam S3 bucket

Operasi upload file bisa dilakukan via beberapa cara, salah satunya menggunakan method `Upload()` milik object uploader (yarn bertipe `*s3manager.Uploader`) yang dalam pemanggilannya, file yang ingin di-upload disisipkan sebagai argument pemanggilan method.

Pada contoh berikut, variabel `file` merepresentasikan file yang akan di upload. Variabel ini disisipkan pada pemanggilan method `Upload()`.

```go
func uploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }

    defer file.Close()

    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(fileName),
        Body:   file,
    })

    return err
}
```

Selanjutnya di fungsi `main()`, siapkan object uploader, caranya mudah yaitu dengan membungkus object session menggunakan fungsi `s3manager.NewUploader()`.

> Pastikan untuk tidak meng-import dependensi `github.com/aws/aws-sdk-go/service/s3/s3manager`

Tak lupa siapkan juga sample file yang akan digunakan untuk testing (untuk di upload ke bucket). Pada contoh ini penulis menggunakan file `./upload/adamstudio.jpg`.

Modifikasi lagi fungsi `main`, tambahkan statement-statement di atas dan panggil fungsi `uploadFile()`.

```go
func main() {
    // ...

    // =============== upload file ===============
    uploader := s3manager.NewUploader(sess)

    fileName := "adamstudio.jpg"
    filePath := filepath.Join("upload", fileName)

    err = uploadFile(uploader, filePath, bucketName, fileName)
    if err != nil {
        fmt.Printf("Failed to upload file: %v", err)
    }

    fmt.Println("Successfully uploaded file!")

    // ...
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_3.png)

## C.39.7 Menampilkan daftar objects dalam bucket

Gunakan method `ListObjectsV2()` untuk menampilkan isi bucket.

```go
func listObjects(client *s3.S3, bucketName string) (*s3.ListObjectsV2Output, error) {
    res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
        Bucket: aws.String(bucketName),
    })
    if err != nil {
        return nil, err
    }

    return res, nil
}
```

Panggil fungsi `listObjects()` yang telah dibuat di atas. Lewat nilai balik fungsi tersebut, akses property `.Contents` untuk mendapatkan list objects.

```go
func main() {
    // ...

    // =============== list objects ===============
    // bucketName := "adamstudio-new-bucket"
    objects, err := listObjects(s3Client, bucketName)
    if err != nil {
        fmt.Printf("Couldn't list objects: %v", err)
        return
    }

    for _, object := range objects.Contents {
        fmt.Printf("Found object: %s, size: %d\n", *object.Key, *object.Size)
    }

    // ...
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_5.png)

## C.39.8 Men-download object dari S3 bucket

Untuk proses download, kita harus mempersiapkan satu file object terlebih dahulu untuk menampung konten hasil operasi download.

Pada contoh berikut, object `file` adalah file yang akan menampung operasi download. Object tersebut disisipkan sebagai argument pemanggilan fungsi `Download()` milik object downloader bertipe `*s3manager.Downloader`.

```go
func downloadFile(downloader *s3manager.Downloader, bucketName string, key string, downloadPath string) error {
    file, err := os.Create(downloadPath)
    if err != nil {
        return err
    }

    defer file.Close()

    _, err = downloader.Download(
        file,
        &s3.GetObjectInput{
            Bucket: aws.String(bucketName),
            Key:    aws.String(key),
        },
    )

    return err
}
```

Buat object downloader menggunakan fungsi `s3manager.NewDownloader()`, siapkan juga variabel untuk path file, lalu panggil method `downloadFile()`.

> Pastikan untuk meng-import dependensi `github.com/aws/aws-sdk-go/service/s3/s3manager`

```go
func main() {
    // ...

    // =============== download file ===============
    downloader := s3manager.NewDownloader(sess)
    fileName := "adamstudio.jpg"
    bucketName := "adamstudio-new-bucket"
    downloadPath := filepath.Join("download", fileName)
    err = downloadFile(downloader, bucketName, fileName, downloadPath)
    if err != nil {
        fmt.Printf("Couldn't download file: %v", err)
        return
    }

    fmt.Println("Successfully downloaded file")

    // ...
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_4.png)

## C.39.9 Menghapus object dari S3 bucket

Operasi delete object bisa dilakukan menggunakan method `DeleteObject()` milik object s3 client:

```go
func deleteFile(client *s3.S3, bucketName string, fileName string) error {
    _, err := client.DeleteObject(&s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(fileName),
    })

    return err
}
```

Panggil fungsi di atas pada `main()`, lalu sisipkan variabel file yang ingin di-delete.

```go
func main() {
    // ...

    // =============== delete file ===============
    fileName := "adamstudio.jpg"
    bucketName := "adamstudio-new-bucket"
    err = deleteFile(s3Client, bucketName, fileName)
    if err != nil {
        fmt.Printf("Couldn't delete file: %v", err)
        return
    }

    fmt.Println("Successfully delete file")
}
```

Jalankan program dan lihat hasilnya.

![S3 test](images/C_aws_s3_6.png)

## C.39.10 Presign URL

Presign URL adalah salah satu metode untuk sharing object bisa diakses oleh publik (lewat internet). Dengan presign url, kita bisa menentukan durasi berapa lama object bisa diakses oleh public.

Cara penerapannya cukup mudah, pertama akses *instance* object menggunakan method `GetObjectRequest()`. Pada argument pemanggilan method, isi `Key` dengan nama object. Lalu gunakan nilai balik pertama statement untuk mengakses method `Presign()` sekaligus tentukan durasinya.

Pada kode berikut, method `Presign()` mengembalikan URL object yang valid selama `15 menit`.

```go
func presignUrl(client *s3.S3, bucketName string, fileName string) (string, error) {
    req, _ := client.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(fileName),
    })

    urlStr, err := req.Presign(15 * time.Minute)
    if err != nil {
        return "", err
    }

    return urlStr, nil
}
```

Panggil method tersebut di fungsi `main()` lalu coba akses URL nya.

```go
func main() {
    // ...

    // =============== presign url ===============
    fileName := "adamstudio.jpg"
    bucketName := "adamstudio-new-bucket"
    urlStr, err := presignUrl(s3Client, bucketName, fileName)
    if err != nil {
        fmt.Printf("Couldn't presign url: %v", err)
        return
    }

    fmt.Println("Presign url:", urlStr)
}
```

Hasilnya:

![S3 test](images/C_aws_s3_7.png)

![S3 test](images/C_aws_s3_8.png)

---

 - [aws-sdk-go](https://github.com/aws/aws-sdk-go), by AWS, Apache 2.0 License

---

<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-C.38-aws-s3">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-C.38...</a>
</div>

---

<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>
