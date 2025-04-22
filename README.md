# 🧋 Product Battle
- [🧠 Summary](#-summary)
- [🧱 System Services](#-system-services)
- [🧾 Flow](#-flow)
    - [Battle](#battle)
    - [Reward](#reward)
    - [Background](#background)
    - [Other](#other)
- [🧱 Build Process](#-build-process)

## 🧠 Summary
This is a multi-round product battle system.\
Users participate in one-on-one drink battles through voting.\
In the end, a "Champion Drink" is selected, and a coupon is provided.

---
## 🧱 System Services
| Service          | Description                                         |
|------------------|-----------------------------------------------------|
| `front-end`      | User interface for product voting                   |
| `broker`         | API Gateway that unifies microservice access        |
| `product-service`| Provides product data                               |
| `coupon-service` | Issues coupons based on voting results              |
| `score-service`  | Processes voting results and writes to the database |
| `bi-service`     | Provides voting stats                               |
---

## 🧾 Flow
### Battle
- Users access the page.
- Through the broker, all drinks are retrieved from the product-service.
- Two products are displayed at a time; users select the winner, replacing the loser.
- Each match result is recorded.
- The drink with the highest score is selected as the champion.

### Reward
- Through the broker, the coupon-service provides coupons for the champion drink.

### Background
- All round results are sent to the MQ via the broker.
- The score-service processes the results from the MQ and writes them to the database.

### Other
- Users access the page and retrieve statistics from the bi-service via the broker.
---

## 🧱 Build Process