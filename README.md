
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

### Screenshots
#### Register
<img width="1416" alt="Screenshot 2025-05-29 at 12 54 34 PM" src="https://github.com/user-attachments/assets/a61c7d87-0a13-43fd-b89c-e1f89ef67f44" />
<img width="1415" alt="Screenshot 2025-05-29 at 12 55 09 PM" src="https://github.com/user-attachments/assets/63a0aa8f-d665-4559-9cee-994770850373" />

#### Login
<img width="1418" alt="Screenshot 2025-05-29 at 12 55 43 PM" src="https://github.com/user-attachments/assets/1be0c9ab-2fe3-47c2-87ca-9b01f8fb2694" />
<img width="1422" alt="Screenshot 2025-05-29 at 12 56 42 PM" src="https://github.com/user-attachments/assets/6f16490d-cc34-4f18-918d-416560a41064" />


## 📦 Product Endpoints


| Method | Endpoint                | Description                    | Auth Required   |
|--------|-------------------------|--------------------------------|-----------------|
| GET    | `/api/products`         | List all products              | ❌ No           |
| POST   | `/api/products/create`  | Add a new product (Admin only) | ✅ Yes (Admin)  |
| PUT    | `/api/products/update`  | Update a product (Admin only)  | ✅ Yes (Admin)  |
| DELETE | `/api/products/delete`  | Delete a product (Admin only)  | ✅ Yes (Admin)  |

### Screenshots
#### Get Products (paginated)
<img width="1424" alt="Screenshot 2025-05-29 at 1 02 02 PM" src="https://github.com/user-attachments/assets/6cca2029-75f4-43d2-9bc0-a58474550b61" />

#### Add a product (with admin token)
<img width="1412" alt="Screenshot 2025-05-29 at 1 04 32 PM" src="https://github.com/user-attachments/assets/c6ed950b-2fba-4f48-915b-924702712602" />

#### Update a product (with admin token)
<img width="1419" alt="Screenshot 2025-05-29 at 1 10 24 PM" src="https://github.com/user-attachments/assets/bfc48432-419a-44e4-a7a0-42faa5fee3b0" />

#### Delete a product (with admin token)
<img width="1411" alt="Screenshot 2025-05-29 at 1 11 23 PM" src="https://github.com/user-attachments/assets/37093d9c-6702-46b2-8f9d-6042e5fe1de3" />
<img width="1416" alt="Screenshot 2025-05-29 at 1 11 45 PM" src="https://github.com/user-attachments/assets/b3ae0a63-82dd-410b-93bf-d42d7220506d" />

## 🛒 Cart Endpoints

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| GET    | /api/cart        | View cart           |
| POST   | /api/cart/add    | Add to cart         |
| DELETE | /api/cart/remove/:id | Remove item     |
| POST   | /api/cart/change/:product_id | Change qty |
### Screenshots

#### Add a product
<img width="1413" alt="Screenshot 2025-05-29 at 1 26 07 PM" src="https://github.com/user-attachments/assets/fffe8db6-b779-4786-a2e9-f3b4a2a22949" />

#### Get cart (paginated)
<img width="1416" alt="Screenshot 2025-05-29 at 1 23 25 PM" src="https://github.com/user-attachments/assets/14600ff3-3fcd-4918-8b71-5ec5e501948d" />
<img width="1417" alt="Screenshot 2025-05-29 at 3 00 20 PM" src="https://github.com/user-attachments/assets/660ceca0-1209-4aee-ad9a-b4c5ddab78a4" />
<img width="1415" alt="Screenshot 2025-05-29 at 3 15 40 PM" src="https://github.com/user-attachments/assets/edb696d4-e541-41c9-b1bc-4f480b77612f" />




## 📦 Order Endpoints

| Method | Endpoint              | Description          |
|--------|-----------------------|----------------------|
| POST   | /api/orders/place     | Place order          |
| GET    | /api/orders           | View user orders     |
| POST   | /api/orders/cancel/:id | Cancel order        |
### Screenshots

#### View Order
<img width="1420" alt="Screenshot 2025-05-29 at 3 56 34 PM" src="https://github.com/user-attachments/assets/e622cc76-aaa1-4508-9842-6170d75dc83d" />

<!-- ## ✅ Todo

- Add unit tests
- Dockerize the app
- Add Swagger docs -->

## 📄 License

MIT

---

Made with ❤️ by [Akshat-AFB](https://github.com/Akshat-AFB)
