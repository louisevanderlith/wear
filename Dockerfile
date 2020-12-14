FROM scratch

COPY cmd/cmd .

EXPOSE 8086

ENTRYPOINT [ "./cmd" ]