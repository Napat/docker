# Docker Buildx

`Docker Buildx` เป็น CLI plugin สำหรับ Docker ที่ขยายความสามารถของคำสั่ง docker build โดยใช้ BuildKit เป็นเครื่องมือสร้างหลัก (build engine) ซึ่งติดตั้งมาพร้อมกับ Docker Desktop และ Docker Engine ตั้งแต่เวอร์ชัน 19.03 ขึ้นไป Buildx ช่วยให้สามารถสร้าง Docker image สำหรับหลายสถาปัตยกรรม (เช่น linux/amd64, linux/arm64, linux/arm/v7) จากคำสั่งเดียว และรองรับคุณสมบัติขั้นสูง เช่น

การสร้าง manifest list สำหรับ multi-architecture images ซึ่งช่วยให้ Docker เลือก image ที่เหมาะสมกับสถาปัตยกรรมของ host โดยอัตโนมัติ
การใช้ distributed caching เพื่อเพิ่มความเร็วในการสร้างโดยแชร์ cache ระหว่างเครื่องต่างๆ
การ export ผลลัพธ์การสร้างในรูปแบบต่างๆ เช่น OCI image tarballs
Buildx ทำงานโดยใช้ drivers ต่างๆ เพื่อกำหนดวิธีและสถานที่ในการสร้าง เช่น ใช้ Docker daemon ท้องถิ่นหรือ remote builders ซึ่งให้ความยืดหยุ่นสูงในการจัดการกระบวนการสร้าง

## Docker Buildx ใช้เพื่อแก้ปัญหาอะไร

ก่อนการมี Docker Buildx การสร้าง Docker image สำหรับหลายสถาปัตยกรรมเป็นกระบวนการที่ซับซ้อนและใช้เวลา โดยเฉพาะอย่างยิ่งเมื่อต้องรองรับ hardware ที่หลากหลาย เช่น x86_64 สำหรับเครื่อง PC และ ARM สำหรับ Raspberry Pi หรือเซิร์ฟเวอร์ cloud บางประเภท Docker Buildx ช่วยแก้ปัญหาดังต่อไปนี้

### การสร้าง image สำหรับหลายแพลตฟอร์ม (Multi-platform Builds)

ในอดีต การสร้าง image สำหรับสถาปัตยกรรมที่แตกต่างกันต้องใช้สภาพแวดล้อมการสร้างที่แยกกันหรือต้องกำหนดค่าการ cross-compile ที่ซับซ้อน Buildx ช่วยให้สามารถสร้าง image สำหรับหลายสถาปัตยกรรมจากคำสั่งเดียว โดยใช้เทคนิคเช่น QEMU สำหรับการ emulate สถาปัตยกรรมที่ไม่ใช่ native หรือใช้ native nodes สำหรับการสร้างบน hardware จริง
ตัวอย่างเช่น คุณสามารถสร้าง image สำหรับ linux/amd64, linux/arm64, linux/arm/v7 ได้พร้อมกันด้วยคำสั่งเดียว

### ประสิทธิภาพการสร้าง (Build Performance)

Buildx รองรับ distributed caching ซึ่งช่วยลดเวลาในการสร้างโดยการแชร์ cache ระหว่างการสร้างบนเครื่องต่างๆ นอกจากนี้ยังรองรับ parallel multi-stage builds ซึ่งเหมาะสำหรับโครงการที่มีขั้นตอนการสร้างที่ซับซ้อน
การใช้ BuildKit เป็น backend ช่วยเพิ่มประสิทธิภาพโดยลดการสร้างซ้ำของขั้นตอนที่ไม่เปลี่ยนแปลง

### ความยืดหยุ่นในการกำหนดสภาพแวดล้อมการสร้าง

Buildx สามารถกำหนดได้ว่าจะสร้าง image ที่ไหนและอย่างไร โดยใช้ drivers ต่างๆ เช่น docker driver สำหรับการสร้างบน Docker daemon ท้องถิ่นหรือ kubernetes driver สำหรับการสร้างบน Kubernetes cluster
นอกจากนี้ยังสามารถ export ผลลัพธ์การสร้างในรูปแบบต่างๆ เช่น OCI image tarballs ซึ่งเป็นประโยชน์สำหรับการใช้งานในสถานการณ์ที่หลากหลาย เช่น CI/CD pipelines

### การจัดการ Manifest List

