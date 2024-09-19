# Nongki App

## Struktur Projek

```bash
├── cmd
│   └── server
│       └── main.go           
├── config
│   └── app.go
│   └── sqlx.go
│   └── redis.go              
├── internal    
│   └── domain
│       └── base_domain.go 
│       └── user.go 
│   └── handler
│       └── auth_handler.go
│       └── user_handler.go  
│   └── repository
│       └── user_repository.go 
│   └── request
│       └── login_request.go
│       └── register_request.go
│       └── update_user_request.go 
│       └── refresh_token_request.go   
│   └── response
│       └── base_response.go
│       └── login_response.go  
│       └── register_response.go 
│       └── user_response.go 
│   └── router
│       └── auth_route.go 
│       └── user_route.go 
│       └── router.go 
│       └── router_setting.go
│   └── usecase     
│       └── auth_usecase.go 
│       └── user_usecase.go
├── pkg
│   └── constant
│       └── response_message.go 
│   └── db
│       └── sqlx.go 
│   └── helpers
│       └── formatter.go 
│   └── jwt
│       └── jwt.go 
│   └── log
│       └── logrus.go 
│   └── middleware  
│       └── middleware.go              
├── migrations
│   └── <timestamp>_create_users_table.sql  
├── Dockerfile                
├── docker-compose.yml        
├── Makefile                  
└── .env                      
```

## Cara Menjalankan Aplikasi

1. **Kloning Repository**
```bash
git clone <repository-url>
cd <repository-directory>
```

2. **Jalankan Menggunakan Makefile**
* Migrasi Database
```bash
make migrate-up
```

* Menjalankan Aplikasi
```bash
make run
```

## ERD
```bash
+-------------------+
|      users        |
+-------------------+
| id (PK)           |
| name              |
| email (UNIQUE)    |
| address           |
| gender            |
| marital_status    |
| created_at        |
| updated_at        |
| deleted_at        |
| deleted_by        |
+-------------------+
```

## Arsitektur Aplikasi

1. **Presentation Layer (Delivery)**

* Responsibility: Menangani input dari pengguna dan memberikan output.
* Components: Handlers dan Middleware.
2. **Use Case Layer (Business Logic)**

* Responsibility: Mengimplementasikan logika bisnis aplikasi.
* Components: Use Cases untuk berbagai fungsi aplikasi.
3. **Domain Layer (Entities)**

* Responsibility: Menyimpan model data dan aturan bisnis dasar.
* Components: Entities seperti User.
4. **Repository Layer (Persistence)**

* Responsibility: Mengelola akses data dan operasi CRUD.
* Components: Repositories untuk query database.
5. **Infrastructure Layer**

* Responsibility: Menyediakan implementasi teknis dan layanan eksternal.
* Components: Database Configuration dan Migrations.
6. **Configuration**

* Responsibility: Menyimpan pengaturan aplikasi dan environment variables.
* Components: Config Files seperti .env

## Diagram Arsitektur:
```bash
+------------------------------------+
|           Presentation Layer       |
|                                    |
| +--------------+     +------------+ |
| |  Handlers    |<--->| Middleware | |
| +--------------+     +------------+ |
+------------------------------------+
              |
              |
              v
+------------------------------------+
|           Use Case Layer           |
|                                    |
| +-----------------------------+      |
| |         Use Cases           |      |
| +-----------------------------+      |
+------------------------------------+
              |
              |
              v
+------------------------------------+
|           Domain Layer             |
|                                    |
| +-----------------------------+      |
| |          Entities           |      |
| +-----------------------------+      |
| |        Value Objects        |      |
| +-----------------------------+      |
+------------------------------------+
              |
              |
              v
+------------------------------------+
|           Repository Layer         |
|                                    |
| +-----------------------------+      |
| |         Repositories        |      |
| +-----------------------------+      |
+------------------------------------+
              |
              |
              v
+------------------------------------+
|           Infrastructure Layer     |
|                                    |
| +-----------------------------+      |
| |  Database Configuration     |      |
| +-----------------------------+      |
| |         Migrations          |      |
| +-----------------------------+      |
+------------------------------------+
              |
              |
              v
+------------------------------------+
|           Configuration            |
|                                    |
| +-----------------------------+      |
| |       Config Files          |      |
| +-----------------------------+      |
+------------------------------------+
```

## Postman Collection
[API Nongki Collection](https://bold-crater-47260.postman.co/workspace/New-Team-Workspace~edbc1147-9fc2-4b06-a265-665a134b7a5d/collection/29819488-27da1788-c107-4929-be41-189d28b9ac7d?action=share&creator=29819488)