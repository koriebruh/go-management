# Format JSON Output Sistem Inventaris

```json
{
  "system_summary": {
    "total_items": 150,
    "total_stock_value": 75000000,
    "average_item_price": 500000,
    "total_categories": 8,
    "total_suppliers": 5,
    "last_updated": "2024-11-07T10:00:00Z"
  },
  "low_stock_alert": {
    "threshold": 5,
    "items": [
      {
        "id": 1,
        "name": "Laptop Asus",
        "category": "Electronics",
        "quantity": 3,
        "price": 12000000,
        "supplier": "PT Supplier Electronics"
      },
      {
        "id": 2,
        "name": "Monitor LG",
        "category": "Electronics",
        "quantity": 4,
        "price": 2500000,
        "supplier": "PT Supplier Electronics"
      }
    ]
  },
  "category_summary": {
    "categories": [
      {
        "id": 1,
        "name": "Electronics",
        "stats": {
          "total_items": 45,
          "total_value": 25000000,
          "average_price": 2500000
        },
        "items": [
          {
            "id": 1,
            "name": "Laptop Asus",
            "quantity": 3,
            "price": 12000000,
            "stock_value": 36000000
          },
          {
            "id": 2,
            "name": "Monitor LG",
            "quantity": 4,
            "price": 2500000,
            "stock_value": 10000000
          }
        ]
      }
    ]
  },
  "supplier_summary": {
    "suppliers": [
      {
        "id": 1,
        "name": "PT Supplier Electronics",
        "stats": {
          "total_items_supplied": 50,
          "total_supply_value": 30000000
        },
        "supplied_items": [
          {
            "id": 1,
            "name": "Laptop Asus",
            "category": "Electronics",
            "quantity": 3,
            "price": 12000000
          },
          {
            "id": 2,
            "name": "Monitor LG",
            "category": "Electronics",
            "quantity": 4,
            "price": 2500000
          }
        ]
      }
    ]
  },
  "inventory_metrics": {
    "stock_status": {
      "healthy_stock": 120,
      "low_stock": 25,
      "out_of_stock": 5
    },
    "value_metrics": {
      "highest_value_category": "Electronics",
      "lowest_value_category": "Stationery",
      "average_item_value": 500000
    },
    "stock_distribution": {
      "by_category": {
        "Electronics": "30%",
        "Furniture": "25%",
        "Office Supplies": "45%"
      }
    }
  }
}
```

## Penjelasan Struktur

### 1. System Summary
- Menampilkan ringkasan keseluruhan sistem
- Mencakup total barang, nilai, dan statistik umum

### 2. Low Stock Alert
- Daftar barang dengan stok di bawah threshold
- Informasi detail setiap barang termasuk supplier

### 3. Category Summary
- Ringkasan per kategori
- Statistik kategori dan daftar barang

### 4. Supplier Summary
- Informasi pemasok dan barang yang disuplai
- Statistik nilai dan jumlah barang per pemasok

### 5. Inventory Metrics
- Metrik kesehatan stok
- Distribusi nilai dan stok
- Statistik perbandingan antar kategori