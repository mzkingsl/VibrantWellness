FROM golang:1.24-alpine

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server .

CMD ["./server"]
