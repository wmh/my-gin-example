# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0] - 2025-12-03

### Added
- 🚀 Upgraded to Go 1.23 with latest dependencies
- 🔐 JWT authentication with role-based access control (RBAC)
- 📦 Full CRUD operations for Users and Products
- 🗄️ GORM integration with SQLite (supports MySQL/PostgreSQL)
- ✅ Request validation using Gin's binding and go-playground/validator
- 📄 Pagination and search functionality for list endpoints
- 🔒 Rate limiting middleware to prevent abuse
- 🧪 Comprehensive unit tests for controllers
- 🌐 WebSocket support for real-time communication
- 🐳 Docker support with Dockerfile and docker-compose
- 📝 Enhanced structured logging with zerolog
- ⚙️ Configuration management with Viper
- 📚 Extensive API documentation and examples
- 🛠️ Makefile for common development tasks
- 🎯 Example scripts for testing all API endpoints

### Changed
- Replaced deprecated `io/ioutil` with `io` package
- Updated all dependencies to their latest stable versions:
  - gin-gonic/gin: v1.7.4 → v1.10.0
  - rs/zerolog: v1.23.0 → v1.33.0
  - spf13/viper: v1.8.1 → v1.19.0
  - stretchr/testify: v1.7.0 → v1.9.0
- Improved project structure with separate models, services, and controllers
- Enhanced error handling and logging throughout the application
- Better graceful shutdown with context timeout

### Fixed
- Security vulnerabilities in outdated dependencies
- Proper HTTP status codes for all error scenarios
- Database connection handling and pooling

### Security
- Added JWT token-based authentication
- Implemented role-based access control
- Added rate limiting to prevent DDoS attacks
- Password hashing with bcrypt
- Input validation on all endpoints

## [1.0.0] - 2021-08-17

### Added
- Initial project setup with Gin framework
- Basic health check endpoints
- Example API with simple authentication
- Configuration management with Viper
- Logging with zerolog
- Basic test coverage
- CircleCI integration

---

## Migration Guide from v1 to v2

### Breaking Changes
- Minimum Go version is now 1.23 (was 1.15)
- Some v1 endpoints are preserved under `/v1/` prefix
- New endpoints use `/v2/` prefix

### New Features to Integrate
1. **Authentication**: All v2 endpoints require JWT tokens (except registration/login)
2. **Database**: SQLite database is now required (auto-created on first run)
3. **Configuration**: New JWT and database settings in `config/app.toml`

### Update Steps
1. Update Go to version 1.23+
2. Run `go mod tidy` to update dependencies
3. Update configuration file with new settings
4. Migrate any custom code to use new patterns
5. Run tests to verify compatibility

For questions or issues, please open a GitHub issue.
