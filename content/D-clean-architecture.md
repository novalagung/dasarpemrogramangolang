# D.4 Clean Architecture Golang
pada chapter ini kita akan praktek penerapan salah satu arsitektur yang banyak digunakan oleh developer, yaitu clean architecture.
> NOTE: Clean architecture tidak ada batasan apa saja teknologi atau bahasa pemrograman yang dipakai, namun di tutorial ini saya menggunakan golang native tanpa menggunakan framework seperti go-fiber, go-chi dan sebagainya untuk implementasi dari architecture ini.
## D.4.1 Penjelasan
Clean Architecture adalah sistem arsitektur yang dibuat oleh Robert C.Martin(Uncle Bob) yang merupakan berasal terdiri dari: Hexagonal Architecture, Onion Architecture, Screaming Architecture, dan sebagainya selama beberapa tahun. Arsitektur ini membuat setiap proses dibagi menjadi layer tersendiri yang berdiri sendiri tanpa ada interfensi dari layer atas.
Menurut Uncle Bob ada 5 keuntungan untuk memakai arsitektur ini, yaitu:
#### • Testable
Proses bisnis bisa di testing tanpa mengubah User Interface, Database, Webserver, atau external komponen. Kita bisa membuat dan menjalankan skenario testing proses tanpa mengubah komponen yang sudah ada.
#### • Independent of UI
User Interface lebih mudah untuk diubah tanpa mengubah seluruh yang ada di sistem. Sebagai contoh User Interface web bisa dengan console ui tanpa mengubah aturan proses bisnis.
#### • Independent of Database
Independen database yang berarti bisa mengganti database Oracle atau SQL Server ke MongoDB, BigTable, CouchDB, dan sebagainya. Jadi aturan proses bisnis tidak melompati ke database.
#### • Independent Framework
Clean architecture tidak mempunyai ketergantungan dengan beberapa library fitur yang ada di perangkat lunak. Ini seperti memperbolehkan menggunakan framework sebagai alat atau tool, daripada harus memaksa menambahkan di sistem kita ke dalam limitnya.
#### • Independent of any External
Secara fakta aturan proses bisnis kita secara sederhana tidak mengetahui apapun yang ada di eksternal layer.

Di bawah ini adalah layer-layer diagram dari clean architecture
 ![](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)
 Seperti yang terlihat ada 4 layer dibuat oleh uncle bob yaitu: Entities, Usecases, Interface Adapters, Framework & Drivers:

#### 1. Entities
Entities digunakan untuk membuat object, method, atau bisa juga data structure dan function. Entities tidak melakukan perubahan yang ada diluar, jadi apabila ada perubahan di layer atas seperti usecase atau controller, entities tidak ada interfensi untuk merubah alur proses bisnis yang terjadi.
#### 2. Usecases
Layer Usecase ini berisi tentang alur proses bisnis di aplikasi spesifik. Layer ini merangkum dan mengimplementasikan dari entities menjadi alur bisnis proses yang dibuat. Di dalam usecase tidak ada perubahan yang berefek kepada entities.
#### 3. Interface Adapters
Interface Adapters memiliki struktur didalamnya antara lain, Controllers, Presenters dan Views. Layer ini dipakai untuk adaptor dengan layer usecase dan entities. Alur berjalannya layer interface adapter ini adalah dari controller mengirim data ke usecase, lalu dari usecase dilempar ke Presenter dan ke Views.
#### 4. Framework & Drivers
Layer ini adalah paling luar dari 3 layer diatas, secara umum layer ini tersusun dari framework dan alat seperti database, web framework, dan sebagainya. Di sini kalian tidak membuat code untuk menyambungkan komunikasi ke layer lainnya. 

## D.4.2 Persiapan 
Kita akan mempersiapkan struktur folder untuk membuat clean architecture, karena architecture ini setiap layer harus dipisah maka kita akan membuat folder untuk setiap layer. Berikut ini struktur foldernya:

