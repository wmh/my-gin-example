# Changelog

All notable changes to this project will be documented in this file.

## [2.0.0] - 2025-12-07

### Added
- 🚀 Upgraded to Go 1.24 with latest stable dependencies
- 🔐 JWT authentication with role-based access control (RBAC)
- 📦 Full CRUD operations for Users and Products
- 📁 File upload management system with security validation
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
- ✅ Complete CI/CD with GitHub Actions and CircleCI
- 🔍 CodeQL security analysis integration
- 🛡️ Dependabot for automated dependency updates

### Changed
- Replaced deprecated `io/ioutil` with `io` package
- Updated all dependencies to their latest secure versions:
  - gin-gonic/gin: v1.7.4 → v1.10.0
  - rs/zerolog: v1.23.0 → v1.33.0
  - spf13/viper: v1.8.1 → v1.19.0
  - stretchr/testify: v1.7.0 → v1.9.0
  - golang-jwt/jwt: v5.2.1 → v5.2.2
  - golang.org/x/crypto: v0.40.0 → v0.45.0
- Improved project structure with separate models, services, and controllers
- Enhanced error handling and logging throughout the application
- Better graceful shutdown with context timeout
- Hardened file permissions (0600 for files, 0750 for directories)
- Added HTTP server timeouts for security

### Fixed
- All security vulnerabilities (14 CVEs resolved)
- All gosec warnings (11 issues resolved)
- Proper HTTP status codes for all error scenarios
- Database connection handling and pooling
- Path traversal vulnerabilities in file operations
- Error handling for all file operations

### Security
- ✅ Zero vulnerabilities (verified by govulncheck)
- ✅ gosec compliant (all warnings resolved)
- ✅ Path traversal protection for file uploads
- ✅ Secure file permissions (0600/0750)
- ✅ HTTP timeout configuration (anti-Slowloris)
- ✅ JWT token-based authentication
- ✅ Role-based access control
- ✅ Rate limiting to prevent DDoS attacks
- ✅ Password hashing with bcrypt
- ✅ Input validation on all endpoints

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
- Minimum Go version is now 1.24 (was 1.15)
- File permissions changed to more secure defaults
- Some HTTP timeouts added (may affect long-running requests)

### New Features to Integrate
1. **Authentication**: All protected endpoints require JWT tokens
2. **File Upload**: New file management system with security validation
3. **Database**: SQLite database is now required (auto-created on first run)
4. **Configuration**: New JWT, database, and upload settings in `config/app.toml`

### Update Steps
1. Update Go to version 1.24+
2. Run `go mod tidy` to update dependencies
3. Update configuration file with new settings (especially upload config)
4. Ensure proper file permissions on data directories
5. Run security checks: `govulncheck ./...` and `gosec ./...`
6. Run tests to verify compatibility

For questions or issues, please open a GitHub issue.
