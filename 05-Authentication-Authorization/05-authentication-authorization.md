# 🔐 Authentication and Authorization in Backend Engineering

Modern applications rely on **authentication** (verifying identity) and **authorization** (controlling access). While they often go hand in hand, they solve different problems:  
- **Authentication** → Who are you?  
- **Authorization** → What are you allowed to do?  

---

## 🚪 Authentication Methods

### 1. OAuth 2.0
- Provides secure delegated access without sharing passwords.  
- Uses access tokens issued by **Authorization Server**.  
- Popular for third-party logins (Google, GitHub, Discord).  

### 2. JWT (JSON Web Token)
- Compact, stateless tokens containing user info.  
- Ideal for **scalable microservices**.  
- Stored in cookies, headers, or (not recommended) localStorage.  

### 3. Zero Trust Architecture
- “Never trust, always verify.”  
- Every request is authenticated and authorized, regardless of network location.  

### 4. Passwordless Authentication
- Email magic links, OTPs, or biometric-based logins.  
- Removes risks of weak/stolen passwords.  

### 5. Future Trends
- **Decentralized Authentication** (Blockchain-based identity).  
- **Behavioral Biometrics** (typing patterns, mouse movements).  
- **Post-Quantum Cryptography** (resistant to quantum attacks).  

---

## 🍪 Sessions, JWTs, and Cookies

### Sessions (Stateful)
- HTTP is stateless → servers need a way to **remember users**.  
- A session ID is created and stored in Redis or memory DB.  
- Sent to the client via **cookies**.  
- Session expires after timeout (e.g., 15 mins).  
- Scales better with distributed stores like Redis/Memcache.  

### JWTs (Stateless)
- Useful for distributed, scalable systems.  
- No server memory required; servers validate JWT with a shared secret.  
- **Challenges:** token theft → can’t revoke until expiry.  

### Cookies
- Store **session IDs** or **JWTs** on the client.  
- Sent with every request to maintain state.  

---

## 🧩 Types of Authentication

1. **Stateful Auth** → Sessions (web apps).  
2. **Stateless Auth** → JWTs (mobile, APIs).  
3. **API Keys** → Machine-to-machine (e.g., accessing OpenAI API).  
4. **OAuth 2.0 & OIDC** → Delegated authorization + authentication.  

---

## 🔑 OAuth Deep Dive

### OAuth 1.0
- Used cryptographic signatures (complex, error-prone).  
- Replaced by OAuth 2.0.  

### OAuth 2.0
- Uses **bearer tokens** (simpler, more vulnerable).  
- Four main flows:  
  1. Authorization Code Flow (web apps).  
  2. Implicit Flow (browser apps, now discouraged).  
  3. Client Credentials (server-to-server).  
  4. Device Code Flow (smart TVs).  

### OpenID Connect (OIDC)
- Layer on top of OAuth 2.0 → adds **authentication**.  
- Introduces **ID Token (JWT)** with user info.  

---

## 🛡️ Authorization & RBAC

- **Authorization** = defining *what actions a user can perform*.  
- Often role-based (RBAC):  
  - Admin → Read/Write/Delete  
  - User → Read/Write  
  - Guest → Read-only  
- Unauthorized actions return **403 Forbidden**.  

---

## ⚠️ Security Concerns

### Error Messages
- Avoid leaking sensitive details:  
  ❌ “User not found” or “Wrong password.”  
  ✅ “Authentication failed.”  

### Timing Attacks
- Attackers analyze response times to guess credentials or hashing algorithms.  
- Mitigation: **constant-time cryptographic checks** and simulated delays.  

---

## 🎯 Key Takeaways
- Use **sessions** for simple web apps, **JWTs** for distributed systems.  
- Adopt **OAuth 2.0 + OIDC** for third-party integrations.  
- Use **API keys** for server-to-server communication.  
- Always consider **security best practices** (error handling, timing attack protection).  
- Future: **passwordless, blockchain identity, quantum-safe cryptography**.  
