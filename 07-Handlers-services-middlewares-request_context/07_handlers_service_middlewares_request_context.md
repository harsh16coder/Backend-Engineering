# Server-Side Request Lifecycle: Handlers, Services, Repositories, and Middleware

This document outlines the internal request lifecycle within a server, focusing on the roles of **Handlers (Controllers)**, **Services**, **Repositories**, **Middleware**, and **Request Context**.

---

## 1. The Internal Request Lifecycle 🌐

When a client sends an HTTP request to a server, an internal request lifecycle begins:

1.  **Port Reception:** The request hits the server's listening **port**.
2.  **Routing:** The port routes the request to an appropriate **route** based on the URL and HTTP method (GET, POST, etc.).
3.  **Handler/Controller Assignment:** The **Router** assigns a specific **Handler** or **Controller** function to process the request.
    * The Handler primarily receives two main objects: a **Request object** (containing all client data) and a **Response object** (used to build and send the server's reply).

---

## 2. Handlers (Controllers) 🖐️

The **Handler** (often called the **Controller**) is the initial entry point for application logic. It deals primarily with **input and output data formats** related to the HTTP protocol.

### Handler Responsibilities:

1.  **Data Extraction (Input):** Taking raw data from the HTTP Request object.
    * **GET:** Extract **query parameters**.
    * **POST/PUT/DELETE:** Extract data from the **request body**.
2.  **Deserialization (Binding):** Converting the raw input data (e.g., JSON) into a native programming language structure.
    * **Go:** Deserialized into a `struct`.
    * **Python:** Deserialized into a `dictionary`.
    * **Node.js/JavaScript:** Often no explicit deserialization is needed as JSON is a native object format.
    * *This step is also called **Binding**.*
3.  **Validation & Transformation:**
    * Validating that the data meets necessary constraints (e.g., email format, required fields).
    * Transforming the data to the required application format (e.g., date formats, case changes).
    * **Best Practice:** Make query parameters optional. If a client doesn't send one, apply a **default value** to prevent the validation pipeline from immediately throwing a `400 Bad Request` error.
4.  **Calling the Service Layer:** Once the data is solid, the controller hands it over to the **Service Layer** for actual business processing.
5.  **Sending the Response (Output):** Receives processed data from the Service Layer, modifies/writes **HTTP headers**, and provides the final response body to the client using the Response object.



---

## 3. Service Layer (Business Logic) ⚙️

The **Service Layer** is where the **actual processing (business logic)** happens.

### Key Principles:

* **Isolation:** The Service Layer should be completely **isolated** from HTTP/frontend/backend concerns. No HTTP-specific actions (like setting headers or dealing with request/response objects) should occur here.
* **Focus:** It represents the application's core logic.
* **Orchestration:** It takes data from the Handler, processes it, and may orchestrate calls to one or more Repository functions to perform database operations.

### Service Responsibilities:

* Executing core business rules (e.g., calculating a price, checking inventory, processing a payment).
* Handling external integrations (e.g., sending an email).
* Receiving data from various **Repository functions**, combining/orchestrating it, and passing the final result back to the Handler.

---

## 4. Repository Layer (Data Access) 💾

The **Repository Layer** is exclusively responsible for **all database calls (data persistence)**.

### Repository Principles:

* **Abstraction:** It hides the details of the database implementation (e.g., whether it's SQL or NoSQL) from the Service Layer.
* **Interface:** A good pattern is that **one repository method should return one kind of data** (or a collection of it).

### Repository Responsibilities:

* Receiving parameters (like sort criteria, IDs, or search terms) from the Service Layer.
* Constructing and executing database queries (e.g., generating a SQL query for a sort operation).
* Mapping database results back into application-friendly objects or data structures.

### The Standard Flow:

$$
\text{Client Request} \rightarrow \text{Handler} \rightarrow \text{Service} \rightarrow \text{Repository} \rightarrow \text{Database}
$$

---

## 5. Middleware 🚧

**Middleware** functions execute **in between** different layers of the API cycle. They are optional and depend entirely on the application's requirements.

### How Middleware Works:

* A middleware receives three main objects: the **Request** object, the **Response** object, and a **`next()`** function.
* The `next()` function is called to pass the request context to the next component (another middleware or the final handler).
* **Order Matters:** The sequence in which middlewares are defined dictates their execution order.

### Why Use Middleware?

Middleware is used to define common operations and avoid repetitive tasks across multiple handlers, following the principle of **Don't Repeat Yourself (DRY)**.

* **Efficiency:** A middleware can perform a task (like authentication) and, if it fails, immediately return a response (e.g., `401 Unauthorized`) without ever passing the request to the main handler. This saves precious server resources.

### Common Middleware Tasks:

| Category | Purpose / Examples |
| :--- | :--- |
| **Security** | **CORS** (checking allowed origins), **Authentication** (token validation, authorization), **Security Headers** (Content Security Policy), **Rate Limiting** (max requests from an IP). |
| **Logging & Monitoring** | Creating access logs (paths, client IP, timing) for debugging and auditing. |
| **Data Handling** | **Compression** (e.g., GZIP), **Data Parsing** (initial serialization/deserialization). |
| **Global Error Handling** | Typically kept as the **last** middleware. It catches generic errors thrown by any component further up the pipeline and provides a standardized, appropriate response to the client. |



---

## 6. Request Context 📦

**Request Context** is a temporary, isolated storage area (**state**) that is **scoped for a particular request**.

### Key Characteristics:

* **Key-Value Pair Storage:** Usually implemented as a simple map or key-value store.
* **Accessibility:** Easily accessible to **all subsequent middlewares and handlers** in the request chain.
* **Decoupling:** It allows data to be shared without closely coupling the different components.

### Example Use Case (Authentication):

1.  The **Authentication Middleware** receives an authorization token.
2.  It validates the token and securely extracts the **`userId`**.
3.  It stores this validated `userId` in the **Request Context**.
4.  Subsequent handlers or services can now retrieve the trustworthy `userId` from the context to query the database.
    * *This prevents a client from maliciously impersonating another user, as the ID used for database lookups comes from a secure, server-validated source (the context), not directly from a client-provided parameter.*

---