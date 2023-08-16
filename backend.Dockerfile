FROM golang:1.20 as builder
COPY ./* /build/forum-backend
WORKDIR /build/forum-backend/cmd
ENV GOROOT=/build/forum-backend/cmd
RUN go build -o ./backend.out .

FROM scratch
COPY --from=builder /build/forum-backend/cmd/backend.out /bin/hello
CMD ["/bin/backend.out"]
