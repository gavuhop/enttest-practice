# Ent Project with Atlas Migration

## Migration với Versioning

Dự án này sử dụng Atlas để quản lý database migration với versioning đàng hoàng.

### Cấu hình

1. **File cấu hình**: `atlas.hcl` - định nghĩa các environment (local, dev, prod)
2. **Migration files**: Thư mục `migrations/` chứa các file migration versioned
3. **Scripts**: `scripts/migrate.sh` - tool migration nâng cao

### Sử dụng

#### Script cơ bản (migration.sh)
```bash
# Apply migration cho environment local
./migration.sh

# Apply migration cho environment khác
./migration.sh dev
./migration.sh prod
```

#### Script nâng cao (scripts/migrate.sh)
```bash
# Xem trạng thái migration
./scripts/migrate.sh local status

# Tạo migration mới từ schema changes
./scripts/migrate.sh local diff

# Apply migrations
./scripts/migrate.sh local apply

# Tạo migration với tên tùy chỉnh
./scripts/migrate.sh local new add_user_table

# Validate migrations
./scripts/migrate.sh local validate
```

### Workflow Migration

1. **Thay đổi schema**: Chỉnh sửa file trong `ent/schema/`
2. **Tạo migration**: `./scripts/migrate.sh local diff`
3. **Review migration**: Kiểm tra file trong `migrations/`
4. **Apply migration**: `./scripts/migrate.sh local apply`

### Cấu trúc Migration

- `migrations/` - Chứa các file migration versioned
- `migrations/atlas.sum` - Hash tracking cho migrations
- `atlas.hcl` - Cấu hình environments
- `scripts/migrate.sh` - Tool migration nâng cao

### Lưu ý

- Luôn backup database trước khi apply migration
- Test migration trên environment dev trước khi apply lên prod
- Review migration files trước khi apply
