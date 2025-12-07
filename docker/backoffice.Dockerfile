# syntax=docker/dockerfile:1

# Base with pnpm
FROM node:20-alpine AS base
WORKDIR /app
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV NEXT_TELEMETRY_DISABLED=1
RUN corepack enable

# Dependencies layer (workspace-aware)
FROM base AS deps
# Use the repo workspace and lockfile for reproducible installs
COPY pnpm-workspace.yaml pnpm-lock.yaml package.json ./
# Only copy manifests for target packages to maximize cache hits
COPY apps/backoffice/package.json ./apps/backoffice/package.json
COPY packages/ui/package.json ./packages/ui/package.json
COPY packages/tsconfig/package.json ./packages/tsconfig/package.json
COPY packages/eslint-config/package.json ./packages/eslint-config/package.json
RUN pnpm install --frozen-lockfile

# Build layer
FROM base AS builder

# Accept API URL at build time for client-side bundling
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}
# Accept app environment at build time
ARG NEXT_PUBLIC_APP_ENV
ENV NEXT_PUBLIC_APP_ENV=${NEXT_PUBLIC_APP_ENV}

# Reuse installed deps (root, backoffice, and ui) so package-level bins are present
COPY --from=deps /app/node_modules ./node_modules
COPY --from=deps /app/apps/backoffice/node_modules ./apps/backoffice/node_modules
COPY --from=deps /app/packages/ui/node_modules ./packages/ui/node_modules
# Copy sources needed to build (include tsconfig so workspace symlinks resolve)
COPY packages/tsconfig ./packages/tsconfig
COPY packages/ui ./packages/ui
COPY apps/backoffice ./apps/backoffice
# Build UI first (generates CSS), then the backoffice
RUN pnpm --filter ui build
RUN pnpm --filter backoffice build

# Runtime layer
FROM node:20-alpine AS runner
WORKDIR /app/apps/backoffice
# ENV NODE_ENV=production
# ENV NEXT_TELEMETRY_DISABLED=1

# Set the API URL at runtime
ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=${NEXT_PUBLIC_API_URL}
# Set the app environment at runtime
ARG NEXT_PUBLIC_APP_ENV
ENV NEXT_PUBLIC_APP_ENV=${NEXT_PUBLIC_APP_ENV}

RUN corepack enable

# Copy runtime artifacts
COPY --from=builder /app/apps/backoffice/.next ./.next
COPY --from=builder /app/apps/backoffice/public ./public
COPY --from=builder /app/apps/backoffice/package.json ./package.json
COPY --from=builder /app/apps/backoffice/next.config.ts ./next.config.ts
# Bring node_modules for runtime resolution
COPY --from=deps /app/node_modules /app/node_modules
COPY --from=deps /app/apps/backoffice/node_modules ./node_modules
# Optional: keep UI package available if referenced at runtime/transpile
COPY --from=builder /app/packages/ui /app/packages/ui

EXPOSE 3000
ENV PORT=3000
CMD ["pnpm", "start"]