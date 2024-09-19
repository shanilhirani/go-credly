# Build stage: Use a base image with certificates to build the final image
FROM alpine:latest AS base

# Copy the CA certificates from Alpine to a separate build stage
RUN apk --no-cache add ca-certificates

# Final stage: Use scratch as the base image
FROM scratch

# Copy the compiled Go binary into the container
COPY go-credly /

# Copy the CA certificates from the previous build stage
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Set the binary as the entrypoint
ENTRYPOINT ["/go-credly"]
