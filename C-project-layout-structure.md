# C.1. Go Project Layout Structure

Mari kita awali pembahasan pada pemrograman Go lanjut dengan topik yang paling penting, yaitu tentang bagaimana manajemen file dan folder pada project Go.

[Sebenarnya tidak ada spesifikasi resmi dari Go mengenai bagaimana struktur project harus disusun](https://github.com/golang-standards/project-layout/issues/117). Akan tetapi ada beberapa project open source yang strukturnya digunakan sebagai basis **standar** dalam menyusun file dan folder program. Dan pada chapter ini kita akan mencoba membahas dan mempergunakan project tersebut sebagai acuan dalam membuat program Go.

## C.1.1. Library `golang-standard/project-layout`

Ada open source project yang sangat menarik untuk dipelajari, yaitu [project-layout](https://github.com/golang-standards/project-layout). Project tersebut isinya adalah project layout pada Go yang merupakan hasil kombinasi dari banyak project layout Go terkenal, seperti kubernetes, nats.io, istio, termasuk juga layout dari source code Go itu sendiri.

Perlu saya tekankan, bahwa Go bukan merupakan bahasa *functional* ataupun *object-oriented*, kita selaku programmer diberikan kebebasan terhadap bagaimana penulisan *source code* aplikasi yang dikerjakan. Akan tetapi, memang ada beberapa fitur milik OOP dan bahasa *functional* dalam Go, jadi ... bebas.

Termasuk juga perihal *project layout structure*, kita diberi kebebasan penuh. Di dokumentasi Go tidak ada panduan perihal bagaimana seharusnya desain struktur kode. Argumentasi ini diperkuat oleh [Russ Cox, yang merupakan Tech Lead proyek Go programming language](https://github.com/golang-standards/project-layout/issues/117).

Nah, dari sini sekarang sudah cukup jelas ya.

Ok, sekarang kembali ke project layout milik `golang-standard`. Saya sarankan untuk mempelajari dan mencoba struktur ini karena sangat umum diadopsi dalam pengembangan aplikasi menggunakan bahasa Go.

Pada chapter ini, saya hanya akan membahas garis besarnya saja, selebihnya jika ingin praktik bisa langsung clone dari https://github.com/golang-standards/project-layout.

## C.1.2. Struktur Layout `golang-standard/project-layout`

Ada cukup banyak folder dan subfolder dalam project layout, berikut kami ringkas beberapa file dan direktori yang umumnya dipakai.

```bash
.
├── go.mod
|   # file go.mod dipergunakan oleh go module (jika go mod diaktifkan).
|
├── Makefile
|   # file Makefile dipergunakan oleh command `make`.
|
├── assets/
|   # folder assets berisi static assets, seperti gambar, logo, dll.
|
├── build/
|   # folder build isinya adalah files untuk keperluan build dan
|   # juga CI (continous integration). Contoh file yang dimaksud adalah
|   # seperti Dockerfile, file CI tool (.travis-ci.yml, .gitlab-ci.yml)
|   # dan file untuk keperluan build ke bentuk lain seperti file deb, rpm, pkg.
|   |
│   ├── ci/
|   |   # tempatkan file untuk CI dalam folder ini
|   |
│   └── package/
|       # tempatkan file untuk keperluan build dalam folder ini
|
├── cmd/
|   # folder cmd isinya adalah source code utama aplikasi.
|   #
|   # jika aplikasi merupakan sebuah app monolith, maka folder ini isinya
|   # adalah langsung source code utama.
|   # salah satu contoh, folder ini isinya adalah file-file bisnis logic utama,
|   # seperti services dan repositories.
|   #
|   # jika arsitektur microservices diadopsi, dengan layout monorepo,
|   # maka isi dari cmd adalah source code yang dibagi per service.
|   |
│   ├── your_app_1/
│   ├── your_app_2/
│   ├── your_app_3/
│   └── ...
|
├── configs/
|   # folder configs isinya adalah file konfigurasi.
|
├── deployments/
|   # folder deployments isinya adalah file yang berhubungan dengan orchestration,
|   # deployments, dan juga CD. Seperti docker-compose.yml, k8s file, dll.
|
├── docs/
|   # folder docs isinya adalah file design dan dokumentasi.
|
├── examples/
|   # folder examples isinya adalah file example.
|
├── init/
|   # folder init isinya adalah file-file system init (systemd, upstart, sysv)
|   # dan file konfigurasi process manager atau supervisor (runit, supervisord).
|
├── internal/
|   # folder internal isinya adalah file private aplikasi dan library.
|   # sebetulnya folder ini kegunaannya sama seperti `pkg`, perbedaannya adalah package
|   # dalam folder internal ini hanya bisa di-import dalam project ini, tidak bisa di-import
|   # ke project lain.
|
├── pkg/
|   # folder pkg isinya adalah file utility yg di-reuse dalam project yang sama,
|   # atau bisa juga di re-use oleh project lain.
|   |
│   ├── your_public_lib_1/
│   ├── your_public_lib_2/
│   ├── your_public_lib_3/
│   └── ...
|
├── test/
|   # folder test isinya adalah file testing. untuk struktur file-nya sendiri bebas,
|   # mau disusun seperti apa.
|   #
|   # khusus untuk unit test, baiknya tidak ditempatkan disini,
|   # tapi ditempatkan di dalam package yang sama dengan file yang akan di-unit-test. 
|
├── vendor/
|   # berisi clone dari 3rd party dependencies. Folder ini digunakan jika konfigurasi vendor diaktifkan
|
├── web/
|   # berisi aplikasi web. untuk microservices saya sarankan untuk menempatkan aplikasi web dalam folder `cmd/app`
|
└── ...
```

Hmm, cukup banyak juga ya yang perlu dipelajari? 😅 Tenang, tidak perlu untuk dihafal, cukup dipahami saja. Selain itu semua direktori diatas juga belum tentu dipakai semua, perlu disesuaikan dengan proyek yang sedang teman-teman kembangkan.

Ok, sampai sini saja pembahasan mengenai project layout, selanjutnya silakan mencoba-coba jika berkenan, bisa dengan men-develop mulai awal, atau *clone* existing project untuk dipelajari strukturnya.

---

 - [Standard Go Project Layout](https://github.com/golang-standards/project-layout/), by Kyle Quest
