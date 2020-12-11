FROM golang:1.15.6 AS builder

WORKDIR /go/src/github.com/joshrwolf/minigit

COPY . .
RUN go build -o bin/minigit

FROM registry.access.redhat.com/ubi8/ubi:8.3

RUN yum install -y git

COPY --from=builder /go/src/github.com/joshrwolf/minigit/bin/minigit /usr/local/bin/minigit

WORKDIR /etc/minigit/repos

RUN git clone --no-checkout https://github.com/rancher/rke2.git

ENTRYPOINT [ "/usr/local/bin/minigit" ]