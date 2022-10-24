FROM golang:1.17.12-alpine AS builder
WORKDIR /go/src/github.com/ddld93/promon/contractor
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/ddld93/promon/contractor ./
EXPOSE 3000
CMD ["./main"] 