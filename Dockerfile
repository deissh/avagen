FROM golang:1.10 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/avagen
COPY . ./
RUN go get -insecure ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server ./server
# Add fronts
COPY wqy-zenhei.ttf wqy-zenhei.ttf

EXPOSE 8080
ENTRYPOINT ["./server"]
