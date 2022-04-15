FROM golang:latest 
WORKDIR /app 
ADD ./ /app/ 
RUN go build ./main.go
ENTRYPOINT ["/app/entrypoint.sh"]
CMD ["sh"]