```bash
.
├── configs
│    # configs berisi konfigurasi seperti database, logger, dan sebagainya.
├── controllers
│    # controller berisi route api dan menyisipkan layer usecase untuk melakukan bisnis proses.
├── entities
│    # entities isinya adalah untuk membuat model data, model method atau juga model function.
├── repositories
│    # repositories berisi function atau method dengan model entities yang sudah dibuat.
├── usecases 
│    # usecases berisi function yang berisi ochestrator dari repositories.
└── main.go
```
Lalu kita akan menyiapkan beberapa library golang yang dipakai untuk aplikasi kita:
- #### Godotenv
Godotenv library untuk membuat environment aplikasi.
``` go get -u github.com/joho/godotenv ```
- #### Gorm
Gorm library digunakan untuk membuat orm database sql.
``` go get -u gorm.io/gorm ```
- #### Go-playground Validator
Go-playground validator dipakai untuk membuat validasi dari request.
``` go get -u github.com/go-playground/validator/v10 ```
- #### Gorm Postgres
Gorm postgres digunakan untuk orm database sql khusus untuk postgres.
``` go get -u gorm.io/driver/postgres ```

## D.4.3 Praktek
Setelah struktur folder project sudah dibuat, kita akan implementasi clean architecture pada tutorial ini. Di setiap layer kita akan membuat dependecy injection dari Entities, Usecase, dan Controller.
### D.4.3.1 Config
Kita akan membuat 1 folder bernama ``` config ```. Folder ini berisi konfigurasi yang akan kita pakai di aplikasi seperti environment, database, validator, dan lain-lain.
Kita buat file ``` .env ``` yang berisi konfigurasi database dan aplikasi.
- #### File .env
```env
DATABASE_HOST=127.0.0.1
DATABASE_PORT=5432
DATABASE_USERNAME=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=testing-clean-architecture

APP_NAME=clean-architeture-example
APP_HOST=127.0.0.1
APP_PORT=9090
APP_VERSIONAPI=v1
```
Untuk database kalian menyesuaikan dengan local kalian masing-masing. Di file ini juga kita membuat konfigurasi dari aplikasi seperti nama aplikasi, host, port, dan version api yang akan kita pakai.
Setelah sudah kita membuat 1 file lagi bernama ``` database.go ```, file ini berisi konfigurasi dan koneksi database postgres.
- #### File database.go
```go
package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func NewDatabaseConfig(config DatabaseConfig) (db *gorm.DB, err error) {
	stringConnection := "host=" + config.Host + " user=" + config.Username + " password=" + config.Password + " dbname=" + config.DBname + " port=" + config.Port + " TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	// SetConnMaxIdleTime: berfungsi untuk menetapkan jumlah maksimum waktu koneksi secara idle(tidak berjalan)
	sqlDB.SetConnMaxIdleTime(1 * time.Minute)
	// SetConnMaxLifetime: berfungsi menentukan maksimum waktu dapat digunakan kembali
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	// SetMaxIdleConns: berfungsi untuk menentukan jumlah connection tidak dijalankan(idle)
	sqlDB.SetMaxIdleConns(20)
	// SetMaxOpenConns: berfungsi untuk menetapkan jumlah open koneksi
	sqlDB.SetMaxOpenConns(5)

	return
}
```

