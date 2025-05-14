# !/bin/bash

# start product-service
cd product-service
echo "Starting product-service..."
go run cmd/api/main.go &
cd ..

# start score-service
cd score-service
echo "Starting score-service..."
go run cmd/queue/main.go &
cd ..

# start widget-service
cd widget-service
echo "Starting widget-service..."
go run cmd/api/main.go &
cd ..

# start broker
cd broker
echo "Starting broker..."
go run cmd/api/main.go &
cd ..

# start frontend
cd frontend
echo "Starting frontend..."
go run cmd/web/main.go &
cd ..

# Wait for all background processes to finish
wait
echo "All services started."