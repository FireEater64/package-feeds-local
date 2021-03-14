FROM golang:1.16.2-buster as builder

# Create and change to the app directory.
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -mod=readonly -v -o logger

FROM gcr.io/distroless/base:nonroot

COPY --from=builder /app/logger /app/logger

# Run the web service on container startup.
CMD ["/app/logger"]
