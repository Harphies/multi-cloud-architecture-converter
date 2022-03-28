FROM golang:1.16-alpine AS builder 

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download  

COPY . . 

RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux go build -o ../../

# EXPOSE 4000

FROM scratch 

# Copy our static executable.
COPY --from=builder /app/main /app/api

COPY --from=builder /app/main /app/api 

ENTRYPOINT ["/app/api"]
