FROM golang:1.16-alpine AS builder

# Move to working dir (/build)
WORKDIR /build

# Coppy and download dependencies using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Set necessary environment variables needed for our image
# and build the API server
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch

# Copy binary and config files from /build
# to the root folder of the scratch container
COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

# Export necessary port,
EXPOSE 3007

# Command to run when starting the container 
ENTRYPOINT ["/apiserver"]
