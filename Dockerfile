FROM golang:buster AS builder

USER root

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

RUN go get github.com/gin-gonic/gin && \
    go get github.com/golang-jwt/jwt/v4 && \
    go get github.com/gorilla/handlers && \
    go get github.com/pkg/sftp && \
    go get golang.org/x/crypto

RUN go build -o app .

FROM golang:buster AS prod

USER root

RUN apt update -y && apt full-upgrade -y

RUN apt install openssl -y

WORKDIR $HOME/dist

COPY --from=builder /build/instrumented-app .

COPY --chmod=755 scripts/run-ssh /usr/local/bin/run-ssh

RUN mkdir -p $HOME/.ssh && \
	chmod 770 $HOME/.ssh

EXPOSE 5000

COPY ./scripts/starts.sh $HOME/starts.sh

RUN chmod +x $HOME/starts.sh && \
    chmod 775 -R $HOME/dist/app

ENTRYPOINT ["/opt/app-root/src/starts.sh"]
