FROM golang:1.19-alpine

WORKDIR /user/app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o entrypoint cmd/app/main.go 

RUN chmod +x entrypoint

EXPOSE 8080

ENTRYPOINT [ "/user/app/entrypoint" ]