
# 📋 Go-Management

**Go-Management** adalah aplikasi manajemen berbasis Golang yang mendukung fungsi autentikasi dan pengelolaan supplier, kategori, dan item. Proyek ini dikembangkan menggunakan berbagai komponen seperti controller, service, repository, dan domain.

---

## ⚙️ Persiapan Awal

**Prasyarat:**
1. 🟡 **Go** - Pastikan Golang telah terpasang.
2. 🐳 **Docker** - Untuk menjalankan database dengan Docker.
3. 📫 **Postman/HTTP Client** - Untuk menguji endpoint (opsional).

---

## 🛠️ Konfigurasi Awal

1. **Clone repository:**
   ```bash
   git clone <repository-url>
   cd go-management
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Sesuaikan konfigurasi** di file `.env` sesuai kebutuhan Anda.

---

## 🚀 Menjalankan Aplikasi

1. **Jalankan Database**  
   Jalankan perintah berikut untuk menjalankan database melalui Docker:
   ```bash
   docker-compose up -d
   ```

2. **Menjalankan Aplikasi Go**  
   Untuk menjalankan aplikasi Go:
   ```bash
   go run main.go
   ```

---

## 🗂️ Struktur Proyek

- **`cnf`**: ⚙️ Konfigurasi aplikasi.
- **`controller`**: 📬 Controller untuk menangani request.
- **`docs`**: 📝 Dokumentasi API.
- **`domain`**: 📦 Struktur domain aplikasi.
- **`dto`**: 📤 Data Transfer Object, struktur data yang diterima atau dikirim.
- **`repository`**: 🗄️ Logika akses data ke database.
- **`service`**: 📊 Logika bisnis aplikasi.
- **`utils`**: 🧰 Kode tambahan atau helper.

---

## 📌 Daftar Endpoint

### Public Endpoints

1. **Hello**
   - **Endpoint**: `/`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan pesan sambutan.

2. **Register**
   - **Endpoint**: `/api/auth/register`
   - **Metode**: `POST`
   - **Deskripsi**: Mendaftarkan pengguna baru.

3. **Login**
   - **Endpoint**: `/api/auth/login`
   - **Metode**: `POST`
   - **Deskripsi**: Mengotentikasi pengguna.

4. **Logout**
   - **Endpoint**: `/api/auth/logout`
   - **Metode**: `POST`
   - **Deskripsi**: Keluar dari sesi pengguna.

---

### Protected Endpoints (Memerlukan Autentikasi)

#### Admin

1. **Daftar Admin**
   - **Endpoint**: `/api/admins`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan daftar admin.

2. **Hello (Autentikasi)**
   - **Endpoint**: `/hi`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan pesan sambutan untuk pengguna terotentikasi.

#### Kategori

3. **Daftar Kategori**
   - **Endpoint**: `/api/categories`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan daftar semua kategori.

4. **Buat Kategori**
   - **Endpoint**: `/api/categories`
   - **Metode**: `POST`
   - **Deskripsi**: Menambahkan kategori baru.

5. **Ringkasan Kategori**
   - **Endpoint**: `/api/categories/info`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan ringkasan kategori.

#### Item

6. **Daftar Item**
   - **Endpoint**: `/api/items`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan daftar semua item.

7. **Buat Item**
   - **Endpoint**: `/api/items`
   - **Metode**: `POST`
   - **Deskripsi**: Menambahkan item baru.

8. **Ringkasan Item**
   - **Endpoint**: `/api/items/info`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan ringkasan item.

9. **Cari Item Berdasarkan Kondisi**
   - **Endpoint**: `/api/items/condition`
   - **Metode**: `GET`
   - **Deskripsi**: Menampilkan item berdasarkan kondisi tertentu.

10. **Metrik Inventaris Item**
    - **Endpoint**: `/api/items/metric`
    - **Metode**: `GET`
    - **Deskripsi**: Menampilkan metrik inventaris untuk item.

11. **Laporan Item Berdasarkan Kategori**
    - **Endpoint**: `/api/items/category`
    - **Metode**: `GET`
    - **Deskripsi**: Menampilkan laporan item berdasarkan kategori.

#### Supplier

12. **Buat Supplier**
    - **Endpoint**: `/api/suppliers`
    - **Metode**: `POST`
    - **Deskripsi**: Menambahkan supplier baru.

13. **Daftar Supplier**
    - **Endpoint**: `/api/suppliers`
    - **Metode**: `GET`
    - **Deskripsi**: Menampilkan daftar semua supplier.

14. **Ringkasan Supplier**
    - **Endpoint**: `/api/suppliers/info`
    - **Metode**: `GET`
    - **Deskripsi**: Menampilkan ringkasan supplier.

---

## 🔌 API Endpoint

Berikut adalah beberapa endpoint utama untuk melakukan operasi register, login, logout, menambah supplier, kategori, dan item.

### 1. 📝 Register
Endpoint ini digunakan untuk melakukan registrasi user baru.

- **Endpoint**: `/auth/register`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "username": "user1",
     "password": "password123",
     "email": "user1@gmail.com"
   }
   ```
