
# 📋 Go-Management

**Go-Management** adalah aplikasi manajemen berbasis Golang yang mendukung fungsi autentikasi dan pengelolaan supplier, kategori, dan item. Proyek ini dikembangkan menggunakan berbagai komponen seperti controller, service, repository, dan domain.

---

## ⚙️ Persiapan Awal

**Prasyarat:**
1. 🟡 **Go** - Pastikan Golang telah terpasang.
2. 🐳 **Docker** - Untuk menjalankan database dengan Docker.
3. 📫 **Postman/HTTP Client** - Untuk menguji endpoint (opsional).

---

## 🛠️ Konfigurasi Awal, `default run int port :3000`

1. **Clone repository:**
   ```bash
   git clone https://github.com/koriebruh/go-management.git
   cd go-management
   ```

2. **Jalankan Database**  
   Jalankan perintah berikut untuk menjalankan API melalui Docker:
   ```bash
   docker-compose up -d
   ```

3. **Sesuaikan konfigurasi** di file `.env` sesuai kebutuhan Anda, pastikan anda meyesuaikan `docker-compose.yaml`.

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

## 📌 Daftar Endpoint Public 

### 1. 📝 Register
Endpoint ini digunakan untuk melakukan registrasi user baru.

- **Endpoint**: `api/auth/register`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "username": "fatleme",
     "password": "fatlem123",
     "email": "fatlem@gmail.com"
   }
   ```
- **Deskripsi**: Melakukan registrasi user baru. Pengguna dapat mendaftar dengan menyediakan username, password, dan email untuk membuat akun baru.
  
### 2. 🔑 Login
Endpoint ini digunakan untuk melakukan login.

- **Endpoint**: `api/auth/login`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "username": "fatlem",
     "password": "fatlem123"
   }
   ```
- **Deskripsi**: Login ke aplikasi menggunakan username dan password yang sudah terdaftar. Mengembalikan token yang dapat digunakan untuk akses lebih lanjut.
  
### 3. 🚪 Logout
Endpoint ini digunakan untuk logout user.

- **Endpoint**: `api/auth/logout`
- **Metode**: `POST`
  - **Response**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": {
       "message": "LogOut Success"
      }
   }
   ```
- **Deskripsi**: Melakukan logout dari aplikasi. Menghapus session atau token yang digunakan oleh pengguna saat ini.

### 4. 🧑‍💼 Daftar Admin
Endpoint ini digunakan untuk menampilkan daftar admin yang terdaftar.

- **Endpoint**: `api/admins`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "username": "fatlem",
         "email": "fatlem@gmail.com",
         "createdAt": "2024-11-10T14:11:41.466Z"
       }
     ]
   }   
   ```
- **Deskripsi**: Menampilkan daftar admin yang terdaftar dalam sistem. Menampilkan informasi admin yang memiliki hak akses ke aplikasi.

## 📌 Daftar Endpoint Terpoteksi
  
### 1. ➕ Menambah Supplier
Endpoint ini digunakan untuk menambah data supplier baru.

- **Endpoint**: `api/suppliers`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "fatlem",
     "contact": "fatlem@gmail.com"
   }
   ```
- **Deskripsi**: Menambah data supplier baru ke dalam sistem. Pengguna dapat memberikan nama dan kontak supplier.
  
### 2. 🏷️ Menambah Kategori
Endpoint ini digunakan untuk menambah kategori baru.

- **Endpoint**: `api/categories`
- **Metode**: `POST`
- **Request**:
   ```json
   {
     "name": "Elektronik",
     "description": "Laptop Gaming"
   }
   ```
- **Deskripsi**: Menambah kategori baru untuk item. Pengguna dapat menambahkan nama kategori dan deskripsi terkait.
  
### 3. 📦 Menambah Item
Endpoint ini digunakan untuk menambah item baru.

- **Endpoint**: `api/items`
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
- **Deskripsi**: Menambah item baru dengan informasi seperti nama, deskripsi, harga, kuantitas, kategori, dan supplier.
  
### 4. 🏷️ Daftar Kategori
Endpoint ini digunakan untuk menampilkan daftar kategori.

- **Endpoint**: `api/categories`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "id": 1,
         "name": "Elektronik"
       },
       {
         "id": 2,
         "name": "Makanan"
       },
       {
         "id": 3,
         "name": "Kecantikan"
       }
     ]
   }
   ```
- **Deskripsi**: Menampilkan daftar kategori yang sudah ada. Menampilkan ID dan nama kategori yang terdaftar dalam sistem.
  
### 5. 🛒 Daftar Supplier
Endpoint ini digunakan untuk menampilkan daftar supplier.

