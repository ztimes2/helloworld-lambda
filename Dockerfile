FROM golang:alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY vendor/ vendor

RUN CGO_ENABLED=0 go build -o app -mod=vendor main.go

FROM amazon/aws-lambda-go

COPY --from=builder /app/app ${LAMBDA_TASK_ROOT}

CMD [ "app" ]