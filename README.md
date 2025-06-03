# Product Battle

[中文](./README-ZHTW.md) | English

## Table of Contents
- [Introduction](#introduction)
- [System Architecture](#system-architecture)
- [Build Steps](#build-steps)
- [System Entry Points](#system-entry-points)

---

## Introduction

Product Battle is a product showdown system using a round-based one-on-one knockout mechanism. It calculates product scores and visualizes product popularity.\
This project aims to provide hands-on experience with Microservices architecture, implemented with Go, Gin, gRPC, RabbitMQ, PostgreSQL, Docker, Kubernetes, and other technologies.

---

## System Architecture

| Service            | Description                        |
|--------------------|------------------------------------|
| `frontend`         | Product battle UI & dashboard      |
| `broker`           | Microservices API gateway          |
| `product-service`  | Provides product data              |
| `score-service`    | Handles score calculation/storage  |
| `widget-service`   | Provides dashboard data            |

**Architecture Diagram:**

![Architecture Diagram](system.drawio.png)

---

## Build Steps

### Docker containers

1. Run the build script:
   ```bash
   bash deployment/docker/build.sh
   ```
2. Modify `.env` as needed

---

### Next: Kubernetes

1. Download and start minikube:
   ```bash
   minikube start --nodes=2
   ```
2. Enable ingress:
   ```
   minikube addons enable ingress
   ```
3. Deploy services:
   ```bash
   kubectl apply -f deployment/k8s/ --recursive
   ```
4. Set up hosts:
   Edit your hosts file with:
   ```bash
   sudo vi /etc/hosts
   ```
   Add the following at the end:
   ```
   127.0.0.1 product-battle.frontend product-battle.broker
   ```

5. Create a route between localhost and the cluster:
   ```bash
   minikube tunnel
   ```

---

## System Entry Points

| Entry Point| Docker| Kubernetes|
|--------------|--------------------------------------------------------------------|----------------------------------------------|
| Game         | [http://localhost:8080](http://localhost:8080)                     | [http://product-battle.frontend](http://product-battle.frontend) |
| Dashboard    | [http://localhost:8080/dashboard](http://localhost:8080/dashboard) | [http://product-battle.frontend/dashboard](http://product-battle.frontend/dashboard) |
| RabbitMQ     | [http://localhost:15672](http://localhost:15672)                   | — |