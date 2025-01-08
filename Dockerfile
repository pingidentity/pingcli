# Use the official Go image as the base for building
FROM golang:1.23 as builder

# GitHub tag
ARG RELEASE_VERSION

# Set the working directory
WORKDIR /app

# Download the source code from GitHub Releases
RUN curl -L https://github.com/pingidentity/pingcli/archive/refs/tags/v${RELEASE_VERSION}.tar.gz | tar xz --strip-components=1

# Build the pingcli binary (ensure static build)
RUN CGO_ENABLED=0 go mod tidy && CGO_ENABLED=0 go build -o /pingcli

# Debugging step: Check if the binary exists
RUN ls -al /pingcli

# Use Alpine as the runtime image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /pingcli /usr/local/bin/pingcli

# Set permissions for the binary
RUN chmod +x /usr/local/bin/pingcli

# Debugging step: Verify binary is in place
RUN ls -al /usr/local/bin/pingcli

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []
