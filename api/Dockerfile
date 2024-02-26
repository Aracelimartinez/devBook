FROM golang:1.22.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app/app

EXPOSE 8000

CMD ["/app/app"]
