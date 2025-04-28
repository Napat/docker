# Docker with non-root user and cap_net_bind_service

## TLDR

- การตั้งค่า USER nonroot ใน Dockerfile ช่วยเพิ่มความปลอดภัยให้กับ container
- ตั้งค่า `setcap 'cap_net_bind_service=+ep' <filename>` เพื่อแก้ปัญหาให้ nonroot user สามารถ bind port ที่ต่ำกว่า 1024 ได้

## Non-root user

การตั้งค่า Dockerfile ให้ container ทำงานใน identity ของ non-root user ที่มี privilege ต่ำกว่า root เป็นการเพิ่มความปลอดภัยให้กับ container สามารถทำได้ดังนี้

``` Dockerfile
RUN adduser -D -u 1000 nonroot
USER nonroot
```

หลังจากตั้งค่า USER ใน Docker container เป็น non-root แล้ว ถ้าเราใช้งานบนเครื่อง mac หรือ windows ก็จะสามานถใช้งานได้ทันทีไม่มีปัญหาอะไร

แต่เมื่อนำ image ไปใช้งานในเครื่อง Linux ตั้งแต่ kernel kernel version 2.6.24 ขึ้นไปซึ่งมีระบบการจัดการ filesystem capabilities อาจจะเกิดปัญหาขึ้นได้
**ถ้ามีการ bind port ที่ต่ำกว่า 1024 ลงมา(privilege port)** เนื่องจากการเปิดพอร์ตกลุ่มนี้จำเป็นต้องใช้ privilege ของ root

### วิธีการแก้ไข

เราสามารถตั้งค่า capability ให้กับ binary นั้น ๆ ได้โดยใช้คำสั่ง `setcap 'cap_net_bind_service=+ep' <filename>` เพื่อตั้งค่าให้ non-root user สามารถ bind กับ port ต่ำกว่า 1024
ตัวอย่าง Dockerfile สำหรับ alpine linux เขียนได้ดังนี้

``` Dockerfile
RUN apk add --no-cache libcap && \
    setcap 'cap_net_bind_service=+ep' /app/main && \
    apk del libcap && \
    adduser -D -u 1000 nonroot
USER nonroot
```

การกำหนดสิทธิ์ให้กับไฟล์เพื่อให้มีสิทธิ์ในการเชื่อมต่อ port ที่มีหมายเลขน้อยกว่า 1024 ได้โดยไม่ต้องใช้สิทธิ์ของ root  ช่วยทำให้ง่ายต่อการเรียกใช้งานโปรแกรมที่ต้องการ  
อย่างไรก็ตาม การกำหนดสิทธิ์นี้ต้องใช้ความระมัดระวังเพราะอาจเป็นช่องทางให้โจมตีได้เช่นกัน **ควรใช้คำสั่งดังกล่าวด้วยความระมัดระวังและหลีกเลี่ยงการให้สิทธิ์แก่ไฟล์ที่ไม่น่าเชื่อถือ**

## How to run

To run the application, make sure you have Docker and Docker Compose installed on your machine.

1. Clone the repository and navigate to the project directory.

    ``` shell
    git clone https://github.com/Napat/docker_nonroot.git
    ```

2. Build the Docker image and start the container by running the following command.

    ``` shell
    make dcup 
    ```

3. Run the following command to test the endpoint at privileged port 80.

    ``` shell
    curl http://localhost:80/health
    ```
