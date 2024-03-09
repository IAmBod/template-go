FROM golang:1.22-alpine as build

WORKDIR /usr/src/app

ENV CGO_ENABLED=0
ENV GOCACHE=/root/.cache/go-build

COPY go.mod go.sum? ./

RUN go mod download && go mod verify

COPY . .

RUN --mount=type=cache,target="/root/.cache/go-build" go build -v -ldflags "-s -w" -o /bin/app

FROM gcr.io/distroless/static-debian12

COPY --from=build /bin/app /

CMD ["/app"]
