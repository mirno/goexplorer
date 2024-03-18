# Clean architecture layers

[layers](https://www.c-sharpcorner.com/article/clean-architecture-in-asp-net-core-web-api/) [1](https://positiwise.com/blog/clean-architecture-net-core)

* [Domain layer](#domain-layer)
* [Application layer](#application-layer)
* [Infrastructure layer](#infrastructure-layer)
* [Presentation layer](#presentation-layer)

## [Domain layer](https://dev.to/julianlasso/clean-architecture-domain-layer-3bdd)

The **Domain Layer** (sometimes called the **Business Logic Layer** or **Model Layer**) is at the core of your application. It encapsulates the business rules and knowledge. This layer is where you define entities, value objects, and domain events that represent concepts within the problem domain your application is addressing. The Domain Layer is kept pure, meaning it has no dependencies on other layers, such as the Application or Infrastructure layers. This purity and isolation make it easier to test and reason about.

**Entities**: Objects that have a distinct identity that runs through time and different states.
**Value Objects**: Objects that contain attributes but have no conceptual identity (e.g., Money, Color).
**Domain Services**: When an operation does not naturally belong to an Entity or Value Object, it can be defined in a service.

## [Application layer](https://dev.to/julianlasso/clean-architecture-application-layer-3n1e)

The **Application Layer** sits on top of the Domain Layer and orchestrates the flow of data to and from the Domain, directing the application's tasks and enforcing its rules. It's where the use cases of your application are implemented. This layer acts as a mediator between the outer layers (like the UI or external interfaces) and the Domain Layer. It's responsible for coordinating application activities, such as logging, transaction management, and security (often through application services).

**Use Cases**: The specific tasks your application can perform, represented as application services.
**DTOs (Data Transfer Objects)**: Objects used to transfer data between the Application Layer and outer layers without exposing domain entities.
**Interfaces/Ports**: Define the contracts that external services or layers must implement to be used by the application.

## Infrastructure layer

The **Infrastructure Layer** provides technical capabilities that support the higher layers. This includes things like database access, file system manipulation, network communication, and essentially any external agency or service required by the application. It implements interfaces defined by the Application Layer, ensuring that details about external libraries, frameworks, or systems are kept separate from the core business logic.

**Repositories**: Implementations of the repository interfaces defined in the Domain Layer, responsible for data access and persistence.
**API Clients**: Code that communicates with external services or APIs.
**Framework Code**: Specific implementations for web frameworks, ORMs, logging frameworks, etc.

## Presentation layer

The Presentation Layer is responsible for how the application's functionality is presented to the end-user or an external system. It's concerned with user interfaces (UI) but can also include adapters that expose your application to external systems, such as web APIs.

Key responsibilities include:

**User Interface (UI)**: Displays information to the user and interprets user commands. This can be a graphical user interface (GUI), a command-line interface (CLI), a web interface, or even a RESTful web service API that external clients consume.
**Input Validation**: Validates inputs before they are sent to deeper layers for processing. This ensures that only valid data is processed by the application and domain layers.
**View Models**: Constructs view models (or data transfer objects) that are optimized for specific views or API responses. These models are designed to shape data in a form that is most useful for the presentation layer, decoupling the UI from the domain model.
**Mapping**: Converts data between the application layer's DTOs and the view models used by the UI or external interfaces. This includes assembling and disassembling complex data structures that the UI components or external systems consume or produce.
**Session/State Management**: Manages the user's session and state, especially in web applications. This can include tracking user authentication, managing cookies, or maintaining state in stateless protocols.

---

### Clean Architecture Principles
The separation into these layers supports the **Open/Closed Principle** by allowing the application to be extended with new features or changes in external dependencies without modifying the existing business logic. It also aligns with the **Dependency Rule** of Clean Architecture, which states that dependencies can only point inwards, meaning higher-level layers can depend on lower-level layers but not the other way around.

#### Practical Example: A Cache Service
Imagine you're building a cache service as part of your application:

**Domain Layer**: Define a CacheItem entity with properties like Key, Value, and ExpirationTime. Include domain services for cache policies (e.g., eviction policies).
**Application Layer**: Implement use cases for adding, retrieving, and deleting cache items. Use DTOs to communicate cache item data.
**Infrastructure Layer**: Implement a repository for cache items, possibly using an in-memory store or an external database. If your cache needs to interact with external services (e.g., for distributed caching), this would also be part of the Infrastructure Layer.
By adhering to these architectural principles, you ensure that your application is well-organized, with each part focusing on its specific responsibilities. This separation makes your application more flexible, easier to test, and simpler to maintain or extend over time.