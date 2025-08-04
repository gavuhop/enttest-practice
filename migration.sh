#!/bin/bash

set -e

# Sửa thông tin kết nối MySQL tại đây
DB_URL="mysql://user:password@localhost:3306/dbname"

# Kiểm tra tham số môi trường
ENV=${1:-local}
echo "==> Using environment: $ENV"

# Tạo migration mới
echo "==> Creating new migration..."
atlas migrate diff --env $ENV --dev-url $DB_URL

# Apply migration
echo "==> Applying migration..."
atlas migrate apply --env $ENV --url $DB_URL

echo "==> Migration completed successfully!"
echo "Note: Check migrations/ directory for versioned migration files"