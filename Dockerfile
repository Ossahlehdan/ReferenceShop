FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 9090

ENTRYPOINT CompileDaemon --build="go build" --command=./Hadhel.ReferenceShop