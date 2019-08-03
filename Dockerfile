# -------- BASE IMAGE ----------
FROM golang:1.12 AS base

# -------- BUILD -----------
FROM base AS builder

WORKDIR /ebox-api
COPY . /ebox-api
RUN make build

# -------- PRODUCTION -----------
FROM base AS production

COPY --from=builder /ebox-api/build/ebox-api /bin
ENV GIN_MODE=release
ENTRYPOINT /bin/ebox-api
