# Hướng Dẫn Migration với Atlas & Ent

## 1. Tổng Quan

Quy trình migration giúp bạn quản lý thay đổi cấu trúc database một cách an toàn, có kiểm soát, phù hợp với nhiều môi trường (local, dev, prod...). Công cụ sử dụng:  
- **Ent**: ORM sinh code Go từ schema.
- **Atlas**: Quản lý migration, tạo file SQL, apply, rollback...

---

## 2. Quy Trình Cơ Bản

### Bước 1: Chỉnh sửa schema Ent

- Sửa đổi file schema trong thư mục `./ent/schema` để cập nhật model dữ liệu.

### Bước 2: Sinh lại code Ent

```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema
```
- Lệnh này sẽ sinh lại code Go từ schema mới.

### Bước 3: Tạo migration

#### a. Tạo migration tự động từ schema

```bash
atlas migrate diff <tên-migration> --env <env>
```
- `<tên-migration>`: Đặt tên cho migration, ví dụ: add-user-table.
- `--env <env>`: Chỉ định môi trường (local, dev, prod...).

#### b. Tạo migration rỗng (nếu cần chỉnh tay)

```bash
atlas migrate new <tên-migration>
```
- Tạo file migration rỗng để tự viết SQL.

### Bước 4: Áp dụng migration vào database

```bash
atlas migrate apply --env <env>
```
- Thực thi các migration chưa áp dụng vào database.

---

## 3. Các Lệnh Hữu Ích Khác

| Lệnh | Ý nghĩa |
|------|---------|
| atlas migrate status | Kiểm tra trạng thái migration (đã/ chưa apply) |
| atlas migrate show | Xem nội dung migration file |
| atlas migrate ls | Liệt kê các migration file |
| atlas migrate validate | Kiểm tra tính hợp lệ migration |
| atlas migrate hash | Tạo lại file hash kiểm tra integrity |
| atlas migrate diff | So sánh schema và database |
| atlas migrate import | Import migration từ tool khác |
| atlas migrate checkpoint | Tạo checkpoint trạng thái migration |
| atlas migrate edit <version> | Chỉnh sửa migration file |
| atlas migrate rm <version> | Xóa migration file |
| atlas migrate reset | Reset database (DANGER! Xóa toàn bộ data) |

---

## 4. Rollback

### Rollback migration
```bash
atlas migrate down 1 --env local
```

### Xóa file migration cũ
```bash
atlas migrate rm <version>
```
**Ví dụ:**
```bash
atlas migrate rm 20221115102552
```

### Sửa lại schema rồi gen lại Ent
```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema
```

### Kiểm tra schema khớp với database sau khi rollback
```bash
atlas migrate status --env local
```

### Validate migration status trước và sau rollback
```bash
atlas migrate validate --env local
```

---

## 5. Ví Dụ Quy Trình Đầy Đủ

1. Sửa file schema trong `./ent/schema`.
2. Sinh lại code Ent:
   ```bash
   go run entgo.io/ent/cmd/ent generate ./ent/schema
   ```
3. Tạo migration:
   ```bash
   atlas migrate diff add-user-table --env local
   ```
4. Áp dụng migration:
   ```bash
   atlas migrate apply --env local
   ```
5. Kiểm tra trạng thái:
   ```bash
   atlas migrate status --env local
   ```
6. Nếu cần rollback:
   ```bash
   atlas migrate down 1 --env local
   ```

---

## 6. Tham Khảo

- [Atlas Documentation](https://atlasgo.io/)
- [Ent Documentation](https://entgo.io/) 