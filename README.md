# Simple Backend with Prometheus Client

This repository provides a simple Go backend application integrated with the Prometheus client library to expose standard application metrics. Additionally, it includes a handler that emulates resource loads such as CPU usage, goroutines, and memory allocations.

## Features

*   **Prometheus Metrics**: Exposes runtime metrics like:
    *   CPU usage
    *   Number of goroutines
    *   Heap memory allocations
*   **Simulated Resource Usage**:
    *   A handler (`/allocate`) to emulate CPU load, allocate memory, and increase goroutines.
    *   Useful for testing how metrics change under different application loads.
*   **Prometheus Integration**: Fully compatible with Prometheus monitoring systems.

## Installation

1.  Clone the repository:
 ```   
    git clone git@github.com:okharch/go-prometheus.git
    cd go-prometheus
```
    
2.  Install dependencies:
```    
    go mod tidy
```                
    
3.  Run the application:
```    
    go run main.go
```
## Usage

### Expose Metrics

The application exposes Prometheus metrics at:

```
http://localhost:8080/metrics
```    

### Simulate Resource Usage

The `/allocate` handler allows you to simulate resource usage. Example query parameters:

*   `size`: Allocate memory (in bytes).
*   `duration`: Simulate CPU load for a specified number of seconds.

Example:
```
http://localhost:8080/allocate?size=1048576&duration=5
```
## Setting Up Prometheus

1.  Download and run Prometheus:
``` 
    wget https://github.com/prometheus/prometheus/releases/latest/download/prometheus-\*.tar.gz
    tar -xvzf prometheus-\*.tar.gz
    cd prometheus-\*
    ./prometheus --config.file=prometheus.yml
```
    
2.  Add the following scrape configuration to `prometheus.yml`:

```    
    scrape\_configs:
      - job\_name: "go-backend"
        static\_configs:
          - targets: \["localhost:8080"\]
```

3.  Access the Prometheus UI:
```    
    http://localhost:9090
```

## Setting Up Grafana (Optional)

1.  Install Grafana and run the server:
```
    grafana-server
```                
    
1.  Access Grafana at:
    
```
    http://localhost:3000
```

2.  Add Prometheus as a data source and create dashboards for:
    *   CPU usage
    *   Goroutines
    *   Heap memory allocations

## Contribution

Contributions are welcome! Feel free to open issues or submit pull requests to enhance the functionality or fix bugs.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments

*   [Prometheus Client Golang Library](https://github.com/prometheus/client_golang)
*   [Grafana](https://grafana.com)

Happy Monitoring! 🚀
