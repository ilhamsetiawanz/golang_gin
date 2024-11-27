
# SIMPLE API USING GO AND MYSQL

Projek ini merupakan salah satu bagian dari proses pembelajaran saya dalam mempelajari Golang, khususnya dalam membuat RestAPI.

Projek ini menggunakan library Gin dan juga Gorm pada proses pembuatannya. diharapkan projek ini dapat menjadi contoh atau referensi dalam mempelajari golang nantinya.




## Instalasi Projek

lakukan clone pada project

```bash
  git clone https://github.com/ilhamsetiawanz/golang_gin.git
```

Masuk ke dalam file project

```bash
  cd golang_gin
```

Menambahkan Go Init

```bash
  go mod init golang_gin
```
Menginstall Library Gin dan Gorm
```bash
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql
```

## API Reference

#### Mengambil Semua Data Buku

```http
  GET http://localhost:8080/get-books
```

#### Mengambil Data Buku Berdasarkan ID

```http
  GET http://localhost:8080/get-books/:id
```

#### Memperbarui Data Buku

```http
    PUT http://localhost:8080/update-books/:id
```

#### Menghapus Data Buku

```http
    DELETE http://localhost:8080/delete-books/:id
```


## Kesimpulan

SIMPULKAN SENDIRI