- **Kode Implementasi**:
   ```go
   func Register(w http.ResponseWriter, r *http.Request) {
       // Logic untuk register user
   }
   ```

### 2. 🔑 Login
Endpoint ini digunakan untuk melakukan login.

- **Endpoint**: `/auth/login`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "username": "user1",
     "password": "password123"
   }
   ```
- **Kode Implementasi**:
   ```go
   func Login(w http.ResponseWriter, r *http.Request) {
       // Logic untuk login user
   }
   ```

### 3. 🚪 Logout
Endpoint ini digunakan untuk logout user.

- **Endpoint**: `/auth/logout`
- **Metode**: `POST`
- **Kode Implementasi**:
   ```go
   func Logout(w http.ResponseWriter, r *http.Request) {
       // Logic untuk logout user
   }
   ```

### 4. ➕ Menambah Supplier
Endpoint ini digunakan untuk menambah data supplier baru.

- **Endpoint**: `/suppliers`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "Supplier A",
     "contact": "supplierA@example.com"
   }
   ```
- **Kode Implementasi**:
   ```go
   func AddSupplier(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menambah supplier baru
   }
   ```

### 5. 🏷️ Menambah Kategori
Endpoint ini digunakan untuk menambah kategori baru.

- **Endpoint**: `/categories`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "Elektronik",
     "description": "Laptop Gaming"
   }
   ```
- **Kode Implementasi**:
   ```go
   func AddCategory(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menambah kategori baru
   }
   ```

### 6. 📦 Menambah Item
Endpoint ini digunakan untuk menambah item baru.

- **Endpoint**: `/items`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "Asus Rog G15 G513RM",
     "description": "Republic of Gamers",
     "price": 22000000,
     "quantity": 2,
     "category_id": 1,
     "supplier_id": 1
   }
   ```
- **Kode Implementasi**:
   ```go
   func AddItem(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menambah item baru
   }
   ```

### 7. 🏷️ Daftar Kategori
Endpoint ini digunakan untuk menampilkan daftar kategori.

- **Endpoint**: `/categories`
- **Metode**: `GET`
- **Kode Implementasi**:
   ```go
   func GetCategories(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan kategori
   }
   ```

### 8. 🛒 Daftar Supplier
Endpoint ini digunakan untuk menampilkan daftar supplier.

- **Endpoint**: `/suppliers`
- **Metode**: `GET`
- **Kode Implementasi**:
   ```go
   func GetSuppliers(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan supplier
   }
   ```

### 9. 📊 Menambah Kategori Baru
Endpoint ini digunakan untuk menambah kategori baru.

- **Endpoint**: `/categories`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "Furniture",
     "description": "Modern office chairs"
   }
   ```
- **Kode Implementasi**:
   ```go
   func AddCategory(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menambah kategori baru
   }
   ```

### 10. 🔍 Cari Item Berdasarkan Kondisi
Endpoint ini digunakan untuk mencari item berdasarkan kondisi tertentu.

- **Endpoint**: `/items/condition`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan item berdasarkan kondisi tertentu seperti harga, kuantitas, dll.
- **Kode Implementasi**:
   ```go
   func SearchItemsByCondition(w http.ResponseWriter, r *http.Request) {
       // Logic untuk mencari item berdasarkan kondisi
   }
   ```

### 11. 📈 Laporan Item Berdasarkan Kategori
Endpoint ini digunakan untuk menampilkan laporan item berdasarkan kategori tertentu.

- **Endpoint**: `/items/category`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan laporan item berdasarkan kategori.
- **Kode Implementasi**:
   ```go
   func GetItemReportByCategory(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan laporan item berdasarkan kategori
   }
   ```

### 12. 💾 Metrik Inventaris Item
Endpoint ini digunakan untuk menampilkan metrik inventaris untuk item.

- **Endpoint**: `/items/metric`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan metrik inventaris item seperti stok, penjualan, dll.
- **Kode Implementasi**:
   ```go
   func GetItemInventoryMetrics(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan metrik inventaris item
   }
   ```

### 13. 🧑‍💼 Daftar Admin
Endpoint ini digunakan untuk menampilkan daftar admin yang terdaftar.

- **Endpoint**: `/admins`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan daftar admin.
- **Kode Implementasi**:
   ```go
   func GetAdmins(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan daftar admin
   }
   ```

### 14. 🏅 Hello (Autentikasi)
Endpoint ini digunakan untuk menampilkan pesan sambutan untuk pengguna yang sudah terautentikasi.

- **Endpoint**: `/hi`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan pesan sambutan untuk pengguna yang sudah login.
- **Kode Implementasi**:
   ```go
   func HelloAuthenticated(w http.ResponseWriter, r *http.Request) {
       // Logic untuk menampilkan pesan sambutan bagi pengguna terautentikasi
   }
   ```

--- 

## 🧪 Menjalankan Uji Coba Endpoint

Gunakan `AUTH_TEST.http`, `CATEGORIES_TEST.http`, `ITEM_TEST.http`, atau alat HTTP client seperti Postman untuk menguji endpoint secara lokal.

---