### D.4.3.2 Entities
Layer entities ini berisi data struct, method, function dan juga entities.Kita akan buat 3 file di folder ```entities``` dengan nama ``` products.go```, ```product_request.go```, dan  ```product_response.go```  dengan isi kode dibawah ini:
- #### File products.go
```go
package entities

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ID     int     `gorm:"column:id"`
	Name   string  `gorm:"column:name"`
	Price  float64 `gorm:"column:price"`
	Weight float64 `gorm:"column:weight"`
}

// Membuat model interface untuk repository product
type ProductRepository interface {
	GetByID(id int) (Products, error)
	Gets() ([]Products, error)
	Create(product *Products) error
	Update(product *Products) error
	DeleteByID(id int) (Products, error)
	Count() (int64, error)
}

// Membuat model interface untuk usecase product
type ProductUsecase interface {
	GetOne(id int) (ProductResponseJSON, error)
	Gets() (ProductResponseJSON, error)
	Create(product CreateProductRequest) (ProductResponseJSON, error)
	Update(product UpdateProductRequest) (ProductResponseJSON, error)
	DeleteByID(id int) (ProductResponseJSON, error)
}
```
Pada file ```products.go``` kita membuat model untuk table database dan model repository dan usecase yang nanti akan di inject.
- #### File product_request.go
```go
package entities

type CreateProductRequest struct {
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}

type UpdateProductRequest struct {
	ID     int     `json:"id" validate:"required"`
	Name   string  `json:"name" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}

type DeleteProductRequest struct {
	ID int `json:"id" validate:"required"`
}
```
Pada file ```product_request.go``` kita membuat model untuk request data untuk proses api seperti insert, update, dan delete.Disamping itu juga di setiap tag request terdapat code ``` validate:"required" ``` yang berfungsi sebagai validasi bahwa setiap request tidak boleh kosong dan akan di lakukan oleh library go-playground validator. 
- #### File product_response.go
```go
package entities

// Struct request and response product
type ProductResponseJSON struct {
	Data    []Products `json:"data"`
	Count   int64      `json:"count"`
	Success bool       `json:"success"`
	Message string     `json:"message"`
}
```
Pada file ```product_response.go``` kita membuat model untuk response data pada setiap api yang dibuat. Di sini kita hanya menambahkan tag json saja.

### D.4.3.2 Repository
Di dalam folder ``` repositories ``` buat 1 file ``` product_repository.go ```. Di layer repository ini digunakan untuk meng-handle method database seperti gets, put, delete, get, dan sebagainya.Layer ini bertanggung jawab untuk database yang akan dipakai di aplikasi.Untuk code-nya seperti di bawah ini.
- #### File product_repository.go
```go
package repositories

import (
	"clean-architecture-golang-example/entities"
	"errors"
	"log"

	"gorm.io/gorm"
)

// ProductRepositores: repository untuk model product
type ProductRepositories struct {
	database *gorm.DB
}

// NewProductRepositories: Injeksi repository product model
func NewProductRepositories(conn *gorm.DB, isMigrate bool) entities.ProductRepository {
	if isMigrate {
		err := conn.AutoMigrate(entities.Products{})
		if err != nil {
			log.Fatal("Migration Error:", err)
		}
	}
	return &ProductRepositories{conn}
}

// Create: digunakan untuk membuat insert data ke model product.
func (p *ProductRepositories) Create(product *entities.Products) error {
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(entities.Products{}).Create(product)
	err = query.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}

// Count: digunakan untuk menghitung jumlah data product yang tersimpan.
func (p *ProductRepositories) Count() (int64, error) {
	var count int64
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(entities.Products{}).Select("*").Count(&count)
	err = query.Error
	if err != nil {
		return count, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return count, err
	}

	return count, err
}

// DelteByID: digunakan untuk menghapus data product dengan id yang dipilih.
func (p *ProductRepositories) DeleteByID(id int) (entities.Products, error) {
	var product entities.Products
	var err error

	queryFind := p.database.Model(entities.Products{}).Where("id = ?", id).Find(&product)
	err = queryFind.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, errors.New("ID is not found")
		}
		return product, err
	}

	queryDelete := queryFind.Delete(&product)
	err = queryDelete.Error
	if err != nil {
		return product, err
	}

	return product, err
}

// GetByID: digunakan untuk menampilkan data product yang sesuai dengan id yang dipilih.
func (p *ProductRepositories) GetByID(id int) (entities.Products, error) {
	var result entities.Products
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(&entities.Products{}).Where("id = ?", id).Where("deleted_at IS NULL").First(&result)
	err = query.Error
	if err != nil {
		return result, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return result, err
	}

	return result, err
}

