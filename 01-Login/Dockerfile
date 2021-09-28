FROM golang:1.17-alpine3.14

# Define current working directory
WORKDIR /01-Login

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add sources
COPY . .

RUN go build -o out/auth0-go-web-app .

# Expose port 3000 for our web app binary
EXPOSE 3000

CMD ["/01-Login/out/auth0-go-web-app"]
