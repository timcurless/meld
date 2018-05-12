FROM golang:1.10.2-alpine3.7 AS build
LABEL maintainer=<tim.curless@thinkahead.com>

RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml /go/src/meld/
WORKDIR /go/src/meld/
RUN dep ensure --vendor-only

COPY . /go/src/meld/
WORKDIR /go/src/meld/cmd/meldsvc
RUN dep ensure && go build -o /bin/meld

FROM alpine:3.7
RUN apk add --no-cache curl && \
    curl -O https://releases.hashicorp.com/terraform/0.11.7/terraform_0.11.7_linux_amd64.zip && \
    unzip terraform_0.11.7_linux_amd64.zip && \
    mv terraform /usr/bin/terraform
COPY --from=build /bin/meld /bin/meld
ENTRYPOINT ["/bin/meld"]
CMD ["--help"]
