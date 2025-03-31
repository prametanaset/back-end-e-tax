# ใช้ Golang image สำหรับการคอมไพล์แอปพลิเคชัน
FROM golang:1.24.1-alpine AS builder


# ตั้งค่าพาธในการทำงานใน container
WORKDIR /app

# คัดลอกไฟล์ Go โมดูลและไฟล์ที่จำเป็นทั้งหมด
COPY go.mod go.sum ./

# ติดตั้ง dependencies
RUN go mod tidy

# คัดลอกไฟล์ทั้งหมดในโปรเจกต์ไปยัง container
COPY . .

# คอมไพล์แอปพลิเคชัน
RUN go build -o main ./cmd

# ใช้ image ของ Alpine Linux เพื่อลดขนาดของอิมเมจ
FROM alpine:latest

# ติดตั้ง dependencies ที่จำเป็นสำหรับรันแอปพลิเคชัน
RUN apk --no-cache add ca-certificates

# ตั้งค่าพาธในการทำงาน
WORKDIR /root/

# คัดลอกไฟล์ที่คอมไพล์จากขั้นตอนก่อนหน้า
COPY --from=builder /app/main .

# เปิดพอร์ตที่แอปจะฟัง
EXPOSE 3000

# คำสั่งในการรันแอปพลิเคชัน
CMD ["./main"]
