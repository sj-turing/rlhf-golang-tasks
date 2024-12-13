# Run the application with pprof
pprof -prof cpu docker run -p 8080:8080 your_image

# Generate a CPU profile
go tool pprof http://localhost:8080/debug/pprof/profile
