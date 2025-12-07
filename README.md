# Qreate

QR code management platform with scan tracking and analytics.

## Tech Stack

- **Frontend**: Next.js 15, React 19, Tailwind CSS 4
- **Backend**: Go, Gin
- **Database**: PostgreSQL
- **Cache**: Redis
- **Monorepo**: Turborepo, pnpm

## Project Structure

    qreate/
    ├── apps/
    │   ├── api/          # Go backend (port 8080)
    │   ├── backoffice/   # Admin dashboard (port 3000)
    │   └── webapp/       # QR redirect service (port 3001)
    ├── packages/
    │   ├── ui/           # Shared UI components
    │   ├── tsconfig/     # Shared TypeScript configs
    │   └── eslint-config/# Shared ESLint configs
    └── docker/           # Dockerfiles

## Getting Started

### Prerequisites

- Node.js 22+
- pnpm 10+
- Go 1.24+
- Docker (for PostgreSQL & Redis)

### Installation

    # Install dependencies
    pnpm install

    # Initialize Go modules
    cd apps/api && go mod tidy && cd ../..

    # Start infrastructure
    docker-compose up -d postgres redis

    # Build UI package
    pnpm build:ui

### Development

    # Run all apps
    pnpm dev

    # Or run individually
    pnpm dev:api        # Backend
    pnpm dev:backoffice # Admin dashboard
    pnpm dev:webapp     # Public webapp
    pnpm dev:ui         # UI package watcher

## Scripts

| Command              | Description                              |
|----------------------|------------------------------------------|
| `pnpm dev`           | Start all apps in dev mode               |
| `pnpm build`         | Build all apps                           |
| `pnpm lint`          | Run linting                              |
| `pnpm check-types`   | TypeScript type checking                 |
| `pnpm clean`         | Remove node_modules and build artifacts  |

## Branch Naming

- `feat/feature-name` - New features
- `fix/bug-name` - Bug fixes
- `enhancement/description` - Improvements
- `refactor/description` - Code refactoring
