[WIP] A simple implementation of a micro-service architecture, inspred by <https://microservices.io/patterns/data/database-per-service.html>

Services:
- Product
- Order
- User

Typical user scenario:
- Logs in to service -> User 
- View available products -> Product
- Store products in cart -> Order
- Order stored products -> Order, User 

This project uses gRPC & protocol buffers for data communication.
