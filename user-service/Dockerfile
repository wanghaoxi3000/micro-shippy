FROM alpine:latest

RUN mkdir /app
WORKDIR /app

ADD user-service /app/user-service
RUN chmod +x /app/user-service

CMD ["./user-service"]