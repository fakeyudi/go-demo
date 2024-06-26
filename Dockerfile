FROM golang:1.22.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin .

EXPOSE 8084

ENTRYPOINT [ "/app/bin" ]