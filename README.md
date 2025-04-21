# 🧋 飲品 PK 折價券活動 — Side Project 規格書

## 🧠 專案簡述

這是一個多回合商品 PK 系統\
使用者對飲品進行一對一對戰投票\
最終選出一位「冠軍飲品」，並提供一張優惠券

---

## 🧱 系統服務
| Service Name     | 說明                                  |
|------------------|---------------------------------------|
| `front-end`      | 使用者操作界面，進行 PK 投票              |
| `broker`         | API Gateway，統一對外與各微服務串接       |
| `product-service`| 提供商品資料                            |
| `coupon-service` | 根據冠軍產出優惠券資訊                   |
| `score-service`  | 消化投票結果，計算分數寫入資料庫           |
| `log-service`    | 統一記錄 log                           |
| `report-service` | 產生 PK 活動報表                        |
| `mail-service`   | 發送報表到指定 Email                    |
---

## 🧾 流程說明

### 1. PK 流程
- 經由 broker 從 product-service 取得所有飲品
- 每次展示兩張飲品，點選勝者，替換敗者
- 記錄每場比賽結果
- 最後選出積分最高的飲品為冠軍

### 2. 提供優惠券
- 經過 broker，由 coupon-service 提供冠軍飲品優惠券

### 3. 結束投票
- 將所有回合結果與最終冠軍傳給 broker
- broker 將結果透過 Message Queue 傳給 score-service
- score-service 統計勝負資訊，寫入 DB

### 其他
- log-service 統一管理 log
- report-service 產生統計報表
- mail-service 寄送報表


---

## 🧱 建置流程