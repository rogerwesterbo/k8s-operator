FROM golang as builder
LABEL stage=builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o dist/operator -ldflags '-w -extldflags "-static"' main.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder --chown=1001:1001 /build/dist/operator /bin/operator
ENTRYPOINT [ "operator" ]