// Gets: digunakan untuk menampilkan semua data product.
func (p *ProductRepositories) Gets() ([]entities.Products, error) {
	var results []entities.Products
	var err error
	var tx *gorm.DB = p.database.Begin()

	query := tx.Model(&entities.Products{}).Select("*").Where("deleted_at IS NULL").Scan(&results)
	err = query.Error
	if err != nil {
		return results, err
	}

	query = tx.Commit()
	err = query.Error
	if err != nil {
		return results, err
	}

	return results, err
}

// Update: digunakan untuk update data product.
func (p *ProductRepositories) Update(product *entities.Products) error {
	var err error
	var tx *gorm.DB = p.database.Begin()

	queryFind := tx.Model(entities.Products{}).Where("id = ?", product.ID).Updates(&product)
	err = queryFind.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("ID is not found")
		}
		return err
	}

	queryFind = tx.Commit()
	err = queryFind.Error
	if err != nil {
		return err
	}

	return err
}
```
Di dalam function ``` NewProductRepositories ``` terdapat 2 parameter yaitu:
- ``` conn *gorm.DB ``` digunakan untuk menerima connection database. Jadi untuk bagaimana query ini akan ke database atau konfigurasi kemana akan di ditambahkan ke file ``` main.go ```
- ```isMigrate bool``` untuk menentukan pilihan apakah model ``` Products ``` di folder entities akan di migrasi ke database atau tidak. Migrasi ini akan di buat setelah server dijalankan.
Selain itu code ``` conn.AutoMigrate(entities.Products{}) ``` di dalam function tersebut digunakan untuk menjalankan migrasi ke database secara otomatis dengan model data yang kita buat. 
Apabila kita lihat nilai return dari function ``` NewProductRepositories ``` adalah struct ``` ProductRepositories ``` dengan meng-injeksi ke model interface ``` entities.ProductRepository ```. Method dalam repository ini berisi query CRUD seperti biasa.

### D.4.3.3 Usecase
Di folder usecase kita membuat 1 file bernama ``` product_usecase.go ```, file ini berisi orkestrator yang berisi method-method entities yang sudah kita siapkan sebelumnya. Pada layer ini bertugas untuk meng-handle logic bisnis proses.
- #### File product_usecase.go
```go
package usecases

import (
	"clean-architecture-golang-example/entities"
	"errors"

	"github.com/go-playground/validator/v10"
)

// ProductUsecase: sebagai orchestrator bisnis proses product.
type ProductUsecase struct {
	repository *entities.ProductRepository
	valid      *validator.Validate
}

// NewProductUsecase: injeksi dari repository ke usecase
func NewProductUsecase(repository *entities.ProductRepository, valid *validator.Validate) entities.ProductUsecase {
	return &ProductUsecase{repository, valid}
}

// Create: digunakan untuk insert product ke repository.
func (usecase *ProductUsecase) Create(product entities.CreateProductRequest) (entities.ProductResponseJSON, error) {
	var err error
	var result entities.ProductResponseJSON
	repo := *usecase.repository
	err = usecase.valid.Struct(product)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error validation:" + err.Error(),
		}
		return result, err
	}

	count, err := repo.Count()
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	var data = entities.Products{
		ID:     int(count) + 1,
		Name:   product.Name,
		Price:  product.Price,
		Weight: product.Weight,
	}

	if err = repo.Create(&data); err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   1,
		Success: true,
		Message: "Create product success",
	}

	return result, err
}

// DeleteByID: digunakan untuk hapus product dengan id ke repository.
func (usecase *ProductUsecase) DeleteByID(id int) (entities.ProductResponseJSON, error) {
	var result entities.ProductResponseJSON
	if id == 0 {
		return result, errors.New("ID must be not empty")
	}
	repo := *usecase.repository
	data, err := repo.DeleteByID(id)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server:" + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   1,
		Success: true,
		Message: "Delete product success",
	}

	return result, nil
}

