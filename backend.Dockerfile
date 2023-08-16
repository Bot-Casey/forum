FROM golang:1.20 as builder
COPY ./* ./forum-backend
WORKDIR ./forum-backend
RUN go build -o /bin/backend cmd/main.go

FROM scratch
COPY --from=builder /bin/hello /bin/hello
CMD ["/bin/backend"]
