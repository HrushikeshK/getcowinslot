FROM golang
COPY . .
EXPOSE 8080
RUN go get github.com/google/uuid
WORKDIR /go/src/
CMD go run *.go