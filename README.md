# Candhi REST API With Golang and MySQL 

This REST API provides data and service needed in candhi app.

## Installation ⚠️

### Golang 🐁

For the initial installation, please make sure you have go installed in your computer.

To check if Golang is installed :

```bash
go version
```

### install Candi REST API to your machine 🚀

To install Candi REST API to your machine, first download / clone this git repository.

go to your favorite terminal, then run this command to install all dependencies to your machine

```bash
go mod tidy
```

## Usage 🧨

To run Candhi REST API, run this command and make sure the port that needed is still available in your machine and your IP is registered to the MySQL connection instance in the Google Cloud Platform ☁️

```bash
go run main.go
```

if the application is running correctly you will see this similar message in your terminal log

```bash
Koneksi Berhasil !
Start Development server on port :3000
```


## API Endpoint Lists served in the cloud 🐝

These are API endpoints available categorized by its table in database

### Trivia Prefix Group 📜

Get all trivias, METHOD = GET
```bash
http://34.101.198.95:3000/api/trivias
```

Create Trivia, METHOD = POST
```bash
http://34.101.198.95:3000/api/trivia/create
```

Update Trivia, METHOD = PUT
```bash
http://34.101.198.95:3000/api/trivia/{id}/update
```

Delete Trivia, METHOD = DELETE
```bash
http://34.101.198.95:3000/api/trivia/{id}/delete
```


### Candi Prefix Group 🕍

Get all Candi, METHOD = GET
```bash
http://34.101.198.95:3000/api/candis
```

Create Candi, METHOD = POST
```bash
http://34.101.198.95:3000/api/candi/create
```

Update Candi, METHOD = PUT
```bash
http://34.101.198.95:3000/api/candi/{id}/update
```

Delete Candi, METHOD = DELETE
```bash
http://34.101.198.95:3000/api/candi/{id}/delete
```


### Article Prefix Group 🗞

Get all Article, METHOD = GET
```bash
http://34.101.198.95:3000/api/articles
```

Get all Article by category, METHOD = GET
```bash
http://34.101.198.95:3000/api/article/{category}
```

Create Article, METHOD = POST
```bash
http://34.101.198.95:3000/api/article/create
```

Update Article, METHOD = PUT
```bash
http://34.101.198.95:3000/api/article/{id}/update
```

Delete Article, METHOD = DELETE
```bash
http://34.101.198.95:3000/api/article/{id}/delete
```

## Contributing 👼🏿
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Disclaimer 🧟‍♀️
Made by cloud-computing team Faqqih and Ayi