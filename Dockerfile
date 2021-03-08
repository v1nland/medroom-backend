# Go image for building any microservice
FROM golang:alpine as builder
RUN apk --no-cache add git dep ca-certificates

# ENV
ENV GOBIN=$GOPATH/bin
ENV GO111MODULE="on"

# Create folder
RUN mkdir -p $GOPATH/src/microservice
WORKDIR $GOPATH/src/microservice

# Copy data
COPY . .
RUN go mod vendor
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $GOBIN/main ./app/main.go

# Runtime image with scratch container
FROM scratch
ARG VERSION
ENV VERSION_APP=$VERSION

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/ /app/

ENTRYPOINT ["/app/main"]
