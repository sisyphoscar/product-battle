# Product Battle

中文 | [English](./README.md)

## 目錄
- [簡介](#簡介)
- [系統架構](#系統架構)
- [建置步驟](#建置步驟)
- [系統入口](#系統入口)

---

## 簡介

Product Battle 是一個商品對決系統，採用逐輪一對一淘汰賽機制，計算商品得分並視覺化各產品熱門度。\
本專案旨在熟悉 Microservices 架構，並實作於 Go、Gin、gRPC、RabbitMQ、PostgreSQL、Docker、Kubernetes 等技術。

---

## 系統架構

| Service            | 說明                       |
|--------------------|----------------------------|
| `frontend`         | 商品對決介面、數據儀表板   |
| `broker`           | 微服務 API gateway         |
| `product-service`  | 提供商品資料               |
| `score-service`    | 處理分數計算與儲存         |
| `widget-service`   | 提供儀表板數據             |

**架構圖：**

![架構圖](system.drawio.png)

---

## 建置步驟

### Docker containers

1. 執行建置指令
   ```bash
   bash deployment/docker/build.sh
   ```
2. 依需求修改 `.env`

---

### 繼續使用 Kubernetes

1. 下載並啟動 minikube
   ```bash
   minikube start --nodes=2
   ```
2. 啟用 ingress
   ```
   minikube addons enable ingress
   ```
3. 部署服務
   ```bash
   kubectl apply -f deployment/k8s/ --recursive
4. 設定 hosts
   ```bash
   sudo vi /etc/hosts
   ```
   在檔案末尾新增：
   ```
   127.0.0.1 product-battle.frontend product-battle.broker
   ```

5. 建立 localhost 與 Cluster 的路由通道
   ```bash
   minikube tunnel
   ```

---

## 系統入口

- Docker container: [http://localhost:8080](http://localhost:8080)
- Kubernetes: [http://product-battle.frontend/dashboard](http://product-battle.frontend/dashboard)
- RabbitMQ: [http://localhost:15672](http://localhost:15672)