# Image to prepare tools

FROM golang:1.23-alpine AS tools

# Install make
RUN apk add make --no-cache

# Install buf
ARG BUF_VERSION=1.41.0

RUN BUF_TAR_NAME="buf-$(uname -s)-$(uname -m).tar.gz" &&\
    wget "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/${BUF_TAR_NAME}" &&\
    tar -xvzf ${BUF_TAR_NAME} -C "/usr/local" --strip-components 1 &&\
    rm -rf ${BUF_TAR_NAME}

# Install proto libraries
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest &&\
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest &&\
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# ------------------------------------

# Image to build protos and go binaries

FROM tools AS builder

# Install go packages
WORKDIR /src
COPY ./go.mod ./go.sum .

RUN go mod download && go mod verify

# Build go binary
COPY . .
RUN make gen && go build -v -o /bin/app ./cmd/server/main.go

# ------------------------------------

# Image to run app

FROM scratch

COPY --from=builder /bin/app /bin/app
USER 1000
CMD ["/bin/app"]
