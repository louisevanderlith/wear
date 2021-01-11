FROM alpine:3.12.0

COPY cmd/cmd .

EXPOSE 8086

ENTRYPOINT [ "./cmd" ]