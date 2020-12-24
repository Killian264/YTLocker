# no work big dumb
FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gin", "go", "main.go", "--port", "5000"]


# FROM golang:alpine

# RUN mkdir /build

# WORKDIR /build

# ADD . /build

# COPY . /build

# # RUN go get -d -v ./build...

# # # Install the package
# # RUN go install -v ./build...

# RUN go build /build

# CMD ["gin", "go", "main.go", "--port", "5000"]
