<div align="center">
  <img src="https://skillicons.dev/icons?i=go,ps" alt="Hexa Code Tech Stack" />
  <h1 align="center">⚡ HEXA CODE</h1>
  <p align="center">
    <strong>Professional Enterprise-Grade Local AI Agent CLI.</strong>
  </p>

  <p align="center">
    <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go" alt="Language" />
    <img src="https://img.shields.io/badge/Backend-Ollama-white?style=for-the-badge&logo=ollama&logoColor=black" alt="Backend" />
    <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" alt="License" />
  </p>
</div>

---

## 🏛️ Project Architecture
**Hexa Code** is a high-performance Command Line Interface (CLI) designed for **local-first autonomous engineering**. Developed in **Go**, it provides a secure, zero-latency bridge between your local file system and Large Language Models (LLMs) via **Ollama**.

- **Privacy First:** Your code never leaves your local machine.
- **Enterprise UI:** Aesthetic terminal components via `lipgloss`.
- **Intelligent Prompting:** Specialized system prompt for software engineering.

---

## 📥 Getting Started

### System Prerequisites
*   **Go SDK:** v1.20+
*   **Ollama Engine:** Running on `localhost:11434`

### Build & Execution
```bash
# Resolve dependencies
go mod tidy

# Run the Agent
go run main.go