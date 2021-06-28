FROM golang:1.14

WORKDIR /go/src/golocker

RUN go get github.com/githubnemo/CompileDaemon

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

# CMD ["go", "run", "main.go"]
ENTRYPOINT CompileDaemon -log-prefix=false -color=true -build="go build main.go" -command="./main"
