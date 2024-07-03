FROM scratch
COPY go-credly /
ENTRYPOINT ["/go-credly"]
