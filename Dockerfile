FROM golang:1.10 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/avagen
COPY . ./
RUN go get -insecure ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server /avagen
# Add fronts
COPY Cousine-Bold.ttf Cousine-Bold.ttf

EXPOSE 8080
CMD ["/avagen", "-address", "127.0.0.1:8080"]
