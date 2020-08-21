FROM golang:latest
LABEL SERVICE servicename

# define environments
ENV SERVICE_NAME servicename

# Create app directory
WORKDIR /app

# Install app dependencies
COPY go.mod go.sum ./
RUN go mod download

# Bundle app source
COPY . .
RUN go build -o app .

EXPOSE 3000
CMD ./app
