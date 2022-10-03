# Build Go Binary
FROM golang:1.18 as build
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container
WORKDIR /usr/src/app
COPY . /usr/src/app

# Build the app binary
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Run the Go binary in Alpine
FROM alpine:3.14
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build /usr/src/app /app
WORKDIR /app
CMD ["./app"]
