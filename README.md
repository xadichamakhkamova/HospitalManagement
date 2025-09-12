# Hospital Management System

This is a **microservices-based Hospital Management System** implemented in **Go**. The system provides comprehensive management of hospital operations including administration, doctors, nurses, patients, and pharmacists. The services communicate via **gRPC** and are exposed through an **API Gateway** with authentication and authorization mechanisms.

---

### Services

1. **Admin Panel**
   - Manages departments, personals, doctors, and beds.
   - Handles CRUD operations and storage.

2. **API Gateway**
   - Serves as the entry point for all clients.
   - Routes requests to appropriate microservices.
   - Provides JWT-based authentication and authorization.
   - Integrates Swagger documentation for API endpoints.

3. **Doctor Service**
   - Handles appointments and prescriptions.
   - Communicates with Admin and Patient services.

4. **Nurse Panel**
   - Manages donor data and related operations.

5. **Patient Management Service**
   - Handles patient records and personal data.

6. **Pharmacist Service**
   - Manages medicines and pharmaceutical inventory.

---

## ⚙️ Features

- **Microservices Architecture** using gRPC
- **JWT Authentication** for secure access
- **Role-Based Authorization**
- **Database Migrations** via SQL scripts
- **Swagger Documentation** for API endpoints
- **HTTPS Support** with `mkcert` self-signed certificates
- **Logging** with `logrus` in JSON format

---
