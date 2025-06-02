#!/bin/bash

# add products
docker run --rm -i --network host \
  -e PGPASSWORD=root \
  postgres:17 \
  psql -h localhost -p 54317 -U root -d product <<EOF
INSERT INTO products (name, description, image_url, price) VALUES
('Product 1', 'Description for Product 1', 'https://example.com/product1.jpg', 19.99),
('Product 2', 'Description for Product 2', 'https://example.com/product2.jpg', 29.99),
('Product 3', 'Description for Product 3', 'https://example.com/product3.jpg', 39.99),
('Product 4', 'Description for Product 4', 'https://example.com/product4.jpg', 49.99);
EOF

# add score_logs
docker run --rm -i --network host \
  -e PGPASSWORD=root \
  postgres:17 \
  psql -h localhost -p 54317 -U root -d score <<EOF
INSERT INTO public.score_logs (game, round, winner_id, loser_id) VALUES
('963c3620-1d40-4dcd-85b0-9f9a6c688b68', 1, 1, 2),
('963c3620-1d40-4dcd-85b0-9f9a6c688b68', 2, 3, 1),
('963c3620-1d40-4dcd-85b0-9f9a6c688b68', 3, 3, 4),
('b0bfe725-a67e-4c15-b44f-98e7dd5a47a7', 1, 1, 2),
('b0bfe725-a67e-4c15-b44f-98e7dd5a47a7', 2, 1, 3),
('b0bfe725-a67e-4c15-b44f-98e7dd5a47a7', 3, 1, 4);
EOF