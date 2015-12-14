
FROM scratch
ENTRYPOINT ["/gonotfound"]

# Add the binary
ADD gonotfound /
EXPOSE 8080
