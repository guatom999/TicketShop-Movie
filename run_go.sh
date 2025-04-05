environments=("./env/dev/.env.movie" "./env/dev/.env.payment" "./env/dev/.env.inventory" "./env/dev/.env.ticket" "./env/dev/.env.customer")

# Loop through environments and run the command
for env in "${environments[@]}"; do
    echo "Running: go run main.go $env"
    go run main.go "$env" &
done

# Wait for all background jobs to complete
wait