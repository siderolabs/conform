ARG GOLANG_IMAGE
FROM ${GOLANG_IMAGE}

ENV CGO_ENABLED 0
ENV GO111MODULES on

WORKDIR /conform
COPY ./ ./
RUN go mod download
RUN go mod verify
RUN go mod tidy
RUN go mod vendor

ARG TAG
ARG SHA
ARG BUILT
ENV GOOS linux
ENV GOARCH amd64
RUN go build -o /build/conform-${GOOS}-${GOARCH} -ldflags "-s -w -X \"github.com/autonomy/conform/cmd.Tag=${TAG}\" -X \"github.com/autonomy/conform/cmd.SHA=${SHA}\" -X \"github.com/autonomy/conform/cmd.Built=${BUILT}\"" .

ARG TAG
ARG SHA
ARG BUILT
ENV GOOS darwin
ENV GOARCH amd64
RUN go build -o /build/conform-${GOOS}-${GOARCH} -ldflags "-s -w -X \"github.com/autonomy/conform/cmd.Tag=${TAG}\" -X \"github.com/autonomy/conform/cmd.SHA=${SHA}\" -X \"github.com/autonomy/conform/cmd.Built=${BUILT}\"" .

ENV GOOS linux
ENV GOARCH amd64
COPY ./hack ./hack
RUN chmod +x ./hack/test.sh
RUN ./hack/test.sh --all

FROM scratch
COPY /build/conform-linux-amd64 /conform
ENTRYPOINT [ "/conform" ]
