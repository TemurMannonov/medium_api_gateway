# Multistage dockerfile (it is for minimizing docker image size)

# -- Build stage --
FROM golang:1.19.1-alpine3.16 AS builder

# workdir is the current working directory inside docker image 
# all dockerfile instructions will be executed inside workdir
WORKDIR /app  

# first dot means that copy everything from current folder (blog folder)
# second dot is the current working directory inside the image (/app folder)
COPY . .

# -- Run stage --
FROM alpine:3.16

WORKDIR /app

# copying main binary file to workdir
COPY --from=builder /app/main .

EXPOSE 8000

CMD ["/app/main"]