// GetOne: digunakan untuk mengambil data product dengan id yang sudah dipilih
func (usecase *ProductUsecase) GetOne(id int) (entities.ProductResponseJSON, error) {
	var result entities.ProductResponseJSON
	if id == 0 {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server: ID must be not empty",
		}
		return result, errors.New("ID must be not empty")
	}
	repo := *usecase.repository

	data, err := repo.GetByID(id)
	if err != nil {
		result = entities.ProductResponseJSON{
			Data:    []entities.Products{},
			Count:   0,
			Success: false,
			Message: "Error Internal Server: " + err.Error(),
		}
		return result, err
	}

	result = entities.ProductResponseJSON{
		Data: []entities.Products{
			data,
		},
		Count:   1,
		Success: true,
	}

	return result, nil
}

// Gets: digunakan untuk menampilkan semua data product
func (usecase *ProductUsecase) Gets() (entities.ProductResponseJSON, error) {
	repo := *usecase.repository
	data, err := repo.Gets()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	count, err := repo.Count()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	result := entities.ProductResponseJSON{
		Data:    data,
		Count:   count,
		Success: true,
	}

	return result, nil
}

// Update: digunakan untuk mengubah data dengan id yang sudah dipilih
func (usecase *ProductUsecase) Update(product entities.UpdateProductRequest) (entities.ProductResponseJSON, error) {
	err := usecase.valid.Struct(product)
	var result entities.ProductResponseJSON
	if err != nil {
		return result, err
	}
	repo := *usecase.repository
	var data = entities.Products{
		ID:     product.ID,
		Name:   product.Name,
		Price:  product.Price,
		Weight: product.Weight,
	}

	if err = repo.Update(&data); err != nil {
		return result, err
	}

	count, err := repo.Count()
	if err != nil {
		return entities.ProductResponseJSON{}, err
	}

	result = entities.ProductResponseJSON{
		Data:    []entities.Products{data},
		Count:   count,
		Success: true,
		Message: "Update product success",
	}

	return result, nil
}

```
Seperti yang kita lihat, code di atas sama secara struktur dengan file ``` product_repository.go ``` namun di sini kita hanya memanggil method dari repository layer.Parameter ``` repository *entities.ProductRepository ``` adalah hasil injeksi dari repository layer yang sudah kita buat.Sebagai Contoh method ``` Gets() ``` di usecase berisi 2 method repository yang dipanggil yaitu ``` repo.Gets() ``` dan ``` repo.Count() ```.Sama seperti method di usecase yang lain.Di usecase ini tidak hanya dipakai untuk memanggil method repository layer saja namun kita juga gunakan untuk mapping data, membuat validasi data, dan masih banyak lagi. Untuk validasi code ``` err := usecase.valid.Struct(product) ``` dipakai untuk memvalidasi data dengan model request yang sudah kita sebelumnnya.Pada layer ini kita bisa memakai lebih dari 1 repository yang kita pakai.

### D.4.3.4 Controller
Pada layer controller digunakan untuk bagaimana bentuk data yang nanti akan di tampilkan. Bentuk data bisa berupa REST API, HTML, XML, atau GRPC. Di layer ini juga layer usecase akan dipanggil dan nanti akan diolah oleh layer usecase secara mandiri. Kita buat 1 file baru dengan nama ``` product_controller.go ``` dengan code seperti di bawah ini.
- #### File product_controller.go
```go
package controllers

import (
	"clean-architecture-golang-example/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ProductController struct {
	versionApi string
	logger     *log.Logger
	http       *http.ServeMux
	usecase    *entities.ProductUsecase
}

func NewProductController(versionApi string, logger *log.Logger, http *http.ServeMux, usecase *entities.ProductUsecase) *ProductController {
	controller := &ProductController{versionApi, logger, http, usecase}
	controller.Route()

	return controller
}

