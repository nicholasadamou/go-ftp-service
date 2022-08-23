# go-ftp-service

The go-ftp-service is a Go Gin application that provides an API for sending files to a remote server via SFTP.

## API

### [sftp](app/routes.go)

**URL**: `/sftp`

**Method**: `POST`

**Consumes**: `application/json`

```text
{
	"fileName": "file.txt",
	"data": []
}
```

**Request Header(s)**: `Authorization`

#### Success Response

**Code**: `200 OK`

## Development

### Requirements

- [Docker](http://docker.com/)
- [Go](http://golang.org/doc/install)
- [Gin](https://gin-gonic.com/)

### Steps

Create a `.env` file with the following properties using the [`env.example`](.env.example) as an example:

```text
SSH_HOST=localhost
SSH_USERNAME=user
SSH_PASSWORD=pass
SSH_DESTINATION_FOLDER=/home/user/

ID_RSA=base64encodedid_rsa
ID_RSA_PUB=base64encodedid_rsa_pub
```

Install dependencies.

```bash
go modules download
```

Build the application.

```bash
go build -o app .
```

## Docker

To build the docker image for Upload Service, run the following command.

```bash
docker-compose build
```

Then to execute the docker container, run the following command.

```bash
docker-compose up -d
```

## License

© Nicholas Adamou.

It is free software, and may be redistributed under the terms specified in the [LICENSE] file.

[license]: LICENSE
