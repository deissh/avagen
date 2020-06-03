FROM golang:1.10 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/avagen
COPY . ./
RUN go get -insecure ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server /avagen
# Add fronts
COPY RobotoMono-Regular.ttf RobotoMono-Regular.ttf

EXPOSE 8080
CMD ["/avagen", "serve", "--addr", "0.0.0.0:8080"]
