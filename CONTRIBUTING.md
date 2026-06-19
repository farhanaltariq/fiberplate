# Contributing to Fiberplate

Welcome! This guide helps you contribute to the codebase without breaking the project's patterns and structures.

## Core Architectural Layers

We follow a structured layering pattern. Do not bypass layers (e.g., calling services directly from routes, or query DB directly from controllers).

```
[ HTTP Request ] 
       │
       ▼
 [ routes/ ]       <-- Define HTTP paths, methods, and attach middleware
       │
       ▼
 [ controllers/ ]  <-- Parse requests, call services, handle HTTP responses
       │
       ▼
 [ services/ ]     <-- Core business logic, transaction handling
       │
       ▼
 [ models/ ]       <-- Database schemas and GORM declarations
```

---

## Step-by-Step Guideline: Adding a Feature

### 1. Database Model
* Create schema in `app/database/models/your_feature.go`.
* Register migrations inside `app/database/connection.go` (if auto-migration is used).

### 2. Services
* Define an interface in `app/services/your_feature.go`.
* Implement the interface structure and create a constructor function `NewYourFeatureService(db *gorm.DB) YourFeatureService`.
* Register the new service in `app/middleware/bootstrap.go` under the `Services` struct.
* Initialize the service in `app/routes/routes.go` inside the `initServices()` function.

### 3. Controllers
* Define a controller interface in `app/controllers/your_feature_controller.go`.
* Implement interface methods binding them to `*controller` (defined in `controllers.go`).
* Expose a constructor function `NewYourFeatureController(service middleware.Services) YourFeatureController`.

### 4. Routes
* Map routes in `app/routes/your_feature.go`.
* Register your routing file/group within `app/routes/routes.go` inside the `Init()` function.

---

## Coding Standards & Conventions

### Directory Layout
Keep all source code in the `app/` directory. Do not add top-level directories unless explicitly agreed upon.

### File Naming
* Use **camelCase** for filenames (e.g., `authenticationController.go`, `userController.go`).

### Code Formatting
* Run `go fmt ./...` and `go vet ./...` before committing.
* Ensure code is clean, self-documenting, and dependencies are correctly ordered.

### Git Commits
Use semantic/conventional commit messages:
* `feat: ...` for new features.
* `fix: ...` for bug fixes.
* `refactor: ...` for structural rewrites.
* `docs: ...` for documentation updates.
