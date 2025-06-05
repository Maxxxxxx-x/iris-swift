ARG GO_VERSION=1.23.9

FROM golang:${GO_VERSION}

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build/prod

CMD ["./iris-swift"]
