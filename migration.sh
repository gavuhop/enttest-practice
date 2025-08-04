#!/bin/bash

set -e

# Sửa thông tin kết nối MySQL tại đây
DB_URL="mysql://user:password@localhost:3306/dbname"

echo "==> Generate ent code"
go run entgo.io/ent/cmd/ent generate ./ent/schema

echo "==> Migration done!"
echo "Note: Use 'go run .' to test with auto-migration"