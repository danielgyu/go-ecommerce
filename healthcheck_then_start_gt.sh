#!/bin/sh

healthcheck() {
    while [ true ]
    do
	RES=$(curl -s $1/health | grep healthy)
	if [ -z "$RES" ]; then
	    echo "waiting for $1..."
	    sleep 3
	else
	    echo "$1 healthy"
	    break
	fi
    done
}

healthcheck product-service:8000
healthcheck user-service:8001
healthcheck order-service:8002

/dist/gateway
