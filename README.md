# :computer: Factorial API

## Intro
* API in Golang that calculates factorial of a given number
* This program has been made for N26

## Installation

 `Before anything make sure you have Go 1.12+ environment installed`.  
This project uses go modules
Link to docs : [https://golang.org/dl/](https://golang.org/dl/).  

```bash
$ git@github.com:cgstag/factorial.git
$ cd factorial
$ go build
$ go run factorial &
```

### With Docker

```bash
$ docker build -t cgstag/factorial .
$ docker run cgstag/factorial
```

## Usage

Then you can do HTTP requests (cURL or Postman should do it)

`curl http://localhost:5000/v1/factorial?n=5` should yield the following result
```json
{
  "n": 5,
  "result": 120
}
```

### Use it with docker

`curl http://172.17.0.2:5000/v1/factorial?n=5` 

(Make sure you replace the IP with the correct container IP. see [my gist about this](https://gist.github.com/cgstag/06d640dd196d634582d7979a00235a52))

## Tests

`go test ./...`






