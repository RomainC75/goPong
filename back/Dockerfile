FROM golang:1.21-rc-alpine3.17

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

RUN mkdir /app
COPY .. /app
WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

RUN go get github.com/spf13/cobra
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/spf13/viper
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm
RUN go install github.com/gzuidhof/tygo@latest
RUN go get github.com/gocolly/colly/v2

# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built
# ENTRYPOINT CompileDaemon -build="go build -o api" -command="./api"
ENTRYPOINT CompileDaemon -build="go build -o api" -command="echo'yeah' && ./api serve"