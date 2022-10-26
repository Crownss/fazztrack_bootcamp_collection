<h1 align="center">
  Dark Market using Go and Gin-Gonic framework
</h1>

<p align="center"><img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="400px" alt="gin-gonic image" /></p>


## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/crownss/dark_market
```

2. Install dependencies

```bash
go mod tidy
```

3. Run the app

```bash
go run .
# or
go build -o start
# then
./start
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
-   Generate random secret key
```
package main;import("fmt";"time";"math/rand");var letters=[]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890");func randSeq(n int)string{b:=make([]rune, n);for i:=range b{b[i]=letters[rand.Intn(len(letters))]};return string(b)};func main(){rand.Seed(time.Now().UnixNano());fmt.Println(randSeq(6))}
```

change transaction in [`models/tx.go`](https://github.com/Crownss/dark_market/blob/master/models/tx.go)
default trasanction using BTC

change users requirements in [`models/users.go`](https://github.com/Crownss/dark_market/blob/master/models/users.go)
default using username and password for decentralization and anonymity

🌟 You are all set!

## 💻 Built with

-   [Golang](https://go.dev/)
-   [Gin Gonic](https://github.com/gin-gonic/gin): for framework and all his depedencies like cors etc.
-   [Postgres](https://www.postgresql.org/): for DBMS

<hr>
<p align="center">
Developed with ❤️ in Asia/Jakarta 	🇮🇩
</p>