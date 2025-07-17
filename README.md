# elephpants-api

This simple API handle a list of Elephpants.
It alllows to:
- list the existing Elephpants
- display the information about a Elephpant
- create a new Elephpant
- delete a Elephpant
- update the path and the URL of a Elephpant

<img src="https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant.png" alt="ElePHPant.png" width="300"/> <img src="https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png" alt="back-to-the-future-v2.png" width="300"/>

## Gitpod integration

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scraly/elephpants-api.git)

## Docker image

Gophers API is available in [Docker Hub](https://hub.docker.com/r/scraly/elephpants-api).

### Run the ElePHPants API with Docker

```bash
docker run -p 8080:8080 scraly/elephpants-api:linux-amd64
```

## How to install 

### Prerequisites

Install Go in 1.16 version minimum.  
Install [Taskfile](https://taskfile.dev/#/installation) (optional):

```bash
brew install go-task/tap/go-task
```

Install go-swagger:

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
swagger version
```

### Build 

``` 
go build -o bin/elephpants-api internal/main.go

// or 

task build
```

### Run app 

``` 
go run internal/main.go

// or 

task run
```

### Serve Swagger UI 

This will open you browser on Swagger UI

``` 
task swagger.serve
```

### Test the API

* Get all Gophers:

```bash
curl localhost:8080/elephpants
```

Response:

```bash
[{"displayname":"ElePHPant.png","name":"ElePHPant","url":"https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant.png"},{"displayname":"ElePHPant harry potter","name":"ElePHPant-harry-potter","url":"https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant-harry-potter.png"}]
```

* Get a Elephpant with the input name

```bash
curl "localhost:8080/elephpant?name=ElePHPant"
```

Response:

```bash
{"displayname":"ElePHPant.png","name":"ElePHPant","url":"https://raw.githubusercontent.com/scraly/elephpants/refs/heads/main/ElePHPant.png"}
```

/!\ Returns a 404 HTTP Error Code if a Elephpant have not been found for the given name.

* Get directly the image of an ElePHPant

```bash
$ curl "localhost:8080/elephpant/image?name=ElePHPant"
```

Response:

```bash
$ curl "localhost:8080/elephpant/image?name=ElePHPant" --output o.png

  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 72529    0 72529    0     0   739k      0 --:--:-- --:--:-- --:--:--  745k
```

* Add a new Elephpant

```
curl -X POST localhost:8080/elephpant \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","displayname":"Yoda Gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}'  
```

Response:

```bash
{"displayname":"Yoda Gopher.png","name":"yoda-gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}
```

Add another Elephpant:

```
curl -X POST localhost:8080/elephpant \
   -H "Content-Type: application/json" \
   -d '{"name":"jurassic-park","displayname":"Gopher Park","url":"https://raw.githubusercontent.com/scraly/gophers/main/jurassic-park.png"}'  
```

* Delete a Elephpant

```bash
curl -X DELETE "localhost:8080/elephpant?name=5th-element"
```

* Update a Elephpant

```bash
curl -X PUT localhost:8080/elephpant \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","displayname":"El mejor Yoda Gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}' 
```

Response:

```bash
{"displayname":"El mejor Yoda Gopher","name":"yoda-gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}
```

## Build docker image

* Build a docker image for our current/host platform:

```
DOCKER_BUILDKIT=1 docker build -t elephpants-api .
```

* Build for GitPod (linux/amd64) and push to the Docker Hub:

```
docker buildx build --platform linux/amd64 -t scraly/elephpants-api:linux-amd64 . --push
```

## GoReleaser

Install the CLI:

```
brew install goreleaser
```

Generate .goreleaser.yml (the firt time):
```
goreleaser init
```

Release:

```
goreleaser release --snapshot --skip-publish --rm-dist
```

## Notes

This API use [go-swagger](https://goswagger.io/install.html)
