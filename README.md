# API Gateway Service

## Giới thiệu

API Gateway Service là một dịch vụ cổng API mạnh mẽ và có khả năng mở rộng cao, được thiết kế để làm điểm trung gian giữa client và các microservices backend. Dịch vụ này cung cấp các tính năng như định tuyến request, reverse proxy, health check và xác thực.

## Tính năng chính

- ✅ **Reverse Proxy**: Chuyển tiếp request đến các service backend
- ✅ **Health Check**: Endpoint kiểm tra tình trạng hoạt động của service
- ✅ **Multi-language Support**: Hỗ trợ nhiều ngôn ngữ (Node.js, Python, Go)
- ✅ **RESTful API**: Thiết kế API theo chuẩn REST
- ✅ **Dễ dàng mở rộng**: Kiến trúc linh hoạt, dễ dàng thêm tính năng mới

## Cấu trúc thư mục

```
api-gateway-service/
├── src/
│   ├── index.js      # Node.js/Express implementation
│   ├── main.py       # Python/FastAPI implementation
│   └── main.go       # Go/Gin implementation
├── README.md
├── LICENSE
└── .gitignore
```

## Yêu cầu hệ thống

### Node.js Version
- Node.js >= 14.x
- npm hoặc yarn

### Python Version
- Python >= 3.8
- pip

### Go Version
- Go >= 1.18

## Cài đặt và chạy

### 1. Node.js/Express

```bash
# Cài đặt dependencies
npm install express axios

# Chạy server
node src/index.js
```

Server sẽ chạy trên `http://localhost:3000`

**Endpoints:**
- `GET /health` - Health check endpoint
- `GET /api/proxy` - Reverse proxy example

### 2. Python/FastAPI

```bash
# Cài đặt dependencies
pip install fastapi uvicorn httpx

# Chạy server
uvicorn src.main:app --reload --port 8000
```

Server sẽ chạy trên `http://localhost:8000`

**Endpoints:**
- `GET /health` - Health check endpoint
- `GET /api/proxy` - Reverse proxy example
- `GET /docs` - Auto-generated API documentation (Swagger UI)

### 3. Go/Gin

```bash
# Khởi tạo Go module (lần đầu)
go mod init github.com/RK-goldengate-co/api-gateway-service

# Cài đặt dependencies
go get github.com/gin-gonic/gin

# Chạy server
go run src/main.go
```

Server sẽ chạy trên `http://localhost:8080`

**Endpoints:**
- `GET /health` - Health check endpoint
- `GET /api/proxy` - Reverse proxy example

## Kiểm tra hoạt động

### Health Check

```bash
# Node.js
curl http://localhost:3000/health

# Python
curl http://localhost:8000/health

# Go
curl http://localhost:8080/health
```

### Reverse Proxy Test

```bash
# Node.js
curl "http://localhost:3000/api/proxy?url=https://jsonplaceholder.typicode.com/todos/1"

# Python
curl "http://localhost:8000/api/proxy?url=https://jsonplaceholder.typicode.com/todos/1"

# Go
curl "http://localhost:8080/api/proxy?url=https://jsonplaceholder.typicode.com/todos/1"
```

## Ví dụ Response

### Health Check Response
```json
{
  "status": "healthy",
  "timestamp": "2025-10-17T15:07:00Z",
  "service": "api-gateway"
}
```

### Proxy Response
```json
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
```

## Phát triển tiếp theo

Các tính năng có thể mở rộng:

- 🔐 Authentication & Authorization (JWT, OAuth2)
- ⚡ Rate Limiting & Throttling
- 📊 Logging & Monitoring
- 🔄 Load Balancing
- 🛡️ Security Headers & CORS
- 📈 Metrics & Analytics
- 🔌 WebSocket Support
- 💾 Caching Layer (Redis)

## Đóng góp

Mọi đóng góp đều được chào đón! Vui lòng tạo issue hoặc pull request.

## License

MIT License - xem file [LICENSE](LICENSE) để biết thêm chi tiết.

## Liên hệ

Nếu có bất kỳ câu hỏi nào, vui lòng tạo issue trên GitHub repository.
