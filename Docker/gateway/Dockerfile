FROM alpine:latest
RUN mkdir /app
WORKDIR /app
EXPOSE 8080
ADD ./gateway /app/gateway
CMD ["./gateway"]