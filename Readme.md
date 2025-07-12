# 🟢 Let's Go Snippets — A Go Web App for Sharing Code Snippets

A simple snippet-sharing web application built in **Go**, inspired by the book **Let's Go by Alex Edwards**.

This project helped me learn the foundational concepts of modern web development using Go — from routing and templates to secure sessions and database integration.

---

## ✨ What I Learned

✅ Key backend web development concepts in Go:

- Building HTTP servers using `net/http`
- HTML templating with Go's `html/template` package
- Form decoding and validation
- Secure session management with `scs`
- Authentication and hashed password storage
- Environment-based configuration with `.env`
- Connecting and querying a MariaDB/MySQL database
- Structuring a modular and maintainable Go web application
- Dockerizing Go apps for easy deployment

---

## 🧠 Tech Stack

- **Language:** Go 1.24+
- **Database:** MariaDB (or MySQL)
- **Router:** [`httprouter`](https://github.com/julienschmidt/httprouter)
- **Session Manager:** [`scs`](https://github.com/alexedwards/scs)
- **Form Decoder:** [`github.com/go-playground/form/v4`](https://github.com/go-playground/form)
- **Env Loader:** [`github.com/joho/godotenv`](https://github.com/joho/godotenv)
- **CSRF Protection:** [`nosurf`](https://github.com/justinas/nosurf)
- **Middleware:** [`alice`](https://github.com/justinas/alice)
- **Containerization:** Docker & Docker Compose

---

## 🗂️ Project Structure

```
lets_go/
├── cmd/
│   └── web/                    # Main application entry point
│       ├── context.go          # Request context helpers
│       ├── handlers.go         # HTTP handlers
│       ├── helpers.go          # Template and error helpers
│       ├── main.go             # Application entry point
│       ├── middleware.go       # Custom middleware
│       ├── routes.go           # Route definitions
│       └── templates.go        # Template rendering logic
├── internal/
│   ├── models/                 # Database models and logic
│   │   ├── errors.go          # Custom error types
│   │   ├── snippet.go         # Snippet model
│   │   └── users.go           # User model
│   └── validator/             # Custom validation helpers
│       └── validator.go
├── ui/
│   ├── html/                  # HTML templates
│   │   ├── base.tmpl.html     # Base template layout
│   │   ├── pages/             # Page templates
│   │   └── partials/          # Reusable template components
│   └── static/                # Static assets
│       ├── css/
│       ├── img/
│       └── js/
├── db/
│   └── init.sql               # Database schema initialization
├── tls/                       # TLS certificates (optional)
│   ├── cert.pem
│   └── key.pem
├── docker-compose.yaml        # Docker Compose configuration
├── Dockerfile                 # Docker build configuration
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
└── README.md                  # This file
```

---

## ⚙️ How to Run

### 🔹 Option 1: Local Setup (Recommended for Learning)

#### 📦 Prerequisites
- Go 1.24+
- MariaDB or MySQL installed
- `git` and `mysql` CLI (optional but helpful)

#### ✅ Steps

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd lets_go
   ```

2. **Create the Database**
   ```bash
   mysql -u root -p < db/init.sql
   ```

3. **Create a `.env` file in the root**
   ```env
   DB_DSN=web:123456@/snippetbox?parseTime=true
   ```

4. **Install Dependencies**
   ```bash
   go mod download
   ```

5. **Run the Application**
   ```bash
   go run ./cmd/web
   ```

6. **Open Your Browser**
   
   Visit: http://localhost:4000

---

### 🔹 Option 2: Docker Setup (Quick & Easy)

#### 📦 Requirements
- Docker & Docker Compose installed

#### ✅ Steps

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd lets_go
   ```

2. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

3. **Open Your Browser**
   
   Visit: http://localhost:4000

**The Docker setup includes:**
- Go backend server (auto-built from source)
- MariaDB container with preloaded schema from `db/init.sql`
- No need to install Go or MySQL manually!

---

## 🔐 Optional: Enable TLS (HTTPS)

By default, the application runs over HTTP. To enable HTTPS:

### 1. 🗂 Create TLS Certificates

From the project root:

```bash
mkdir -p tls
openssl req -x509 -newkey rsa:4096 -nodes \
  -keyout tls/key.pem -out tls/cert.pem -days 365
```

*Just press Enter to accept the defaults when prompted.*

### 2. ✅ Enable TLS in `main.go`

Uncomment this line in `cmd/web/main.go`:

```go
// err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
```

And comment out this one:

```go
err = srv.ListenAndServe()
```

### 3. 🔗 Access via HTTPS

Visit: https://localhost:4000

> **Note:** Your browser will warn you about an untrusted certificate — it's safe to proceed for local testing.

---

## 🌟 Features

- **Snippet Management:** Create, view, and manage code snippets
- **User Authentication:** Secure signup and login system
- **Session Management:** Secure session handling with CSRF protection
- **Responsive Design:** Clean, mobile-friendly interface
- **Database Integration:** Persistent storage with MariaDB/MySQL
- **Form Validation:** Client and server-side validation
- **Security:** Password hashing, CSRF protection, and secure headers

---

## 📖 Learning Resources

This project is based on **"Let's Go" by Alex Edwards** - an excellent book for learning Go web development. Check it out at [lets-go.alexedwards.net](https://lets-go.alexedwards.net/).

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📬 Contact

- **GitHub:** [@muskiteer](https://github.com/muskiteer)


---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- **Alex Edwards** for the excellent "Let's Go" book
- The Go community for amazing packages and documentation
- Contributors to all the open-source packages used in this project