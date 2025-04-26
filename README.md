# Go Backend Boilerplate
A very opinionated boilerplate for Go backend applications.

## Who should use this boilerplate
Basically you don't need HOC (Hexagonal, Onion, Clean) Architecture unless you have proven that your project is very complicated. By saying complicated, it means your project has a very complex business logic to the point that traditional `presentation-business-dal` (3 layer) architecture can't handle scaling your application (and of course your business)



## Architecture

![Architecture Simplified](https://oms-public-dev.s3.ap-southeast-1.amazonaws.com/hoc_simplified.png)

The project uses a mixture of different variants layered architecture with dependency inversion; primarily Clean Architecture is used as a baseline. However, some components may be different. Below is the summary

- The domain layer consists of DDD-oriented domain objects (Aggregate Roots, Child Entities, and Value Objects)
- Ubiquitous languages between aggregate roots are wrapped inside the `domain services`
- The term `usecase` is still used due to the explicit nature (instead of `application service` which is more ambiguous)
- Presenter does not exist in this architecture, choosing how to map the domain objects to dtos depend on the developers. You can simply map the domain objects to a dto or use additional mapping functions
- The `Gateway` represents the external interfaces for integrating with external systems. This could be other microservices (bounded contexts) or 3rd party services

## How to run the application
### Local
1. Install Go 1.22+
2. Make sure you have `.env` file in the root with the following environment variables set
3. Run the command to start an application
`go run cmd/api/main.go`

### Docker
1. Have Docker installed on your machine
2. Build the image
```bash
docker build -t <image_name> .
```



## Scenarios for the boilerplate

Order management system with the following behaviours

1. The user can place the order
2. The system will assign customerId and automatically generate GUID for the order (aggregate root) along with other fields e.g. address (value object), order items (child entities)
3. The discount is applied to the order item that has the matched id. In the real implementation, this is performed by the domain services with HTTP client connecting to `promotion` bounded context to retrieve the discount
4. The discount value is then applied to the order
5. The use case then maps the aggregate root to the DTO
6. The users can always query the order using known GUID




 
======= 

