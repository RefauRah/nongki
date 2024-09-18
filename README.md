# Nongki App

## Struktur Projek

```bash
├── cmd
│   └── server
│       └── main.go           # Entry point aplikasi
├── config
│   └── config.go             # File untuk pengaturan konfigurasi
├── domain
│   └── user.go               # Model User dan Usecase User
├── handler
│   └── user_handler.go       # Handler untuk HTTP request (register, login, dll.)
├── middleware
│   └── auth.go               # Middleware untuk JWT Authentication
├── migrations
│   └── <timestamp>_create_users_table.sql  # File migrasi untuk membuat tabel users
├── repository
│   └── user_repository.go    # Repository layer untuk query database
├── usecase
│   └── user_usecase.go       # Business logic untuk user
├── Dockerfile                # Dockerfile untuk membuat image aplikasi
├── docker-compose.yml        # Docker Compose file untuk menjalankan aplikasi dan database
├── Makefile                  # Makefile untuk task seperti migrasi database
└── .env                      # File environment
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
