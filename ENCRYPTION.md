## Encryption Overview

This document provides an in-depth explanation of the encryption and security mechanisms implemented in this project.

---

### Password Hashing (bcrypt)
All user passwords are hashed using bcrypt before being stored in the database. This ensures:

- **Plaintext passwords are never saved**: Only the hashed version of the password is stored, making it impossible to retrieve the original password even if the database is compromised.
- **Salted hashes**: bcrypt automatically generates a unique salt for each password, ensuring that even identical passwords produce different hashes. This prevents rainbow table attacks.
- **Work factor (cost)**: The bcrypt algorithm includes a configurable cost parameter (set to 14 in this project). A higher cost increases the computational effort required to hash and verify passwords, making brute-force attacks more difficult.

The Go implementation uses the `golang.org/x/crypto/bcrypt` package, which is a widely trusted library for secure password hashing.

#### Why bcrypt?
- bcrypt is resistant to brute-force attacks due to its computational cost.
- It has been extensively tested and is widely used in production systems.

---

### JWT (JSON Web Tokens)
After a successful login, users receive a signed JWT (JSON Web Token) for authentication. JWTs are used to securely transmit information between the client and server.

#### Key Features of JWTs in This Project:
- **Claims**: The token includes claims such as:
  - `sub` (subject): Represents the user ID.
  - `exp` (expiration): Specifies the token's expiration time, ensuring tokens cannot be used indefinitely.
- **Signing Algorithm**: The tokens are signed using HMAC-SHA256, a secure hashing algorithm. This ensures:
  - The token cannot be tampered with. Any modification to the token invalidates it.
  - The server can verify the token's authenticity using a shared secret key (`jwtKey`).
- **Expiration**: Tokens are configured to expire after 24 hours, reducing the risk of misuse if a token is leaked.

#### Token Validation
The `ParseJWT` function in the [`utils/jwt.go`](utils/jwt.go) file validates tokens by:
1. Verifying the signature using the secret key.
2. Checking the expiration time to ensure the token is still valid.

#### Why JWT?
- JWTs are stateless, meaning the server does not need to store session data.
- They are compact and can be easily transmitted in HTTP headers.
- They are tamper-proof due to their cryptographic signature.

---

### HTTPS Recommendation
This project is designed to run locally on `http://localhost:8080`. However, for production environments, it is critical to use HTTPS to encrypt all traffic between the client and server. HTTPS ensures:
- **Data confidentiality**: Prevents eavesdropping on sensitive information such as passwords and tokens.
- **Data integrity**: Protects against data tampering during transmission.
- **Authentication**: Verifies the server's identity to prevent man-in-the-middle attacks.

#### Steps for Production Security:
1. Obtain an SSL/TLS certificate from a trusted certificate authority (CA).
2. Configure your server to use HTTPS.
3. Use a secure secret management solution (e.g., `.env` files) to store sensitive keys like `jwtKey`.

---

### SQLite Database Encryption (Optional)
While SQLite is used as the database in this project, it does not natively encrypt data. If encryption is required:
- Use an encrypted SQLite extension such as [SQLCipher](https://www.zetetic.net/sqlcipher/).
- Alternatively, encrypt sensitive data before storing it in the database.

---

### Future Enhancements
To further improve security, implementing:
- **Refresh Tokens**: Use refresh tokens to allow users to obtain new access tokens without re-authenticating.
- **Rate Limiting**: Limit the number of login attempts to prevent brute-force attacks. Distributed Systems Project.
- **IP Whitelisting**: Restrict access to certain endpoints based on IP addresses.
- **Environment-Specific Secrets**: Use different JWT secrets for development, staging, and production environments.