FROM golang:1.24.4

RUN echo "Setting up Go environment..."
WORKDIR /app
COPY go.mod go.sum ./

RUN echo "Downloading dependencies..."
RUN go mod download

RUN echo "Copying source files..."
COPY . .

RUN echo "Building the application..."
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-storage

EXPOSE 8080

RUN echo "RUNNING the application..."
CMD ["/docker-storage"] 