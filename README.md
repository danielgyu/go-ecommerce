# A simple implementation of a micro-service architecture
inspred by 
- <https://www.cncf.io/blog/2021/07/19/think-grpc-when-you-are-architecting-modern-microservices/>
- <https://microservices.io/patterns/data/database-per-service.html>

# Services:
- Product
- Order
- User

# Typical user scenario:
- Logs in to service -> User 
- View available products -> Product
- Store products in cart -> Order
- Order stored products -> Order, Product, User

This project uses gRPC & protocol buffers for data communication. Each service is intended to work as a microservice with an gateway orchestrating in front of them.

//image

# Highlight
The intention of this project was to get a grasp of using grpc & protocol buffers, and implement them as microservices. The overall structure follows the API Gateway pattern, which knows which endpoint of the available microservices to call for a given request. Here's an example.

1. `/order/purchase` endpoint is called when the user wants to buy all the stored products in the user's cart
2. When the requst is made, the gateway service would first call `GetCredit` rpc call to the user service to get the user id and the user's current credit
3. Then it uses the user id to use the `OrderInCart` rpc call to get all the ids of the stored products of the users' cart
4. After that, it uses the product ids to get the aggregate sum of all the products by calling `GetProduct`
5. Finally, it compares the users' current credit(gotten from step 2) and the aggregate sum to determine whether the user can make the purchase

# How to Run
