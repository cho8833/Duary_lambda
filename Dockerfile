FROM golang:1.22.1 as build
WORKDIR /getuser
COPY go.mod go.sum ./

COPY main.go .

RUN go build -tags lambda.norpc -o main main.go

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /getuser/main ./main
ENTRYPOINT [ "./main" ]