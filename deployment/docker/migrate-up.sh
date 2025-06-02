#!/bin/bash

# migrate to product database
docker run --rm -v $(pwd)/../../product-service/internal/infra/db/migrations:/migrations --network host migrate/migrate \
  -path=/migrations -database "postgres://root:root@localhost:54317/product?sslmode=disable" up

# migrate to score database
docker run --rm -v $(pwd)/../../score-service/internal/infra/db/migrations:/migrations --network host migrate/migrate \
  -path=/migrations -database "postgres://root:root@localhost:54317/score?sslmode=disable" up