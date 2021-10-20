[WIP] A simple implementation of a micro-service architecture, inspred by <https://microservices.io/patterns/data/database-per-service.html>

Services:
- Product
- Order
- Authentication
- Customer

Typical user scenario:
- Logs in to service -> Authentication
- View available products -> Product
- Store products in cart -> Order
- Order stored products -> Order, Customer
