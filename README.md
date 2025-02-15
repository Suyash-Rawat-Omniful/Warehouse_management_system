# Warehouse Management System

A Warehouse Management System (WMS) built using **PostgreSQL** and **Go** to efficiently manage SKUs, hubs, inventory, and tenants. **Redis (via Docker)** is used for caching frequently accessed SKU and Hub data.

## 🚀 Features

- **SKU Management**: Store and track product SKUs.
- **Hub Management**: Maintain warehouse hubs.
- **Inventory Management**: Keep records of stock levels.
- **Tenant Management**: Handle multiple tenants for multi-warehouse operations.
- **Caching with Redis**: Frequently accessed SKU and Hub data is cached to improve performance.

## 🛠️ Tech Stack

- **Database**: PostgreSQL
- **Backend**: Golang
- **Caching**: Redis (via Docker)

## 📂 Database Schema

### 1️⃣ SKU Table (`sku`)
| Column     | Type           | Description                      |
|-----------|---------------|----------------------------------|
| id        | INT (PK)      | Primary Key                      |
| product_id| INT           | ID of the associated product     |
| name      | VARCHAR(255)  | Name of SKU                      |
| price     | INT           | Price of the SKU                 |
| fragile   | BOOLEAN       | Indicates if the item is fragile |
| image_url | TEXT          | URL of the SKU image             |
| created_at| TIMESTAMP     | Timestamp when record was created |
| updated_at| TIMESTAMP     | Timestamp when record was updated |
| deleted_at| TIMESTAMP (Nullable) | Soft delete field |

### 2️⃣ Hub Table (`hub`)
| Column     | Type        | Description               |
|-----------|------------|---------------------------|
| id        | INT (PK)   | Primary Key               |
| tenant_id | INT        | Foreign Key (References Tenant) |
| created_at| TIMESTAMP  | Timestamp when record was created |
| updated_at| TIMESTAMP  | Timestamp when record was updated |

### 3️⃣ Inventory Table (`inventory`)
| Column     | Type        | Description                     |
|-----------|------------|---------------------------------|
| id        | INT (PK)   | Primary Key                     |
| hub_id    | INT        | Foreign Key (References Hub)    |
| sku_id    | INT        | Foreign Key (References SKU)    |
| quantity  | INT        | Available stock quantity        |
| created_at| TIMESTAMP  | Timestamp when record was created |
| updated_at| TIMESTAMP  | Timestamp when record was updated |

### 4️⃣ Tenant Table (`tenant`)
| Column     | Type           | Description               |
|-----------|---------------|---------------------------|
| id        | INT (PK)      | Primary Key               |
| name      | VARCHAR(255)  | Name of the tenant        |
| email     | VARCHAR(255)  | Unique email for tenant   |
| created_at| TIMESTAMP     | Timestamp when record was created |
| updated_at| TIMESTAMP     | Timestamp when record was updated |

## 🛠️ Setup Instructions

### 1️⃣ Clone the Repository
```sh
git clone https://github.com/Suyash-Rawat-Omniful/Warehouse_management_system.git
cd Warehouse_Management_System
