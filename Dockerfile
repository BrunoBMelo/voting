FROM golang:1.17-alpine as build
WORKDIR /app
COPY go.mod go.sum  /app/
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -installsuffix cgo -o /cmd

FROM gcr.io/distroless/static
COPY --from=build /cmd/ /
EXPOSE 80
CMD ["/cmd"]