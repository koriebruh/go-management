
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

2. **Jalankan Database**  
   Jalankan perintah berikut untuk menjalankan database melalui Docker:
   ```bash
   docker-compose up -d
   ```

3. **Sesuaikan konfigurasi** di file `.env` sesuai kebutuhan Anda.

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

### 3. 🚪 Logout
Endpoint ini digunakan untuk logout user.

- **Endpoint**: `/auth/logout`
- **Metode**: `POST`
  - **Request**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": {
       "message": "LogOut Success"
      }
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

### 7. 🏷️ Daftar Kategori
Endpoint ini digunakan untuk menampilkan daftar kategori.

- **Endpoint**: `/categories`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 201,
     "status": "Created",
     "data": {
       "message": "success created new category"
     }
   }
   ```

### 8. 🛒 Daftar Supplier
Endpoint ini digunakan untuk menampilkan daftar supplier.

- **Endpoint**: `/suppliers`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 201,
     "status": "Created",
     "data": {
       "message": "success created new supplier"
     }
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

### 10. 🔍 Cari Item Berdasarkan Kondisi
Endpoint ini digunakan untuk mencari item berdasarkan kondisi tertentu.

- **Endpoint**: `/items/condition`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "id": 6,
         "name": "Asus Rog G15 G513RM",
         "description": "Republik Of Gammer",
         "price": 22000000,
         "quantity": 2,
         "category": "",
         "supplier": "",
         "created_by": "",
         "created_at": "2024-11-10T15:18:54.041Z",
         "updated_at": "2024-11-10T15:18:54.041Z"
       }
     ]
   }
   ```
- **Deskripsi**: Menampilkan item berdasarkan kondisi tertentu seperti harga, kuantitas, dll.


### 11. 📈 Laporan Item Berdasarkan Kategori
Endpoint ini digunakan untuk menampilkan laporan item berdasarkan kategori tertentu.

- **Endpoint**: `/items/category`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 500,
     "status": "Internal Server Error",
     "data": "error transactional item"
   }
   ```
- **Deskripsi**: Menampilkan laporan item berdasarkan kategori.


### 12. 💾 Metrik Inventaris Item
Endpoint ini digunakan untuk menampilkan metrik inventaris untuk item.

- **Endpoint**: `/items/metric`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": {
       "stock_status": {
         "healthy_stock": 0,
         "low_stock": 1,
         "out_of_stock": 0
       },
       "value_metrics": {
         "highest_value_category": "Elektronik",
         "lowest_value_category": "Elektronik",
         "average_item_value": 22000000,
         "total_stock_value": 44000000,
         "total_items": 2
       },
       "stock_distribution": {
         "by_category": {
           "Elektronik": "100.00%"
         },
         "total_categories": 1,
         "total_suppliers": 1
       }
     }
   }
   ```
- **Deskripsi**: Menampilkan metrik inventaris item seperti stok, penjualan, dll.


### 13. 🧑‍💼 Daftar Admin
Endpoint ini digunakan untuk menampilkan daftar admin yang terdaftar.

- **Endpoint**: `/admins`
- **Metode**: `GET`
- **Request**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "username": "fatlem",
         "email": "fatlem@example.com",
         "createdAt": "2024-11-10T14:11:41.466Z"
       }
     ]
   }   
   ```
- **Deskripsi**: Menampilkan daftar admin.


### 14. 🏅 Hello (Autentikasi)
Endpoint ini digunakan untuk menampilkan pesan sambutan untuk pengguna yang sudah terautentikasi.

- **Endpoint**: `/hi`
- **Metode**: `GET`
- **Deskripsi**: Menampilkan pesan sambutan untuk pengguna yang sudah login.

--- 

## 🧪 Menjalankan Uji Coba Endpoint

Gunakan `AUTH_TEST.http`, `CATEGORIES_TEST.http`, `ITEM_TEST.http`, atau alat HTTP client seperti Postman untuk menguji endpoint secara lokal.

---
