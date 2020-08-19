FROM dependencies AS builder
                           
WORKDIR /shorty

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -o /go/bin/shorty /shorty/cmd/shorty


FROM alpine:latest         

COPY --from=builder /go/bin/shorty /bin/shorty
COPY --from=builder /shorty/scripts /scripts

ENTRYPOINT ["/bin/shorty"] 
