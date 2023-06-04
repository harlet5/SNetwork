FROM golang:1.20.2-alpine
WORKDIR /app/social-network
RUN apk add --no-cache build-base
RUN apk update && apk add bash
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY . .
RUN go build -o social-network
EXPOSE 8000
#RUN chmod +x social-network
CMD ["./social-network"]
