FROM golang:1.23-alpine

# Cài đặt Atlas
RUN go install ariga.io/atlas/cmd/atlas@latest

# Cài đặt các dependencies cần thiết
RUN apk add --no-cache bash

# Tạo thư mục làm việc
WORKDIR /app

# Copy source code
COPY . .

# Cài đặt Go dependencies
RUN go mod download

# Default command
CMD ["./migration.sh"] 