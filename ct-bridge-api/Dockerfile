FROM golang:1.19-alpine3.18

# Set up apk dependencies
ENV PACKAGES make git libc-dev bash gcc linux-headers eudev-dev curl ca-certificates

# Set up app flags
ENV AP_SCHEME http
ENV AP_HOST 0.0.0.0
ENV AP_PORT 8080
ENV AP_CLEANUP_TIMEOUT 4s

# Set working directory for the build
WORKDIR /opt/app

# Add source files
COPY . .

# Install minimum necessary dependencies, remove packages
RUN apk add --no-cache $PACKAGES && \
    make build-linux && \
    rm -rf ./vendor

# Run the app
#CMD ./build/ct-evm-compatible-bridge-api --scheme=$AP_SCHEME --host=$AP_HOST --port=$AP_PORT --cleanup-timeout=$AP_CLEANUP_TIMEOUT --config-file=config.json
