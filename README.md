# Gocourse_Web

Curso de Desarrollo REST y Microservicios en Go 

Para construir la base de datos con docker se utiliza docker compose y este archivo yml. 
```yml
version: "3.5"
services:
  go-course-web:
    platform: linux/amd64
    container_name: go-course-web
    build:
      context: "Path to DockerFile"
      dockerfile: Dockerfile
    environment:
      MYSQL_ROOT_PASSWORD: "<YOUR-PASSWORD>"
      MYSQL_DATABASE: go_course_web
    ports:
      - "3306:3306"
    volumes:
      - "paht:/docker-entrypoint-initdb.d/init.sql"
```
