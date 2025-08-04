#!/bin/bash

set -e

# Cấu hình mặc định
ENV=${1:-local}
ACTION=${2:-apply}

# Sửa thông tin kết nối MySQL tại đây
DB_URL="mysql://user:password@localhost:3306/dbname"

echo "==> Atlas Migration Tool"
echo "Environment: $ENV"
echo "Action: $ACTION"
echo ""

case $ACTION in
  "diff")
    echo "==> Creating new migration..."
    atlas migrate diff --env $ENV --dev-url $DB_URL
    echo "Migration file created in migrations/ directory"
    ;;
  "apply")
    echo "==> Applying migrations..."
    atlas migrate apply --env $ENV --url $DB_URL
    echo "Migrations applied successfully!"
    ;;
  "status")
    echo "==> Migration status:"
    atlas migrate status --env $ENV --url $DB_URL
    ;;
  "hash")
    echo "==> Computing migration hash..."
    atlas migrate hash --env $ENV
    ;;
  "validate")
    echo "==> Validating migrations..."
    atlas migrate validate --env $ENV --url $DB_URL
    ;;
  "new")
    echo "==> Creating new migration with name..."
    MIGRATION_NAME=${3:-"new_migration"}
    atlas migrate diff --env $ENV --dev-url $DB_URL $MIGRATION_NAME
    ;;
  *)
    echo "Usage: $0 [environment] [action] [migration_name]"
    echo ""
    echo "Environments: local, dev, prod"
    echo "Actions:"
    echo "  diff     - Create new migration from schema changes"
    echo "  apply    - Apply pending migrations"
    echo "  status   - Show migration status"
    echo "  hash     - Compute migration hash"
    echo "  validate - Validate migrations"
    echo "  new      - Create new migration with custom name"
    echo ""
    echo "Examples:"
    echo "  $0 local apply"
    echo "  $0 dev diff"
    echo "  $0 local new add_user_table"
    ;;
esac 