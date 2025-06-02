# Sistem Notifikasi

Sample aplikasi sistem notifikasi dengan menggunakan Redis pub/sub yang ada di golang.

Sistem ini akan memberikan suatu notifikasi untuk menginformasikan berbagai kegiatan atau pengumuman ke user.

Tapi, kali ini kita akan membuatnya dengan cukup mudah dengan menggunakan fitur Pub/Sub dari Redis.

Langkah-langkahnya:

```powershell
go mod init pub-sub-go
```

Install Dependency yang dibutuhkan

```powershell
go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/template/html/v2
go get -u github.com/redis/go-redis/v9
go get -u github.com/valyala/fasthttp
```

Buat folder

```powershell
md configs
md controllers
md views
```

Buat file docker-compose.yml di root folder, lakukan setting seperti dibawah ini

```docker
services:
  redis:
    image: redis:latest
    ports:
      - 6379:6379
    volumes:
      - ./_docker-data/redis:/data
```

Setelah dibuat, dari command prompt ketik:

```powershell
docker-compose up -d
```

Tunggu beberapa saat, karena docker sedang melakukan building container untuk redisnya.

Membuat suatu config yang akan menghubungkan aplikasi dengan redis database

```go
package configs

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_CHANNEL_NOTIFICATION = "notification"
)

var RDS *redis.Client
var RDS_CTX context.Context

func InitRedisClient() {
	RDS_CTX = context.Background()
	RDS = redis.NewClient(&redis.Options{
		Addr:     `localhost:6379`,
		Password: ``,
		DB:       0,
	})
}

```

Initiate koneksi ini main.go dengan cara:

```go
package main

import "pub-sub-go/configs"

func main() {
	// init redis client, menghubungkan aplikasi ke redis server
	configs.InitRedisClient()

	// cek apakan ada error ketika connect ke redis server
	if err := configs.RDS.Ping(configs.RDS_CTX).Err(); err != nil {
		panic(err)
	} else {
		println("Connected to redis server")
	}
}
```

Lakukan eksekusi main.go

```powershell
go run main.go
```

Seharusnya akan keluar message seperti dibawah ini jika koneksi dan konfigurasi nya sudah benar

![](C:\Users\erymn\AppData\Roaming\marktext\images\2025-06-02-11-18-01-image.png)

s


