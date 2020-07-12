FROM golang:latest

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

# Install dependencies (bad way)
RUN go get github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/calculate
RUN go get github.com/KHYehor/gRPCGolang/src/grpc/grpcModules/health
RUN go get github.com/KHYehor/gRPCGolang/src/modules/health
RUN go get github.com/KHYehor/gRPCGolang/src/modules/server
RUN go get google.golang.org/grpc

RUN go build -o server
CMD ["./server"]

