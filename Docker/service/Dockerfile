FROM alpine:latest
RUN mkdir /app
WORKDIR /app
EXPOSE 50051
ADD ./servicegrpc /app/servicegrpc
CMD ["./servicegrpc"]