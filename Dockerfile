FROM golang:buster
WORKDIR /home
COPY main.go .
RUN go build ./main.go
EXPOSE 80
CMD [ "./main" ]