# Start from the official Go image.
FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o llm_checker ./cmd/llm_checker

######## Start a new stage from scratch #######
FROM node:14 as frontend_builder

WORKDIR /web

COPY web/package.json ./
RUN npm install

COPY web/public ./public
COPY web/src ./src

# Build the React app
RUN npm run build

######## Start a new stage from scratch #######
#FROM alpine:3.19.1  
FROM python:3.10

#RUN apk --no-cache add ca-certificates python3 py3-pip

# Install required Python packages

RUN pip install flair

WORKDIR /root

# Copy the Pre-built backend binary file from the previous stage
COPY --from=builder /app/llm_checker .

# Copy the Pre-built frontend files from the previous stage
COPY --from=frontend_builder /web/build ./web/build

# Copy python script
COPY scripts/compare_texts.py ./python/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the Python script
CMD ["./llm_checker"]