FROM alpine:latest

RUN mkdir /app
WORKDIR /app

ADD user-cli /app/user-cli
RUN chmod +x /app/consignment-cli

CMD ["./user-cli"]