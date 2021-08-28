######## Start a builder stage #######
FROM golang:1.16-alpine as builder

WORKDIR /app
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main cmd/oreilly-trial/main.go

######## Start a new stage from scratch #######
# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

WORKDIR /opt/
COPY --from=builder /app/bin/main .
USER 65532:65532

CMD ["./main"]
