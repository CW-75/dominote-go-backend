# CLAUDE.md - Go EDA Backend Guidelines

## 🚀 Project Commands
- Run local: `go run cmd/server/main.go`
- Run with Docker: `docker-compose up --build`
- Run Tests: `go test -v -race ./...`
- Format Code: `go fmt ./... && goimports -w .`
- Linter: `golangci-lint run`

## 🏗️ Event-Driven Architecture (EDA)
1. **Decoupling:** Services must not communicate directly. All data flows via events using a message broker (e.g., Kafka, RabbitMQ) or internal Go channels.
2. **Directory Structure:**
   - `cmd/`: Application entry points.
   - `internal/domain/`: Pure domain models, match game structs, and event contracts (interfaces). Strict zero-external-dependency policy.
   - `internal/infrastructure/`: Database, Docker, Kubernetes, and message broker implementations.
   - `internal/event/`: Event producers, consumers, and event handlers.
3. **Immutability:** Match events represent facts about the past. Once published, event payloads/structs must be treated as read-only.

## 🛠️ Design Principles (SOLID, DRY, KISS, YAGNI)
- **KISS & YAGNI:** Write simple, straightforward Go code. Do not introduce premature abstractions, interfaces, or microservices if a modular monolith leveraging goroutines suffices for current requirements.
- **DRY:** Centralize parsing and validation logic for match records, but do not combine different domain structs just because they look similar.
- **SOLID (Single Responsibility & Interface Segregation):**
  - Each struct or package must have a single responsibility (e.g., `Parser`, `Repository`, `Publisher`).
  - Define small interfaces (1 or 2 methods max) in the consumer package where they are used, not where they are implemented.

## ⚡ High-Performance Go Best Practices
1. **Memory Allocation Optimization:**
   - Pre-allocate slices and maps using `make([]T, 0, capacity)` or `make(map[K]V, capacity)` when processing match log batches to prevent heap reallocation.
   - Pass large structs by pointer (`*Match`) to avoid expensive stack/memory copying, but pass small structs by value to keep allocations on the stack.
   - Use `sync.Pool` to reuse byte buffers or JSON objects if handling high-throughput log ingestions.
2. **Safe Concurrency:**
   - Process match logs concurrently using a Worker Pool bounded by `sync.WaitGroup` or buffered channels.
   - Avoid memory contention; follow the Go proverb: *"Do not communicate by sharing memory; instead, share memory by communicating."*
3. **Error Handling:** Explicit and clear, no exceptions. Use `errors.Is` or `errors.As` and wrap contexts with `fmt.Errorf("...: %w", err)`. Never use `panic()` in handlers or services.

## 🐳 Containerization & Orchestration (Docker & Kubernetes)
1. **Docker Multi-stage Builds:**
   - Stage 1 (`builder`): Use `golang:1.26-alpine`. Compile a statically linked binary with CGO disabled (`CGO_ENABLED=0`) and stripped flags (`-ldflags="-s -w"`).
   - Stage 2 (`final`): Use `scratch` or `alpine:latest` to keep the production image size under 30MB, minimizing the attack surface and deployment lag.
2. **Kubernetes Readiness:**
   - Implement health check endpoints (`/healthz` for Liveness, `/readyz` for Readiness).
   - Handle Graceful Shutdown by capturing OS signals (`SIGTERM`, `SIGINT`) to flush pending event buffers and drain DB connections cleanly before the Pod terminates.

## 🎯 Feature-Driven Development (FDD)
1. **Domain Object Modeling:** The first step for any new feature is modeling the domain objects. Features must be based on a clear and robust domain model.
2. **Feature Lists:** Break down requirements into small, client-valued features. Use the format: `<action> the <result> <by|for|of|to> a(n) <object>` (e.g., "Calculate the total score of a game table").
3. **Plan by Feature:** Group features together and plan their iterative development based on dependencies.
4. **Design & Build by Feature:**
   - **Design:** Define the exact structs, interfaces, and methods required for the feature before coding.
   - **Build:** Implement the code, write unit tests, and perform continuous integration. Avoid long-lived development branches; code should be integrated as soon as the feature is complete and tested.

## 🤖 Agent Session Continuity (AGENTS.md)
1. **End-of-Session Protocol:** When the user indicates a chat session is ending, wrapping up, or switching contexts, the AI **MUST** write or append a summary to a file named `AGENTS.md` in the root directory.
2. **`AGENTS.md` Structure:** The AI must update the file using the following schema:
   - **Current State:** Brief summary of what was accomplished in this session.
   - **Architecture Decisions:** Any new structs, event schemas, or patterns agreed upon.
   - **Next Steps:** Bullet points detailing what the next AI agent or session should focus on immediately.
   - **Known Blockers/Tech Debt:** Leftover TODOs or performance trade-offs made during the session.
3. **Goal:** Ensure the next LLM session can read `AGENTS.md` and instantly pick up the work without losing context.