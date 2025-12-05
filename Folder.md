### Root files
```
README.md          → Project documentation, how to run, API examples
go.mod, go.sum     → Go module definition and dependencies
```

### cmd/api/main.go
**Purpose**: Entry point of the application  
**What you do here**:
- Load config
- Initialize logger
- Connect to PostgreSQL + Redis
- Create repositories, services, handlers
- Set up router with middlewares
- Start HTTP server

```go
func main() {
    cfg := config.Load()
    logger := logger.New()
    db := postgres.Connect(cfg)
    redisClient := redis.NewClient(...)
    
    userRepo := postgres.NewUserRepository(db)
    pasteRepo := postgres.NewPasteRepository(db)
    cache := redis.NewCache(redisClient)

    authService := service.NewAuthService(userRepo, cfg)
    pasteService := service.NewPasteService(pasteRepo, cache)

    handlers := handlers.New(pasteService, authService)
    router := router.SetupRouter(handlers)

    server := server.New(router, cfg.Port)
    server.Start()
}
```

### internal/config/config.go
Loads environment variables or config file

```go
type Config struct {
    Port           string
    JWTKey         string
    PostgresURL    string
    RedisURL       string
    RateLimitRPS   int
}

func Load() *Config { ... } // uses viper or env
```

### internal/http/
#### handlers/*.go
Thin layer — only extract from request → call service → return JSON

**paste_handler.go example**:
```go
func (h *PasteHandler) CreatePaste(c *gin.Context) {
    var req CreatePasteRequest
    if err := c.ShouldBindJSON(&req); err != nil { ... }

    paste, err := h.pasteService.Create(c.Request.Context(), userID, req)
    if err != nil { response.Error(c, err); return }

    response.OK(c, paste)
}
```

#### middlewares/
- `auth_middleware.go` → checks JWT, puts userID into context
- `rate_limit.go` → per-IP or per-user rate limiting (using Redis)
- `logging.go` → logs request method, path, latency, status

#### router.go
Sets up Gin (or Chi/Fiber) routes and applies middlewares

```go
func SetupRouter(h *handlers.Handler) *gin.Engine {
    r := gin.New()
    r.Use(middlewares.Logging(), middlewares.Recovery())

    public := r.Group("/")
    public.POST("/register", h.User.Register)
    public.POST("/login", h.User.Login)

    protected := r.Group("/")
    protected.Use(middlewares.Auth())
    protected.Use(middlewares.RateLimit())
    {
        protected.POST("/pastes", h.Paste.Create)
        protected.GET("/pastes/:id", h.Paste.Get)
        // ...
    }
    return r
}
```

### internal/domain/
Pure business models and interfaces (no framework, no DB)

#### domain/user/
```go
// model.go
type User struct {
    ID        string
    Email     string
    Password  string // hashed
    CreatedAt time.Time
}

// repository.go – interface only!
type UserRepository interface {
    Create(ctx context.Context, u *User) error
    FindByEmail(ctx context.Context, email string) (*User, error)
}

// service.go – interface
type UserService interface {
    Register(ctx context.Context, email, password string) (*User, error)
    Login(ctx context.Context, email, password string) (string, error) // returns JWT
}
```

Same structure for `paste/`

### internal/repository/
Actual implementations of domain interfaces

#### repository/postgres/
```go
// user_repo.go
func (r *UserPostgresRepo) Create(ctx context.Context, u *User) error {
    return r.db.WithContext(ctx).Create(u).Error
}
```

#### repository/redis/cache_repo.go
Used for caching hot pastes or rate-limit counters

```go
func (c *Cache) GetPaste(id string) (*domain.Paste, error) { ... }
func (c *Cache) SetPaste(p *domain.Paste, ttl time.Duration) error { ... }
```

#### repository/memory/paste_repo.go
In-memory version used only in tests

### internal/service/
Business logic lives here (uses repositories)

**auth_service.go**:
```go
func (s *AuthService) Register(email, password string) (*domain.User, error) {
    hashed, _ := bcrypt.GenerateFromPassword(...)
    user := &domain.User{ID: uuid.New(), Email: email, Password: hashed}
    return user, s.userRepo.Create(ctx, user)
}

func (s *AuthService) Login(...) (string, error) {
    user, err := s.userRepo.FindByEmail(...)
    if !checkPassword(...) { return "", ErrInvalidCredentials }
    return jwt.GenerateToken(user.ID, s.jwtKey)
}
```

**paste_service.go** – creates short ID, sanitizes content, sets expiration, etc.

### internal/utils/
Helper functions that don’t depend on external packages too much

- `id_generator.go` → nanoid or custom base62 short IDs ("abc123")
- `time.go` → Now() wrapper for testing
- `sanitizer.go` → bleach HTML, prevent XSS in pastes

### internal/app.go
Sometimes people put the "application wiring" here instead of main.go (dependency injection style)

### internal/server.go
Wrapper around http.Server with graceful shutdown

```go
func (s *Server) Start() {
    go func() { s.http.ListenAndServe() }()
    <-s.stopCh
    s.http.Shutdown(context.Background())
}
```

### pkg/
Code that could be reused by other projects (safe to import from outside)

- `pkg/jwt/token.go` → GenerateToken(), ValidateToken()
- `pkg/logger/` → zap or zerolog wrapper
- `pkg/response/` → standardized JSON responses (OK, Error, Pagination)
- `pkg/validation/` → custom validators for gin

### migrations/
SQL files for goose, migrate, or golang-migrate

```sql
-- 001_users.sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);
```

### build/
#### Dockerfile
Multi-stage build for small final image

#### docker-compose.yml
Spins up postgres + redis + the API

```yaml
services:
  api:
    build: .
    ports: ["8080:8080"]
    depends_on: [postgres, redis]
  postgres:
    image: postgres:16
    environment:
      POSTGRES_DB: pastebin
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
  redis:
    image: redis:7
```

### Summary table

| Path                              | Responsibility                                      |
|-----------------------------------|-------------------------------------------------------|
| `cmd/api/main.go`                 | Bootstrap everything and run server                   |
| `internal/config`                 | Load env variables                                    |
| `internal/http/handlers`          | HTTP → JSON glue code                                 |
| `internal/http/middlewares`       | Auth, rate limit, logging                             |
| `internal/domain`                 | Pure business models + interfaces                     |
| `internal/repository`             | Database implementations (Postgres, Redis, in-memory)|
| `internal/service`                | All business logic                                    |
| `internal/utils`                  | Small pure-Go helpers                                 |
| `pkg/*`                           | Reusable libraries (JWT, logger, response helpers)   |
| `migrations/`                     | Database schema changes                               |
| `build/`                          | Dockerfiles and compose                               |