- **Endpoint**: `api/suppliers`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "name": "Windows",
         "contact_info": "2813412"
       }
     ]
   }
   ```  
- **Deskripsi**: Menampilkan daftar supplier yang sudah terdaftar dalam sistem. Menampilkan informasi supplier yang ada.

### 6. 🔍 Daftar semua item
Endpoint ini digunakan untuk mencari semua item 

- **Endpoint**: `api/items/condition?condition=under&threshold=4`
- **Metode**: `GET`
  - **Resonse**:
     ```json
     {
      "code": 200,
      "status": "OK",
      "data": [
        {
          "id": 1,
          "name": "Iphone 12x",
          "description": "Limited edition for jamal Only",
          "price": 32000000,
          "quantity": 3,
          "category": "Comsetic",
          "supplier": "Windows",
          "created_by": "jamal",
          "created_at": "2024-11-10T14:54:46.585Z",
          "updated_at": "2024-11-10T14:54:46.585Z"
        },
        {
          "id": 2,
          "name": "RC48",
          "description": "Limited edition for jamal Only",
          "price": 32000000,
          "quantity": 3,
          "category": "Electronic",
          "supplier": "Windows",
          "created_by": "jamal",
          "created_at": "2024-11-10T14:55:09.403Z",
          "updated_at": "2024-11-10T14:55:09.403Z"
        }
      ]
    }
     ```
- **Deskripsi**: Mencari semua jenis item yang tersedia.
  
### 7. 🔍 Cari Item Berdasarkan Kondisi 
Endpoint ini digunakan untuk mencari item berdasarkan kondisi tertentu.

- **Endpoint**: `api/items/condition?condition=under&threshold=4`
- **Metode**: `GET`
- **Resonse**:
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
- **Deskripsi**: Mencari item berdasarkan kondisi tertentu, seperti harga atau kuantitas. Membantu menemukan item yang sesuai dengan kriteria yang diinginkan.


### 8. 📈 Laporan Item Berdasarkan Kategori 
Endpoint ini digunakan untuk menampilkan laporan item berdasarkan kategori tertentu.

- **Endpoint**: `api/items/category?category=Electronic`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": {
       "ByCategory": "Electronic",
       "Description": "electronic",
       "TotalItem": 1,
       "TotalQuantity": 3,
       "TotalValue": 96000000,
       "Items": [
         {
           "Name": "RC48",
           "Quantity": 3,
           "SupplierID": 1
         }
       ]
     }
   }
   ```
- **Deskripsi**: Menampilkan laporan mengenai item berdasarkan kategori tertentu. Memudahkan untuk melihat semua item yang ada dalam kategori tersebut.

### 9. 💾 Metrik Inventaris Item
Endpoint ini digunakan untuk menampilkan metrik inventaris untuk item.

- **Endpoint**: `api/items/metric`
- **Metode**: `GET`
- **Resonse**:
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
- **Deskripsi**: Menampilkan metrik terkait inventaris item. Menyediakan informasi tentang stok, nilai rata-rata item, distribusi stok, dan lainnya.

### 10. 📊 Laporan items
Endpoint ini digunakan untuk menampilkan laporan item 

- **Endpoint**: `/api/items/info`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": {
       "total_items": 2,
       "total_stock_value": 192000000,
       "average_item_price": 32000000,
       "total_categories": 2,
       "total_supplier": 1,
       "updated_at": "2024-11-11T13:42:26.601338354Z"
     }
   }
   ```
- **Deskripsi**: Memberikan laporan item yang terperinci berdasarkan total item, total value dll.

### 11. 📊 Laporan Categories
Endpoint ini digunakan untuk menampilkan laporan categories

- **Endpoint**: `api/categories/info`
- **Metode**: `GET`
- **Resonse**:
   ```json
  {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "category_id": 1,
         "category_name": "Comsetic",
         "item_count": 1,
         "total_stock_value": 96000000,
         "average_item_price": 32000000
       },
       {
         "category_id": 2,
         "category_name": "Electronic",
         "item_count": 1,
         "total_stock_value": 96000000,
         "average_item_price": 32000000
       },
       {
         "category_id": 3,
         "category_name": "Glass",
         "item_count": 0,
         "total_stock_value": 0,
         "average_item_price": 0
       }
     ]
   }
   ```
- **Deskripsi**: Memberikan laporan dari setiap category beserta detail nya.

### 11. 📊 Laporan supplier
Endpoint ini digunakan untuk menampilkan laporan supplier 

- **Endpoint**: `/api/supplier/info`
- **Metode**: `GET`
- **Resonse**:
   ```json
   {
     "code": 200,
     "status": "OK",
     "data": [
       {
         "supplier_name": "Windows",
         "total_items": 2,
         "total_value": 192000000
       },
       {
         "supplier_name": "PT PELITA HARAPAN",
         "total_items": 1,
         "total_value": 44000000
       }
     ]
   }
   ```
- **Deskripsi**: Memberikan laporan dari masing masing supplier.

--- 

## 🧪 Menjalankan Uji Coba Endpoint

Gunakan `AUTH_TEST.http`, `CATEGORIES_TEST.http`, `ITEM_TEST.http`,`SUPPLIER_TEST.http`, yg tersedia didalam source code atau alat HTTP client seperti Postman untuk menguji endpoint secara lokal.
