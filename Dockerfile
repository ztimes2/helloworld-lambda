FROM golang:alpine as builder

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o app main.go

FROM amazon/aws-lambda-go

COPY --from=builder app ${LAMBDA_TASK_ROOT}

CMD [ "app" ]