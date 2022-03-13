FROM golang:1.16-alpine

ARG EXEC_NAME=server

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
RUN [ "go", "build", "-o", "$EXEC_NAME", "./cmd/server" ]

EXPOSE 81
CMD [ "./$EXEC_NAME" ]
