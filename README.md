### Project setup & run 
**Generate GraphQL files**: Run the following command to generate the GraphQL files:

`go run github.com/99designs/gqlgen generate`

**Docker Compose**: In order to build and run the project via docker compose please run:

`docker-compose build` 

`docker-compose up` 

### Structure

**cmd**: This directory contains your application's main entry points. Each subdirectory here can represent a different executable or command-line tool for your application. For example, you might have a cmd/api directory for your API server and a cmd/cli directory for a command-line interface.

**config**: Store your configuration files here. These files might include environment-specific configuration settings, database configuration, and other application settings.

**graph**: This directory contains the gqlgen-generated GraphQL files.

**internal**: The internal directory is for packages that are specific to your project and should not be imported by other projects. It's a way to enforce encapsulation and limit access to certain parts of your codebase. So this is a directory meant for code that is specific to your application and not meant to be reused.

**migrations**: This directory contains DB migrations

**pkg** : The pkg directory is where you place packages (Go libraries) that are designed for reuse within your project but could potentially be used by other projects in the future. These packages should be relatively self-contained and independent of your application's business logic.

**schema**: This directory holds the GraphQL schema files.

**tools**: This folder is used to manage tool dependencies

### Clean Architecture
The Clean Architecture, proposed by Robert C. Martin, focuses on separating concerns into layers with strict dependencies among them. It consists of several concentric circles, with the innermost circle representing the core business logic or domain, surrounded by layers like Use Cases, Interface Adapters, and Frameworks & Drivers. In Clean Architecture, the emphasis is on maintaining independence of the inner layers from the outer layers, allowing for easier testing, maintenance.

The key point in Clean Architecture is that the inner layers (use cases/entities) should not depend on the outer layers (such as frameworks, databases, or UI). Instead, the dependencies should flow from the outer layers towards the inner layers, allowing for flexibility, testability, and easier maintenance.

**Some Notes**: In this project i was using DI (dependency injection) container for managing and handling dependencies between various components of an application.

**Terminology**: Some terminology may differ from team to team

- Domain / Model / Entity
- Use Cases / Services / Interactors
- Presenters / Controllers / Delivery
- Repository / Gateway

**Clean Flow**: The flow typically involves:

**1. Resolvers Depending on Frameworks**: In Clean Architecture for a GraphQL API, resolvers may depend on frameworks like gqlgen to handle incoming GraphQL queries and mutations. This dependency is isolated to the outer layers, ensuring the core business logic remains independent of framework-specific details.

**2. Resolvers Depending on Services/Use Cases**: After handling requests, resolvers delegate business logic to services or use cases. This maintains the separation of concerns, with resolvers acting as the entry point and passing control to the services layer to execute application-specific logic.

**3. Services/Use Cases Depending on Models/Domain**: In line with Clean Architecture, services or use cases depend on domain models, which encapsulate core business rules and data structures. This keeps the business logic at the core of the application, isolated from external concerns.

**4. Services/Use Cases Depending on Repository Interfaces**: By depending on repository interfaces rather than concrete implementations, services or use cases remain decoupled from specific data storage mechanisms. This allows flexibility in changing data sources without impacting business logic.

**GraphQL Use Cases**:

<img width="943" alt="Screenshot 2024-09-24 at 21 48 48" src="https://github.com/user-attachments/assets/3d711991-12d3-4c7f-bb00-fcc3fce123d8">