func (Controller *ProductController) Gets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		controller := *Controller.usecase

		if id == 0 {
			result, err := controller.Gets()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Controller.logger.Println("[ERROR]:" + err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				Controller.logger.Println("[SUCCESS]: Gets product is success")
			}

			json.NewEncoder(w).Encode(result)
		} else {
			result, err := controller.GetOne(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Controller.logger.Println("[ERROR]:" + err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				Controller.logger.Println("[SUCCESS]: Get product by id is success")
			}

			json.NewEncoder(w).Encode(result)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		var data entities.CreateProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.Create(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusCreated)
			Controller.logger.Println("[SUCCESS]: Create product is success")
		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPut {
		var data entities.UpdateProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.Update(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			Controller.logger.Println("[SUCCESS] Update product is success")

		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodDelete {
		var data entities.DeleteProductRequest

		controller := *Controller.usecase
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		}

		result, err := controller.DeleteByID(data.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			Controller.logger.Println("[ERROR]:" + err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			Controller.logger.Println("[SUCCESS] Delete product is success")

		}
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		Controller.logger.Println("[ERROR]: method not allowed")
	}
}

func (Controller *ProductController) Route() {
	Controller.http.HandleFunc(Controller.versionApi+"/product/gets", Controller.Gets)
	Controller.http.HandleFunc(Controller.versionApi+"/product/create", Controller.Create)
	Controller.http.HandleFunc(Controller.versionApi+"/product/update", Controller.Update)
	Controller.http.HandleFunc(Controller.versionApi+"/product/delete", Controller.Delete)
}
```
Seperti yang kita lihat di layer ini juga menambahkan method ``` Route() ``` yang berfungsi untuk register url api yang akan kita pakai. Selain itu terdapat juga parameter ``` logger *log.Logger ``` yang dipakai untuk membuat log history proses dan ```http *http.ServeMux``` digunakan untuk membuat routing url api. Untuk routing lebih lengkapnya kalian bisa membaca ulang artikel ini: [B.2. Routing http.HandleFunc](https://dasarpemrogramangolang.novalagung.com/B-routing-http-handlefunc.html).Setelah 3 layer sudah dibuat maka kita hanya tinggal meng-injeksi tiap-tiap layer itu ke main function.
### D.4.3.5 Main Function
Setelah 3 layer sudah dibuat, kita akan membuat file ```main.go``` di directory yang sama. Tuliskan code-nya seperti di bawah ini.
```go
package main

import (
	"clean-architecture-golang-example/config"
	"clean-architecture-golang-example/controllers"
	"clean-architecture-golang-example/repositories"
	"clean-architecture-golang-example/usecases"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		panic(fmt.Errorf("Error Environment : %s", err))
	}

	appConfig := struct {
		name       string
		host       string
		port       string
		versionapi string
	}{
		name:       os.Getenv("APP_NAME"),
		host:       os.Getenv("APP_HOST"),
		port:       os.Getenv("APP_PORT"),
		versionapi: "/api/" + os.Getenv("APP_VERSIONAPI"),
	}

	envDBConfig := config.DatabaseConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DBname:   os.Getenv("DATABASE_Name"),
	}

	Logger := log.New(os.Stdout, appConfig.name+" || ", log.LstdFlags)
	validate := validator.New()

	db, err := config.NewDatabaseConfig(envDBConfig)
	if err != nil {
		panic(fmt.Errorf("Error Connection : %s", err))
	}

	serveMux := http.NewServeMux()

	// layer repository
	productRepository := repositories.NewProductRepositories(db, true)
	// layer usecase
	productUsecase := usecases.NewProductUsecase(&productRepository, validate)
	// layer controller
	controllers.NewProductController(appConfig.versionapi, Logger, serveMux, &productUsecase)

	s := &http.Server{
		Addr:         ":" + appConfig.port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second, // waktu maksimal untuk koneksi menggunakan TCP
		ReadTimeout:  5 * time.Second,   // waktu maksimal untuk membaca request dari client
		WriteTimeout: 5 * time.Second,   // waktu maksimal untuk menulis respon untuk client
	}

	go func() {
		Logger.Printf("Starting server on port :%s\n", appConfig.port)
		err := s.ListenAndServe()
		if err != nil {
			Logger.Fatal(err)
		}
	}()

	// mendapatkan signal ketika signal tertangkap ada interupsi dan akan shutdown server 
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// block sampai signal diterima
	sig := <-sigChan
	Logger.Println("Received terminate, Graceful shutdown: ", sig)

	// menunggu maksimal 30 detik untuk operasi sudah selesai
	tContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tContext)
}

