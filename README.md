# 🚀 Prodexa

**Prodexa** is a scalable, multi-tenant SaaS platform for **Product & Inventory Management**, built to help businesses efficiently manage stock, orders, and operations from a single system.

---

## 🧠 Overview

Prodexa is designed for:

* Retail businesses
* Warehouses
* D2C brands
* Distributors

It provides a centralized platform to:

* Track inventory in real-time
* Manage products and orders
* Monitor business performance

---

## ✨ Key Features

### 🔐 Multi-Tenant Architecture

* Secure data isolation using `tenant_id`
* Supports multiple businesses on a single platform

### 📦 Product Management

* Create, update, and manage products
* SKU-based tracking
* Category support

### 📊 Inventory Management

* Real-time stock tracking
* Stock in/out logs
* Low stock alerts

### 🛒 Order Management

* Sales and purchase orders
* Order lifecycle tracking

### 👥 Customer & Vendor Management

* Maintain customer database
* Vendor tracking for procurement

### 📈 Dashboard & Analytics

* Sales insights
* Inventory status
* Business metrics

---

## 🏗️ Tech Stack

### Backend

* Go (Golang)
* Gin / Fiber
* REST APIs

### Frontend

* Next.js (React)
* TailwindCSS

### Database

* MySQL (multi-tenant design)

### Infrastructure

* AWS (EC2, RDS, S3)
* Docker

---

## 🧩 Architecture

* Multi-tenant system using shared database
* `tenant_id` used for logical data isolation
* Middleware-based tenant enforcement
* Scalable service-oriented structure

---

## 🔄 Core Workflows

### Inventory Flow

1. Purchase order created
2. Stock updated
3. Inventory logs recorded

### Sales Flow

1. Sales order created
2. Stock deducted
3. Revenue tracked

---

## 🔐 Security

* JWT-based authentication
* Role-based access control (RBAC)
* Tenant-level data isolation
* Input validation & secure APIs

---

## 📦 Project Structure

```
prodexa/
│── cmd/                # Entry point
│── internal/
│   ├── handlers/      # API handlers
│   ├── services/      # Business logic
│   ├── repositories/  # DB operations
│   ├── models/        # Data models
│   ├── middleware/    # Auth & tenant middleware
│── pkg/               # Shared utilities
│── configs/           # Config files
│── migrations/        # DB migrations
│── docker/            # Docker setup
```

---

## ⚙️ Getting Started

### Prerequisites

* Go 1.21+
* MySQL
* Docker (optional)

### Setup

```bash
git clone https://github.com/your-username/prodexa.git
cd prodexa
```

### Environment Variables

Create `.env` file:

```
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=password
DB_NAME=prodexa
JWT_SECRET=your_secret
```

### Run Application

```bash
go run cmd/main.go
```

---

## 🧪 Testing

```bash
go test ./...
```

---

## 🚀 Roadmap

* [ ] Multi-warehouse support
* [ ] Barcode scanning
* [ ] Invoice generation (PDF)
* [ ] GST integration
* [ ] Mobile application
* [ ] AI-based demand forecasting

---

## 💰 SaaS Vision

Prodexa aims to become a **complete business operating system** for SMEs:

* Inventory
* Orders
* Analytics
* Automation

---

## 🤝 Contribution

Contributions are welcome!

1. Fork the repo
2. Create a new branch
3. Submit a PR

---

## 📄 License

MIT License

---

## 👨‍💻 Author

Built with ❤️ by Keshav Kumar
