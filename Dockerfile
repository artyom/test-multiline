FROM --platform=$BUILDPLATFORM public.ecr.aws/docker/library/golang:alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0 GOTOOLCHAIN=auto
COPY go.??? ./
RUN go mod download
COPY . .
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o main

FROM scratch
COPY --from=builder /app/main /
ENTRYPOINT ["/main"]
