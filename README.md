# Go URL Shortener Service

> **Author:** Phùng Xuân Dương  
> **GitHub:** [NhinhoD/GOLANG-NEWBIE](https://github.com/NhinhoD/GOLANG-NEWBIE)  
> **Tech Stack:** Golang, Gin Framework, GORM, PostgreSQL (Supabase)

## 1. Giới thiệu (Overview)
Đây là 1 bài tập Backend xây dựng dịch vụ rút gọn liên kết (URL Shortener).
* **Input:** URL dài (ví dụ: `https://www.google.com/search?q=golang`).
* **Output:** URL ngắn (ví dụ: `http://short.url/abc1234`).
* **Tính năng chính:**
    * Tạo link rút gọn.
    * Redirect về link gốc.
    * Đếm lượt click (Real-time).
    * API xem thống kê và danh sách link.

## 2. Hướng dẫn Cài đặt & Cấu hình (Setup Guide)

### Bước 1: Clone dự án

```bash
git clone [https://github.com/NhinhoD/GOLANG-NEWBIE.git](https://github.com/NhinhoD/GOLANG-NEWBIE.git)
cd GOLANG-NEWBIE

```
### Bước 2: Cấu hình biến môi trường (Environment)

Điền thông tin thật vào file .env:

 Dự án sử dụng Supabase. Nếu chạy trên mạng IPv4 (Wifi/4G Việt Nam), bắt buộc dùng Transaction Mode (Port 6543).

#### .env

### Thay [Your_Database_URL_Here] bằng đường dẫn kết nối DB của bạn
* DB_URL="Your_Database_URL_Here"

#### Domain trả về
* BASE_URL="http://localhost:8080/"

### Bước 3: Chạy ứng dụng

### Chạy Terminal dự án và gõ các lệnh:

### Tải dependencies
* go mod tidy

### Chạy server
* go run cmd/server/main.go

### Server sẽ khởi động tại: http://localhost:8080


## 3. API Documentation & Testing

Bạn có thể sử dụng **Postman** hoặc **cURL** để kiểm thử các API bên dưới.

| Chức năng        | Method | URL               | Body (JSON)                                      | Mô tả                                  |
|------------------|--------|-------------------|--------------------------------------------------|----------------------------------------|
| Tạo link rút gọn | POST   | `/shorten`        | `{ "original_url": "https://example.com/very/long/path/to/resource?param1=value1&param2=value2" }`       | Tạo link rút gọn mới                   |
| Redirect         | GET    | `/:code`          | *(Trống)*                                        | Chuyển hướng & tăng số lượt click      |
| Xem thông tin    | GET    | `/api/links/:code`| *(Trống)*                                        | Xem chi tiết link & thống kê click     |
| Danh sách link   | GET    | `/api/links`      | *(Trống)*                                        | Liệt kê tất cả link đã tạo             |

## 4. Thiết kế & Quyết định kỹ thuật 

* Database: PostgreSQL ()
 
 Supabase: Cung cấp PostgreSQL Managed Service giúp triển khai nhanh database bằng giao diện web mà không cần setup cầu kỳ.
  SQL đảm bảo dữ liệu không bị sai lệch khi có nhiều request cập nhật cùng lúc.

* thiết kế API kiểu RESTful với Gin

Gin Framework: Hiệu năng cao, cú pháp Router rõ ràng, dễ bảo trì.

REST: Chuẩn giao tiếp phổ biến, dễ dàng tích hợp với mọi nền tảng Client.

* Thuật toán sinh mã rút gọn.

Sử dụng phương pháp Random String Generation (Base62).Charset: [a-z][A-Z][0-9].

Độ dài: 6 ký tự.

Không gian mẫu: $62^6 \approx 56.8$ tỷ kết hợp.

* Xử lý conflict/duplicate

* xử lý theo cơ chế "Check-and-Retry" 

 - Cột short_code trong Database được đánh Index UNIQUE.

 - Trước khi Insert, hệ thống query kiểm tra xem mã đã tồn tại chưa.

 - Nếu trùng, thuật toán tự động sinh lại mã mới ngay lập tức.

## 5. Trade-offs

* 1. Random String vs. Auto-Increment ID

Quyết định: Chọn Random String.

Lý do: Bảo mật tốt hơn. Nếu dùng Auto-Increment (như 1, 2, 3...), có thể đoán được hệ thống có bao nhiêu link và dễ dàng cào dữ liệu.

Nhược điểm: Phải tốn thêm chi phí query DB để kiểm tra trùng lặp.

2. Atomic Update vs. Application Locking

Quyết định: Chọn Atomic Update (UPDATE table SET count = count + 1).

Lý do: Đẩy việc xử lý đồng thời xuống Database Engine, đảm bảo chính xác tuyệt đối 100% mà không làm chậm ứng dụng.

Nhược điểm: Code phụ thuộc vào cú pháp của SQL/ORM.

## 6. Challenges (Thử thách & Bài học)

Là một sinh viên mới ra trường với nền tảng **.NET**, khi tiếp cận **Golang** tôi gặp một số thách thức chính:

* **Khác biệt ngôn ngữ & tư duy**
  * Golang không theo OOP truyền thống như C#.
  * Cần làm quen với cách tổ chức code bằng struct, interface và composition.

* **Tổ chức project**
  * Không có sẵn Dependency Injection container.
  * Phải tự thiết kế structure (handler, service, repository) ngay từ đầu.

* **Xử lý đồng thời**
  * Bài toán đếm click cần đảm bảo chính xác khi có nhiều request cùng lúc.
  * Giải quyết bằng cách sử dụng **Atomic Update ở Database**.

Thông qua project này, tôi hiểu rõ hơn cách xây dựng một **Backend service đơn giản, hiệu quả và an toàn** với Golang.

## 6. Hạn chế ## 7. Hạn chế & Cải tiến (Limitations & Improvements)

Do giới hạn thời gian, dự án hiện tại là bản MVP. Để hoàn thiện và sẵn sàng cho Production, tôi sẽ bổ sung:

* **Caching (Redis):** Giảm tải cho Database bằng cách lưu các link truy cập thường xuyên vào bộ nhớ đệm (Redis), giúp phản hồi cực nhanh.
* **Unit Tests:** Viết kiểm thử tự động (dùng thư viện `testify`) để đảm bảo tính ổn định của code thay vì test thủ công.
* **Rate Limiting:** Giới hạn số lượt tạo link từ một IP trong một khoảng thời gian để ngăn chặn Spam và DDoS.
* **Docker:** Viết Dockerfile để đóng gói ứng dụng, giúp việc triển khai (Deploy) lên server dễ dàng và đồng bộ.
