FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main ./main.go

EXPOSE 5432
CMD [ "/app/main" ]
