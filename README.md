# A simple implementation of a micro-service architecture
Inspred and ideas taken from the following articles.
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

![architecture](./components.jpeg?raw=true "architecture")

# Highlight
The intention of this project was to get an understading of using grpc & protocol buffers, and implement them in a microservice architecture. The overall structure follows the API Gateway pattern, which knows which endpoint of the available microservices to call for a given request. Here's an example.

1. `/order/purchase/` endpoint is called when the user wants to buy all the stored products in the user's cart
2. When the requst is made, the gateway service would first call `GetCredit` rpc call to the user service to get the user id and the user's current credit
3. Then it uses the user id to use the `OrderInCart` rpc call to get all the ids of the stored products of the users' cart
4. After that, it uses the product ids to get the aggregate sum of all the products by calling `GetProduct`
5. Finally, it compares the users' current credit(gotten from step 2) and the aggregate sum to determine whether the user can make the purchase

# How to Run & usage scenario
- `docker-compose build` then `up` to bring up service containers
- send `localhost:8080/initdb` request to initialize database
- signup `localhost:8080/signup/` include username and password in post request
- login `localhost:8080/user/login/` same as signup, and remember the token
- see products `localhost:8080/product/all/`, remember product ids in order to purchase them
- add to cart `localhost:8080/order/new/`, include token and productIds
- purchase items in cart `localhost:8080/order/purchase`, include token, the resopnse will tell if you have enough credit to buy all the stored products
- a sample walkthrough can be imported by using postman's collection json file
