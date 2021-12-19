=========================================================MEMBUAT REST API DENGAN GIN GOLANG=========================================================

1.  install golang di golang.org
2.  buat folder pustaka-api
3.  generate projek pustaka-api dengan cara ketik di terminal dalam folder pustaka-api "go mod init pustaka-api"
4.  nanti ada file baru didalam folder tersebut, yaitu go.mod
5.  install gin-gonic untuk membangun rest api
6.  nanti ada file baru didalam folder, yaitu go.sum
7.  buat main.go untuk indexnya
8.  untuk menjalankan pakai perintah go run main.go
9.  install gorm untuk terhubung dengan database
10. jika ada error ketika import, gunakan perintah go get. contohnya go get gorm.io/driver/mysql
11. repository di dlm package book, bertanggung jawab terhadap entity Book/ tabel Books
12. soalnya dalam prakteknya, kita harus membatasi menggunakan db. (objek db) secara langsung, namun harus lewat repository
13. bisa dibilang repository adalah sebuah kode khusus yang dipakai sebagai perantara untuk ngambil data ke database (berhubungan dengan database, seperti find/create/save)
14. bisa dibilang juga repository itu seperti func yang ada di model pada laravel
15. berikut adalah layer dari database ke main dalam golang:
    mysql -> db (gorm) -> repository -> service -> handler -> main
16. layer service adalah yang bergubungan dengan bisnis/ logic
17. contoh bisnis/ logic misalnya fitur upload produk, update, dan menyimpan data ke database (yg berhubungan dengan logic atau aturan untuk mengolah database) 
18. jika ingin mengcustom response JSON, dapat menambahkan file response.go di package book
19. nah kalau handler baru bisa disebut seperti controller pada laravel 
20. contoh yg dilakukan handler, menampilkan json/menampilkan error/menampilkan custom response
21. biasanya handler dibilang endpoint, dan berhuubngan langsung dengan routes