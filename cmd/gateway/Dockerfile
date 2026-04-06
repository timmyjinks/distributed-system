FROM golang:1.25

WORKDIR /app

COPY go.sum go.mod .
RUN go mod download

COPY . .
RUN go build -o /bin/app ./cmd/gateway

CMD ["/bin/app"]
