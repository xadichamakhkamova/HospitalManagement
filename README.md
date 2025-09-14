# 🏥 Hospital Management System

A **microservices-based Hospital Management System** built with **Go**.  
This project helps manage hospital operations such as administration, doctors, nurses, patients, and pharmacists.  
All services communicate using **gRPC** and are exposed via an **API Gateway**.

---

## 📌 Microservices

### 🔹 Admin Panel
- Manage **departments**, **personals**, **doctors**, and **beds**  
- CRUD operations with PostgreSQL

### 🔹 API Gateway
- Single entry point for clients  
- Routes requests to microservices  
- Integrated **Swagger documentation**  
- Supports **HTTPS** with self-signed certificates  

### 🔹 Doctor Service
- Handles **appointments** and **prescriptions**  
- Communicates with Admin & Patient services  

### 🔹 Nurse Panel
- Manages **donor data** and related operations  

### 🔹 Patient Management Service
- Handles **patients’ records** and personal data  

### 🔹 Pharmacist Service
- Manages **medicines** and pharmaceutical inventory  

---

## ⚙️ Features
- **Microservices Architecture** with gRPC  
- **PostgreSQL + SQLC** for database queries  
- **Database Migrations** via SQL scripts  
- **Swagger API Documentation**  
- **TLS/HTTPS** with `mkcert`  
- **Structured Logging** with `logrus` (JSON format)  

---