```

- ``` sigChan := make(chan os.Signal, 1) ``` adalah variable untuk 1 channel yang berisi signal notifikasi. Ini akan menerima notifikasi signal dengan kondisi diinterupsi atau dishutdown.
- ```signal.Notify``` code ini untuk mendaftarkan channel yang nantinya akan menerima notifikasi signal.
- Untuk http server kita menggunakan goroutine untuk running server. Jadi server bisa jalan terlebih dahulu tanpa menunggu yang lain.
- Di file ``` main.go ``` berisi banyak konfigurasi seperti database, logger, dan http server. Di sini juga layer-layer yang sudah kita buat di tempelkan di file ini sehingga bisnis proses dari 3 layer tersebut bisa digunakan.
## D.4.4 Testing
Setelah sudah membuat clean architecture, kita akan mengetest hasil dari implementasi yang sudah kita buat. Jalankan aplikasi server ``` main.go ```.
![](image/D-run-server.png)
Seperti yang kita bahas sebelumnya, model entities yang sudah kita buat akan migrasi ke konfigurasi database environment. Lalu kita akan mencoba testing menggunakan api dengan Postman.
Untuk testing di sini kita akan menggunakan api untuk insert data. Isikan request data seperti dibawah ini:
```json
{
    "name": "Pasta gigi",
    "price": 100.0,
    "weight": 9
}
```
Setelah itu langsung trigger, maka apabila sukses akan menampilkan response seperti ini.
```json
{
    "data": [
        {
            "CreatedAt": "2023-10-08T02:07:02.1239668+07:00",
            "UpdatedAt": "2023-10-08T02:07:02.1239668+07:00",
            "DeletedAt": null,
            "ID": 1,
            "Name": "Pasta gigi",
            "Price": 100,
            "Weight": 9
        }
    ],
    "count": 1,
    "success": true,
    "message": "Create product success"
}
```
Dengan ini kita sudah berhasil membuat clean architecture sederhana menggunakan golang. Untuk api yang lain kalian bisa testing sendiri. Seperti catatan yang saya lampirkan di atas architecture ini bebas bagaimana kalian implementasikannya seperti apa selama kalian paham dari konsep dari clean architecture ini. Kalian bisa menggunakan framework golang, library yang lain, perubahan nama layer, dan sebagainya.

---
 - [Godotenv](https://github.com/joho/godotenv), MIT license
 - [Gorm](https://github.com/go-gorm/gorm), MIT license
 - [Go-playground-Validator](https://github.com/go-playground/validator), MIT Licence
 - [Gorm-Postgres](https://github.com/go-gorm/postgres), MIT Licence
---
<div class="source-code-link">
    <div class="source-code-link-message">Source code praktek chapter ini tersedia di Github</div>
    <a href="https://github.com/novalagung/dasarpemrogramangolang-example/tree/master/chapter-D.4-clean-architecture-golang">https://github.com/novalagung/dasarpemrogramangolang-example/.../chapter-D.4...</a>
</div>
---
<iframe src="https://novalagung.substack.com/embed" width="100%" height="320" class="substack-embed" frameborder="0" scrolling="no"></iframe>