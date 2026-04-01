# build stage
FROM golang:1.25 AS builder

ARG VERSION=dev

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-X github.com/michaeljmartin28/minikms/package/version.Version=$VERSION" -o /bin/minikms ./cmd/main.go

# runtime stage
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR  /app

COPY --from=builder /bin/minikms /app/minikms

EXPOSE 8080
EXPOSE 9090

USER nonroot:nonroot

ENTRYPOINT ["/app/minikms"]
