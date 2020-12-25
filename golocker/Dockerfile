# no work big dumb
FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/codegangsta/gin

# CMD ["./app"]

CMD ["go", "run", "main.go"]

# CMD ["gin", "go", "main.go"]
