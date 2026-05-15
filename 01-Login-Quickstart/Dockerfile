FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o out/auth0-go-web-app .

EXPOSE 3000

CMD ["/app/out/auth0-go-web-app"]
