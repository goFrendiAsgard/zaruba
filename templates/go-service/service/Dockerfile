FROM golang:latest
LABEL SERVICE servicename

# define environments
ENV SERVICE_NAME servicename

# Create app directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install app dependencies
RUN go mod download

# Bundle app source
COPY . .
RUN go build -o app .

EXPOSE 3000
CMD ./app