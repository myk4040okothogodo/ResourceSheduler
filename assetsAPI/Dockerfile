# Start with the golang base image
FROM golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

#Install git. (alpine image deosnt have git in it)
RUN apk update && apk add --no-cache git

#Set current worrking directory
WORKDIR   /app

#Note here: To avoid downloading dependencies every time we build an image
#Here, we are caching all the dependencies by first copying go.mod and go.sum file and downloading them
#to be used every time we buildd the image if the dependenciew are not changed

#Copy go mod and sum files
COPY go.mod  ./
COPY go.sum  ./

#Download all dependencies
RUN go mod download

#Now, copy the source code
COPY . .

# CGO_ENABLED is disabled for cross system compilation
# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ,/bin/main .

#Finally our multi-stage to build a small image, start a new stage from scratch
COPY --from=builder  /app/bin/main .

#Run executable
CMD ["./main"]


