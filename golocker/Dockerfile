# no work big dumb
FROM golang:1.14

WORKDIR /go/src/app
COPY . .

# ARG MYSQL_USER     
# ARG MYSQL_PASSWORD 
# ARG MYSQL_HOST     
# ARG MYSQL_TCP_PORT  
# ARG MYSQL_DATABASE 
# ARG GO_API_HOST     
# ARG GO_API_PORT     

# ENV MYSQL_USER      ${MYSQL_USER}
# ENV MYSQL_PASSWORD  ${MYSQL_PASSWORD}
# ENV MYSQL_HOST      ${MYSQL_HOST}
# ENV MYSQL_TCP_PORT  ${MYSQL_TCP_PORT}
# ENV MYSQL_DATABASE  ${MYSQL_DATABASE}
# ENV GO_API_HOST     ${GO_API_HOST}
# ENV GO_API_PORT     ${GO_API_PORT}


RUN go get -d -v ./...
RUN go install -v ./...
RUN go get github.com/codegangsta/gin

# RUN source .env

# CMD ["./app"]

CMD ["gin", "go", "main.go"]
# ENTRYPOINT ["sh", "bin/start.sh"]

# RUN gin go main.go --port 5000

# CMD ["gin"]

# CMD ["gin", "go", "main.go", "--port", "5000"]


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
