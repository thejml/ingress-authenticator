FROM golang:latest as build

RUN useradd -m app
COPY main.go /home/app
WORKDIR /home/app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-w" -a -x main.go

FROM scratch
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /home/app/main /main

EXPOSE 8081
ENTRYPOINT ["/main"]
