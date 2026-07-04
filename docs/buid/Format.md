# Default Console Log Format

Đây là format mặc định được khuyến nghị cho thư viện **OTMC Logger**. Format này phù hợp cho các ứng dụng **CLI, Desktop, API, Server** và giúp log dễ đọc, nhất quán giữa các dự án.

## Mẫu hiển thị

```
2026-07-04 08:24:27.126 +07:00     execSQLFile()        db.go:114   | INFO  | 📝 Executed SQL file: migration/samples/03-instructions.sql
2026-07-04 08:24:27.126 +07:00     Initializer()        main.go:104 | INFO  | ✅ Database connection established.
2026-07-04 08:24:27.127 +07:00     Initializer()        main.go:112 | INFO  | 🤖 LLM Provider: groq
2026-07-04 08:24:27.128 +07:00     Runner()             main.go:154 | INFO  | 🌿 Running application...
2026-07-04 08:24:27.128 +07:00     Runner()             main.go:157 | INFO  | 🌐 Registering APIs...
```

---

# Cấu trúc log

```
┌──────────────────────────────┬──────────────────────┬────────────────┬────────┬───────────────────────────────┐
│ Timestamp                    │ Function             │ File:Line      │ Level  │ Message                       │
├──────────────────────────────┼──────────────────────┼────────────────┼────────┼───────────────────────────────┤
│ 2026-07-04 08:24:27.126      │ execSQLFile()        │ db.go:114      │ INFO   │ 📝 Executed SQL file...       │
└──────────────────────────────┴──────────────────────┴────────────────┴────────┴───────────────────────────────┘
```

## Ý nghĩa các cột

| Cột | Mô tả |
| --- | --- |
| Timestamp | Thời gian đầy đủ kèm múi giờ |
| Function | Tên hàm gọi logger |
| File:Line | File nguồn và số dòng |
| Level | Mức độ log (TRACE, DEBUG, INFO, WARN, ERROR, CRIT) |
| Message | Nội dung log |

---

# Căn lề tự động

Thư viện nên tự động căn lề các cột để log luôn thẳng hàng, giúp dễ đọc ngay cả khi có hàng nghìn dòng.

Ví dụ:

```
2026-07-04 08:24:27.126 +07:00     Runner()             main.go:154   | INFO  | 🌿 Running application...
2026-07-04 08:24:27.126 +07:00     Initializer()        main.go:104   | INFO  | ✅ Database connection established.
2026-07-04 08:24:27.127 +07:00     execSQLFile()        db.go:114     | INFO  | 📝 Executed SQL file...
```

---

# Thông tin tự động

Người dùng chỉ cần ghi:

```
logger.Info("🌐 Registering APIs...")
```

Logger sẽ tự động bổ sung:

- Timestamp
- Timezone
- Function
- File
- Line
- Log Level
- Padding và căn lề
- Màu sắc (nếu console hỗ trợ)

Người dùng không cần truyền các thông tin này.

---

# Màu sắc Console

Chỉ áp dụng khi ghi ra Console.

| Level | Màu |
| --- | --- |
| TRACE | Gray |
| DEBUG | Blue |
| INFO | Green |
| WARN | Yellow |
| ERROR | Red |
| CRIT | Bright Red |

> Chỉ nên tô màu phần **Level** để giữ log rõ ràng và dễ đọc. Khi ghi ra file, màu sắc sẽ tự động bị loại bỏ.

---

# Emoji

Logger không tự sinh emoji mà hiển thị đúng nội dung do ứng dụng truyền vào.

Ví dụ:

```
logger.Info("✅ Database connected")
logger.Info("🌐 Registering APIs")
logger.Warn("⚠️ Memory usage is high")
logger.Error("❌ Database connection failed")
```

---

# JSON Mode

Ngoài Console Format, thư viện nên hỗ trợ JSON để phục vụ Logging Platform hoặc ELK.

Ví dụ:

```
{
    "time": "2026-07-04T08:24:27.126+07:00",
    "level": "INFO",
    "function": "Initializer",
    "file": "main.go",
    "line": 112,
    "message": "LLM Provider: groq"
}
```

---

# Thiết kế mặc định

Format trên nên được tích hợp sẵn trong thư viện và sử dụng làm **Default Console Formatter**.

Người dùng chỉ cần:

```
logger.Info("Application started")
```

Thư viện sẽ tự động hiển thị theo đúng định dạng chuẩn mà không cần cấu hình thêm.

Nếu cần, người dùng vẫn có thể chuyển sang các formatter khác:

- Pretty Format *(mặc định)*
- Text Format
- JSON Format

---

# Nguyên tắc thiết kế

- Không yêu cầu người dùng cấu hình formatter để có log đẹp.
- Căn lề tự động giúp log dễ đọc.
- Tự động lấy Function, File và Line từ runtime.
- Hỗ trợ màu sắc trên Console.
- Không ghi mã màu ANSI vào file log.
- Tương thích với Structured Logging.
- Giữ API đơn giản, chỉ cần gọi `logger.Info()`, `logger.Warn()`, `logger.Error()` là đủ.