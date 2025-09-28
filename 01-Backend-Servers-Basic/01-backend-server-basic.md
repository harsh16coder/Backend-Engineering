# Understanding Backend Servers in Web Applications

A backend server plays a crucial role in web applications by handling
requests, processing data, and serving content to clients. At its core,
the backend is a computer system that listens for requests over the
internet using protocols such as HTTP, WebSockets, or gRPC. When a
client --- typically a web browser or mobile app --- makes a request,
the backend responds by providing static files like images or dynamic
data in formats such as JSON.

When a request is made, it follows a specific flow: it starts from the
browser, travels through DNS servers, reaches cloud infrastructure like
AWS EC2 instances, passes through security layers such as firewalls and
reverse proxies (e.g., NGINX), and finally arrives at the backend
server, which often runs code written in languages like Golang. DNS
servers play the important role of mapping human-readable domain names
to machine-readable IP addresses, allowing the browser to locate the
correct server. On cloud platforms like AWS, security groups regulate
which ports are open, ensuring that only authorized traffic reaches the
server. Reverse proxies like NGINX handle tasks such as managing SSL
certificates and redirecting requests to the appropriate backend
application.

It is important to understand the distinction between frontend and
backend. The frontend refers to the user interface and code that runs
directly in the user's browser. While the backend serves the frontend
code and provides data, the frontend operates in a sandboxed environment
with strict security limitations. Browsers enforce security policies
such as Cross-Origin Resource Sharing (CORS) to restrict how web pages
can interact with external resources, preventing malicious actions.

Frontend environments have several key limitations compared to the
backend. They lack direct access to the file system, cannot maintain
persistent database connections, and are restricted by the device's
limited computing resources. Additionally, executing sensitive business
logic or handling secure data directly in the frontend is a major
security risk. Instead, the backend provides a centralized environment
to execute business logic securely, manage persistent connections to
databases, handle large-scale client requests efficiently, and maintain
overall system security.

In summary, the backend is an essential component of modern web
applications. It centralizes data storage, executes complex business
logic, manages client-server communication, and ensures security.
Understanding the backend's role is a foundational step toward mastering
backend engineering principles and developing robust web applications.

------------------------------------------------------------------------

## Key Insights from the Reference Video

The video provides a comprehensive explanation of what a backend server
is, how it operates, and why it is necessary in web applications.
Starting from a basic definition, the backend is described as a computer
that listens for requests over the internet (HTTP, websockets, gRPC) and
serves content such as static files or JSON data. The video then dives
into the detailed flow of a request, tracing it from the browser,
through DNS servers, AWS EC2 instances, firewalls, reverse proxies (like
NGINX), and finally to the Node.js backend server running locally on the
instance. It explains how DNS maps domain names to IP addresses, how AWS
security groups control port access, and how NGINX manages SSL and
request redirection.

The video also contrasts backend servers with frontend applications by
demonstrating how frontend code is served by the backend but executed on
the client's browser, highlighting the sandboxed and restricted nature
of browser environments. The video explains browser security policies
including CORS, and why frontend environments cannot replace backends.
It details the limitations of running backend logic in the frontend ---
such as security concerns, lack of file system access, inability to
maintain persistent database connections, and limited computing
resources on client devices.

Ultimately, the backend is essential for managing centralized data
storage, executing business logic, handling multiple client requests
efficiently, and maintaining security. The video concludes by
positioning this foundational knowledge as a crucial starting point for
learning backend engineering principles and practices.

------------------------------------------------------------------------

### 🌍 Request Lifecycle from Browser to Backend

The video meticulously traces a user request starting from a browser,
passing through DNS resolution, firewall filtering, and reverse proxy
redirection, until it arrives at the backend server running Node.js.
This showcases the complexity and multiple layers involved in serving
web content, emphasizing the importance of each component in ensuring
seamless communication. Understanding this flow is fundamental for
backend engineers to troubleshoot and optimize web applications.

### 🔐 Role of Firewalls and Security Groups

AWS security groups function as virtual firewalls that define which
ports are open to the internet, directly impacting whether requests can
reach the backend server. This layer is crucial in protecting backend
infrastructure from unauthorized access and potential attacks,
underscoring the security responsibilities developers must consider when
designing backend deployments.

### 🔄 Reverse Proxy and SSL Management

Using NGINX as a reverse proxy simplifies managing SSL certificates and
consolidates configuration for multiple backend services. This approach
improves security by enforcing HTTPS and provides flexibility in routing
requests, which is vital in scalable and maintainable backend
architectures.

### 🖥️ Difference Between Backend and Frontend Execution

Frontend applications deliver JavaScript code to browsers, where it runs
locally, while backend servers process requests and data centrally. This
distinction affects performance, security, and capabilities --- with
backend servers able to handle sensitive operations and maintain
databases, while frontend environments are constrained by sandboxing.

### 🔒 Browser Security and CORS

The browser enforces cross-origin resource sharing (CORS) policies to
protect users from malicious scripts accessing unauthorized data. This
restriction is a key reason backend logic cannot be executed safely on
the frontend, as it limits the frontend's ability to communicate freely
with external systems. Backend servers can securely handle these
cross-domain interactions.

### 🗄️ Database Access and Connection Pooling

Backend servers use native database drivers to maintain persistent
connections (connection pools), allowing efficient and scalable database
queries. Browsers lack this capability, and opening individual
connections from each client would overwhelm databases. This technical
limitation necessitates the backend as the intermediary for data
operations.

### ⚙️ Computing Resources and Scalability

Client devices vary widely in computing power and memory, often
insufficient for heavy business logic or processing large volumes of
data. Backend servers provide centralized, scalable computing resources
that can be upgraded as needed, ensuring consistent performance across
all users regardless of their devices.

### 🔄 Centralization of Data and Logic

Backend servers act as the centralized authority that stores user data,
enforces business logic, and manages state across many users. This
centralization is essential for personalized experiences, such as
Instagram's like-notification flow, where the server coordinates data
updates and triggers notifications reliably.

### 🤔 Why Not Put Backend Logic in Frontend?

The video addresses this common question by highlighting the multiple
restrictions and security risks of running backend logic in the frontend
environment, including sandboxing, lack of database connectivity, and
potential data exposure, making backend servers indispensable.

### 🚀 Foundational Knowledge for Backend Engineering

By understanding these backend fundamentals --- request flow, security,
proxies, frontend-backend differences, and database communication ---
learners are well equipped to begin mastering backend engineering,
building secure, scalable, and performant applications.

------------------------------------------------------------------------

## Highlights

-   🌐 Backend servers listen for internet requests (HTTP, websockets,
    gRPC) and serve or accept data.
-   🖥️ DNS servers translate domain names to IP addresses pointing to
    backend servers hosted on cloud instances like AWS EC2.
-   🔥 AWS security groups act as firewalls controlling which ports
    (HTTP/HTTPS) are accessible externally.
-   ⚙️ Reverse proxies (NGINX) handle SSL termination and route incoming
    requests to local backend services.
-   💻 Frontend apps execute JavaScript code in browsers, which are
    sandboxed and limited environments.
-   🔒 Browser security policies (like CORS) restrict frontend code from
    freely accessing external APIs or system resources.
-   🗄️ Backend servers are necessary for secure data persistence,
    managing database connections, and handling heavy computation.

------------------------------------------------------------------------

## Reference Video

🔗 [What is Backend
Engineering?](https://www.youtube.com/watch?v=6Ss4dJD9Kzg&list=PLui3EUkuMTPgZcV0QhQrOcwMPcBCcd_Q1&index=3)
