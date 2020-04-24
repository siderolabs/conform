ARG GOLANG_IMAGE
FROM ${GOLANG_IMAGE} AS common
ENV CGO_ENABLED 0
ENV GO111MODULES on
WORKDIR /conform
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify
COPY ./ ./
RUN go list -mod=readonly all

FROM common AS build
ARG TAG
ARG SHA
ARG BUILT
ENV GOOS linux
ENV GOARCH amd64
RUN go build -o /conform-${GOOS}-${GOARCH} -ldflags "-s -w -X \"github.com/talos-systems/conform/cmd.Tag=${TAG}\" -X \"github.com/talos-systems/conform/cmd.SHA=${SHA}\" -X \"github.com/talos-systems/conform/cmd.Built=${BUILT}\"" .

ARG TAG
ARG SHA
ARG BUILT
ENV GOOS darwin
ENV GOARCH amd64
RUN go build -o /conform-${GOOS}-${GOARCH} -ldflags "-s -w -X \"github.com/talos-systems/conform/cmd.Tag=${TAG}\" -X \"github.com/talos-systems/conform/cmd.SHA=${SHA}\" -X \"github.com/talos-systems/conform/cmd.Built=${BUILT}\"" .

FROM common AS test
ENV GOOS linux
ENV GOARCH amd64
COPY ./hack ./hack
RUN chmod +x ./hack/test.sh
RUN ./hack/test.sh --all

FROM alpine:3.11 as ca-certificates
RUN apk add --update --no-cache ca-certificates

FROM scratch AS image
COPY --from=ca-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /conform-linux-amd64 /conform
ENTRYPOINT [ "/conform" ]
CMD [ "enforce" ]
