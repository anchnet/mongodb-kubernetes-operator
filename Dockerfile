
FROM golang:1.14 AS builder



ENV GO111MODULE=on
ENV GOPATH ""

WORKDIR /workspace
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn go mod download -x


ADD ./scripts/dev/templates/Dockerfile.operator ./scripts/dev/templates/Dockerfile.operator


# build the binary
# Copy the go source
COPY cmd/manager/main.go cmd/manager/main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY build/bin/ build/bin/

# Build the operator
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager cmd/manager/main.go

# build and second stage image if necessary
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=manager \
    USER_UID=1001 \
    USER_NAME=mongodb-kubernetes-operator

# Copy the operator binary and version manifest
WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /workspace/build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

USER ${USER_UID}
ENTRYPOINT ["/usr/local/bin/entrypoint"]