Buildx ช่วยสร้าง manifest list ซึ่งเป็นรายการที่รวม manifest ของ image สำหรับแต่ละสถาปัตยกรรม ทำให้ Docker สามารถเลือก image ที่เหมาะสมกับสถาปัตยกรรมของ host โดยอัตโนมัติ ซึ่งเป็นประโยชน์เมื่อ deploy บน cloud หรือ edge devices ที่มีสถาปัตยกรรมหลากหลาย

## ใช้ร่วมกับ Golang

GoLang เป็นภาษาโปรแกรมที่เหมาะสมอย่างยิ่งสำหรับใช้งานกับ Docker Buildx เนื่องจาก Go มีความสามารถในการ cross-compile ที่ดีเยี่ยม หมายความว่าสามารถสร้าง binary สำหรับสถาปัตยกรรมที่แตกต่างกันได้โดยไม่ต้องมี hardware ของสถาปัตยกรรมนั้นๆ อยู่จริง การใช้ Buildx กับ GoLang จึงช่วยให้สามารถสร้าง Docker image ที่รองรับหลายแพลตฟอร์มได้อย่างมีประสิทธิภาพ

### วิธีการใช้งาน

เขียน Dockerfile สำหรับโครงการ Go

ใช้ multi-stage build เพื่อแยกขั้นตอนการสร้าง binary ของ Go ออกจากขั้นตอนการสร้าง image สุดท้าย
ขั้นตอนแรกใช้ image ของ Go (เช่น golang:1.19) เพื่อสร้าง binary สำหรับสถาปัตยกรรมเป้าหมาย โดยใช้ตัวแปรเช่น TARGETARCH และ TARGETOS
ขั้นตอนที่สองใช้ image เล็กๆ เช่น alpine เพื่อ copy binary ที่สร้างได้เข้าไป

ตัวอย่าง Dockerfile สำหรับ Go ทั่วไป:

```Dockerfile
# syntax=docker/dockerfile:1
FROM golang:1.19 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o /app/myapp

FROM alpine
COPY --from=build /app/myapp /myapp
CMD ["/myapp"]
```

สำหรับโครงการที่ใช้ CGO (ซึ่งเกี่ยวข้องกับการ compile C code) ต้องติดตั้ง cross-compiler สำหรับสถาปัตยกรรมเป้าหมาย เช่น gcc-aarch64-linux-gnu สำหรับ arm64 และปรับ Dockerfile ดังนี้:

ตัวอย่าง Dockerfile สำหรับ Go กับ CGO:

```Dockerfile
# syntax=docker/dockerfile:1
FROM golang:1.19 AS build
ARG TARGETARCH
WORKDIR /app
COPY . .
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        CC=aarch64-linux-gnu-gcc \
        CC_FOR_TARGET=gcc-aarch64-linux-gnu; \
    fi \
    && CGO_ENABLED=1 GOOS=linux GOARCH=${TARGETARCH} CC=$CC CC_FOR_TARGET=$CC_FOR_TARGET go build -a -ldflags '-extldflags "-static"' -o /app/myapp main.go

FROM alpine
COPY --from=build /app/myapp /myapp
CMD ["/myapp"]
```

### ใช้ Buildx เพื่อสร้าง image สำหรับหลายแพลตฟอร์ม

ก่อนอื่นต้องตั้งค่า Buildx builder (ถ้ายังไม่มี) ด้วยคำสั่ง docker buildx create --use
จากนั้นใช้คำสั่ง docker buildx build โดยระบุแพลตฟอร์มที่ต้องการสร้าง เช่น linux/amd64, linux/arm64, linux/arm/v7
ใช้ flag --push เพื่อ push image ไปยัง registry โดยตรง
ตัวอย่างคำสั่งสร้าง:

สำหรับ Go ทั่วไป

```sh
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t myregistry/myapp:latest --push .
```

สำหรับ Go กับ CGO

```sh
docker buildx build --platform linux/amd64,linux/arm64 -t myregistry/mycgoapp:latest --push .
```

### การใช้งาน image ที่สร้างได้

หลังจากสร้างและ push image ไปยัง registry คุณสามารถ run image บนสถาปัตยกรรมที่แตกต่างกันได้โดยอัตโนมัติ
Docker จะเลือก variant ของ image ที่เหมาะสมกับสถาปัตยกรรมของ host

ตัวอย่างการ run

บน amd64: `docker run --rm -t myregistry/myapp:latest`
บน arm64 (emulated): `docker run --platform linux/arm64 --rm -t myregistry/myapp:latest`
