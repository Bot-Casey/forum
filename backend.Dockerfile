FROM golang:1.20 as builder
COPY ./* /build/forum-backend
WORKDIR /build/forum-backend/cmd
ENV GOROOT=/build/forum-backend/cmd
RUN go build -o /bin/backend main.go

FROM scratch
COPY --from=builder /bin/hello /bin/hello
CMD ["/bin/backend"]
