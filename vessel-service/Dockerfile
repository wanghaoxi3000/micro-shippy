FROM alpine:latest

RUN mkdir /app
WORKDIR /app

ADD vessel-service /app/vessel-service
RUN chmod +x /app/vessel-service

CMD ["./vessel-service"]