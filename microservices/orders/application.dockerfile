FROM golang:1.20-bookworm

WORKDIR /usr/local/application

COPY . .

RUN apt-get update --yes
RUN apt-get upgrade --yes

RUN apt-get install --yes make

RUN export PATH="$PATH:$(go env GOPATH)/bin"

RUN make download
RUN make build

EXPOSE 3000

CMD ["./bin/main"]
