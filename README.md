Simple Backend with Prometheus Client
=====================================

This repository provides a simple Go backend application integrated with the Prometheus client library to expose standard application metrics. Additionally, it includes a client application to interact with the backend and simulate resource loads such as CPU usage, goroutines, and memory allocations.

Features
--------

*   **Prometheus Metrics**: Exposes runtime metrics such as:
    *   CPU usage
    *   Number of goroutines
    *   Heap memory allocations
*   **Simulated Resource Usage**:
    *   The backend exposes a handler (`/allocate`) to emulate CPU load, allocate memory, and increase goroutines.
    *   The client application allows you to test various scenarios dynamically.
*   **Prometheus and Grafana Integration**:
    *   Monitoring metrics in real-time with Prometheus.
    *   Visualizing metrics with Grafana dashboards.

Project Structure
-----------------

    .
    ├── docker-compose.yml      # Docker Compose setup for Prometheus and Grafana
    ├── prometheus.yml          # Prometheus configuration
    ├── server/                 # Backend server code
    │   └── main.go             # Main entry point for the backend
    ├── client/                 # Client application code
    │   └── main.go             # Main entry point for the client
    └── README.md               # Project documentation


Installation
------------

1.  Clone the repository:

        git clone git@github.com:okharch/go-prometheus.git
        cd go-prometheus


2.  Ensure Docker and Docker Compose are installed. Start Prometheus and Grafana using:

        docker compose up -d


    This will:
    *   Start Prometheus on [http://localhost:9090](http://localhost:9090).
    *   Start Grafana on [http://localhost:3000](http://localhost:3000).

Backend Usage
-------------

### Expose Metrics

The backend exposes Prometheus metrics at:

    http://localhost:8080/metrics


### Simulate Resource Usage

The `/allocate` handler allows you to simulate resource usage. Example query parameters:

*   `size`: Allocate memory (in bytes).
*   `duration`: Simulate CPU load for a specified number of seconds.

Example:

    http://localhost:8080/allocate?size=1048576&duration=5


Client Usage
------------

1.  Navigate to the `client` directory:

        cd client


2.  Run the client with default parameters:

        go run main.go


3.  Customize parameters to test different scenarios:
    *   Allocate 10 MB of memory for 10 seconds:

            go run main.go -size=10485760 -duration=10


    *   Allocate 100 MB of memory for 30 seconds:
        
            go run main.go -size=104857600 -duration=30
            
        
    *   Specify a custom backend URL:
        
            go run main.go -url=http://your-backend-url:8080/allocate -size=2097152 -duration=15



Setting Up Prometheus and Grafana
---------------------------------

1.  Ensure the following `docker-compose.yml` file is present in the root directory:

        version: "3.8"
        
        services:
          prometheus:
            image: prom/prometheus:latest
            container_name: prometheus
            volumes:
              - ./prometheus.yml:/etc/prometheus/prometheus.yml
            ports:
              - "9090:9090"
            restart: always
        
          grafana:
            image: grafana/grafana:latest
            container_name: grafana
            ports:
              - "3000:3000"
            restart: always


2.  Add the following `prometheus.yml` file to the root directory:

        scrape_configs:
          - job_name: "go-backend"
            static_configs:
              - targets: ["host.docker.internal:8080"]


3.  Start the services:

        docker compose up -d


    Access Prometheus at [http://localhost:9090](http://localhost:9090) and Grafana at [http://localhost:3000](http://localhost:3000).

4.  Configure Grafana:
    *   Login with default credentials (`admin`/`admin`).
    *   Add Prometheus as a data source with URL: `http://prometheus:9090`.
    *   Create a dashboard and add panels for metrics like:
        *   `app_cpu_usage`
        *   `app_memory_usage`
        *   `go_goroutines`

Stopping Services
-----------------

To stop Prometheus and Grafana, run:

    docker compose down


Contribution
------------

Contributions are welcome! Feel free to open issues or submit pull requests to enhance the functionality or fix bugs.

License
-------

This project is licensed under the MIT License. See the `LICENSE` file for details.

Acknowledgments
---------------

*   [Prometheus Client Golang Library](https://github.com/prometheus/client_golang)
*   [Grafana](https://grafana.com)