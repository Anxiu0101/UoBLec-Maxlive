# Maxlive Micro Service

## How to Run the Stubs

Create a `.json` file for each service with the relevant addresses.
Use Dockerfile to build and deploy each service.
Example Docker build commands for each service:

```bash
# For API Gateway
docker build -t api-gateway .
docker run -p 8000:8000 api-gateway

# For Content Delivery Service
docker build -t content-delivery .
docker run -p 8001:8000 content-delivery

# For Content Creation Service
docker build -t content-creation .
docker run -p 8002:8000 content-creation
```

## Key Functionality

- **Echo Function**: Each service will have an endpoint (/ping) that echoes back a simple response.

- **Service Address Handling**: Each service reads its address from a services.json file and provides it via an endpoint (/services).