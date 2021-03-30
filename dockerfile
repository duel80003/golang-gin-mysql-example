# FROM golang:1.16.2-alpine
# WORKDIR /backend
# ADD . /backend
# RUN go mod download
# RUN cd /backend && go build -o api
# EXPOSE 8080
# ENTRYPOINT [ "./api" ]

FROM scratch
WORKDIR /backend
COPY . /backend
EXPOSE 8080
CMD ["./api"]