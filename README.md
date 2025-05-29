
# Ecommerce-Golang

A simple E-Commerce backend written in Go, using a layered architecture with raw SQL (no ORM). This project includes user authentication, product management, cart functionality, order processing, and JWT-based role access.

## 🔧 Features

- User registration and login (JWT auth)
- Role-based access control (admin/public)
- Product CRUD (admin only)
- Add to cart, view cart, remove/change quantity
- Place order, cancel order, view order history
- Pagination for products, cart, and orders
- Raw SQL with `database/sql`
- Clean layered architecture:
  - routes → controllers → services → repositories → database

## 🧱 Tech Stack

- Go (Golang)
- PostgreSQL
- JWT for authentication
- HTML/CSS/JS frontend (optional)
- No ORM (pure SQL)

## 🗂️ Folder Structure

```
├── controllers/
├── services/
├── repositories/
├── models/
├── database/
├── middleware/
├── routes/
├── main.go
└── go.mod
```

## 🚀 Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/Akshat-AFB/Ecommerce-Golang.git
cd Ecommerce-Golang
```

### 2. Install Dependecies

```bash
go mod tidy
```
### 2. Configure Environment

Update database credentials in `database/connection.go` or use environment variables.

### 3. Initialize Database

Ensure PostgreSQL is running and create necessary tables. You can use the SQL scripts provided (or manually create):

```sql
-- Example
CREATE TABLE users (...);
CREATE TABLE products (...);
CREATE TABLE cart_items (...);
CREATE TABLE orders (...);
CREATE TABLE order_items (...);
```

### 4. Run the Application

```bash
go run main.go
```

### 5. Test API

Use Postman or curl to interact with endpoints.

## 🔐 Auth Endpoints

| Method | Endpoint          | Description       |
|--------|-------------------|-------------------|
| POST   | /api/register     | Register a user   |
| POST   | /api/login        | Login & get token |

## 📦 Product Endpoints

| Method | Endpoint            | Description                |
|--------|---------------------|----------------------------|
| GET    | /api/products       | List all products          |
| POST   | /api/products/create | Add product (admin)       |
| PUT    | /api/products/:id/update | Update product (admin) |
| DELETE | /api/products/:id/delete | Delete product (admin) |

## 🛒 Cart Endpoints

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| GET    | /api/cart        | View cart           |
| POST   | /api/cart/add    | Add to cart         |
| DELETE | /api/cart/remove/:id | Remove item     |
| POST   | /api/cart/change/:product_id | Change qty |

## 📦 Order Endpoints

| Method | Endpoint              | Description          |
|--------|-----------------------|----------------------|
| POST   | /api/orders/place     | Place order          |
| GET    | /api/orders           | View user orders     |
| POST   | /api/orders/cancel/:id | Cancel order        |

<!-- ## ✅ Todo

- Add unit tests
- Dockerize the app
- Add Swagger docs -->

## 📄 License

MIT

---

Made with ❤️ by [Akshat-AFB](https://github.com/Akshat-AFB)
