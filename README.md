<h1 align="center">
  Vehicles Rental
</h1>

<p align="center"><img src="https://i.ytimg.com/vi/kd-8mb6HfGA/maxresdefault.jpg" width="600px" alt="golang and mux" /></p>


## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/crownss/vehicles-rent
```

2. Install dependencies

```bash
go mod tidy
```

3. Migrate up (use this for auto migration)
```bash
go run . migrate -u
```

4. Migrate down (use this if you want delete all db make it from 0)
```bash
go run . migrate -d
```

5. Run the app

```bash
go run . serve
# or
go build -o start
# then
./start serve
# or
docker build -t <your_tag_name>:<your_version> .
# then
docker run -dit <your_tag_name>
```

4. Intruction
-   .env file: 
```
DBUSER=
DBPASSWORD=
DBNAME=
DBHOST=
DBPORT=
SSLMODE=
DBTIMEZONE=

RUN_HOST=
RUN_PORT=

SECRET_KEY=
```

5. Intruction


🌟 You are all set!

## 💻 Built with

-   [Golang](https://go.dev/)
-   [Gorilla/Mux](https://github.com/gorilla/mux): for routers framework
-   [Cobra](https://github.com/gorilla/mux): a commander for modern go cli interactions
-   [Postgres](https://www.postgresql.org/): for DBMS
-   [AutoMigrate](https://github.com/go-gormigrate/gormigrate): for auto migration using command
-   [Gorm](https://gorm.io/): for orm in golang
-   [Golang-jwt](https://github.com/golang-jwt/jwt): jwt for golang

<hr>
<p align="center">
Developed with ❤️ in Asia/Jakarta 	🇮🇩
</p>