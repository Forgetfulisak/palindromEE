FROM golang:1.19


WORKDIR /app
COPY . ./

# sqlite3 uses CGO
ENV CGO_ENABLED=1
ENV PORT=8080
ENV DATABASE_PATH="./database/db"


# Install dependencies
RUN go mod tidy

# Build webserver
RUN go build ./cmd/webserver

EXPOSE 8080/tcp

CMD [ "./webserver" ]