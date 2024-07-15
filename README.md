# 유저 관리 시스템

고등학교 동아리 프로젝트.

## 개요

- Language: Golang (1.21.5)
- Database: Mysql
- 개발 기간: (2024-07-12 ~ 2024-07-15)
- 개발 환경: macOS Monteray (12.7.2)


## 기능

- 회원가입, 회원탈퇴, 로그인, 로그아웃의 기본 기능 구현
- JWT 사용
- 기능 3

### 요구 사항
Required package:
- [Godotenv](github.com/joho/godotenv)
- [Logrus](github.com/sirupsen/logrus)
- [Fiber v2](github.com/gofiber/fiber/v2)
- [Jwt-go](github.com/dgrijalva/jwt-go)
- [Crypto](golang.org/x/crypto)
- [Gorm mysql driver](gorm.io/driver/mysql)
- [Gorm](gorm.io/gorm)

### 설치

1. 다운로드

   ```bash
   git clone https://github.com/caaasperr/UserAPI.git
   cd UserAPI

2. 모듈 다운로드
   
   ```bash 
   go mod init UserAPI
   go mod tidy
   go get 
  
3. 실행
   
    ```bash
    go run main.go
  
