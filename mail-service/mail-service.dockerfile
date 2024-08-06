
FROM alpine:latest

RUN mkdir /app

COPY mailerApp /app
COPY templates /templates


#COPY --from=builder /app/brokerApp /app

CMD ["/app/mailerApp"]