FROM golang:1.18

RUN mkdir /app
WORKDIR /app

COPY . .

COPY .env .

RUN go get -d -v ./...

RUN go install -v ./...

# # hot reload
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT /go/bin/CompileDaemon --build="go build /app/main.go" --command=./main
