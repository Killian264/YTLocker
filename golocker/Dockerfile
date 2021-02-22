FROM golang:1.14

WORKDIR /go/src/golocker

COPY . .

RUN go get
RUN go get github.com/githubnemo/CompileDaemon

# CMD ["go", "run", "main.go"]
ENTRYPOINT CompileDaemon -log-prefix=false -color=true -build="go build main.go" -command="./main"
