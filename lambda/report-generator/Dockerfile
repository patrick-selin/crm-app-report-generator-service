# syntax=docker/dockerfile:1.4

FROM public.ecr.aws/lambda/provided:al2 as builder

WORKDIR /src

COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /asset-output/bootstrap main.go

CMD [ "bootstrap" ]
