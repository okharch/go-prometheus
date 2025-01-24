# Simple Backend with Prometheus Client

This repository provides a simple Go backend application integrated with the Prometheus client library to expose standard application metrics. Additionally, it includes a handler that emulates resource loads such as CPU usage, goroutines, and memory allocations.

---

## Features

- **Prometheus Metrics**: Exposes runtime metrics like:
  - CPU usage
  - Number of goroutines
  - Heap memory allocations
- **Simulated Resource Usage**:
  - A handler (`/allocate`) to emulate CPU load, allocate memory, and increase goroutines.
  - Useful for testing how metrics change under different application loads.
- **Prometheus Integration**:
  - Fully compatible with Prometheus monitoring systems.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/repository-name.git
   cd repository-name
   ```
2. Install dependencies:

```
go mod tidy
Run the application:

bash
Копіювати
Редагувати
go run main.go
