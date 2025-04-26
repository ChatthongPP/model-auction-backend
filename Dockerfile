# ========== BUILD STAGE ==========
FROM golang:1.19-alpine AS build

WORKDIR /usr/src/app

ARG BUILD_TARGET=api

# Need to add tzdata to ensure correct timezone
RUN apk update && apk add --no-cache gcc && apk add --no-cache make bash tzdata

COPY go.mod .
COPY go.sum .

RUN go mod download

# Enforce UTC for all operations
ENV TZ=UTC

COPY . .

RUN GOCMD=GO111MODULE=on go build -ldflags="-s -w" -o "bin/main" "./cmd/${BUILD_TARGET}" 
RUN go get -d github.com/swaggo/swag/cmd/swag
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -h
RUN swag fmt
RUN swag init -g cmd/api/main.go --parseDependency --output docs/ 


# ========== FINAL SHIPPING STAGE ==========
FROM alpine:latest

ARG BUILD_TARGET=api

WORKDIR /usr/src/app

# Need to add tzdata to ensure correct timezone
RUN apk update && apk add --no-cache gcc && apk add --no-cache make bash tzdata

RUN echo "${BUILD_TARGET}"

# Enforce UTC for all operations
ENV TZ=UTC

COPY --from=build "/usr/src/app/bin/main" .

RUN touch .env

CMD ["./main"]




