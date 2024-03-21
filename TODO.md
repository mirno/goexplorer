# Project Todos

## Immediate

- [ ] Add detailed logging to both server and client
- [ ] Write basic unit tests for existing server and client code.
- [ ] Usage of env variables and flags for scalable and secure solutions

 
## Short-Term

- [ ] Harden TLS configuration for production
- [ ] 
- [ ] Implement the primary API routes on the server
- [ ] Start documenting API endpoints and setup instructions
- [ ] Write unit tests for HTTP handlers
- [ ] Implement linters and strategies and add them within an ADR
- [ ] tlsConfig.BuildNameToCertificate is depricated

## Long-Term

- [ ] Implement advanced authentication and authorization mechanisms (OAUTH2)
- [ ] Implement Lefthook with pre-commit and pre-push
- [ ] Optimize performance based on profiling results
- [ ] Dockerize (containerize) the server and client applications
- [ ] Set up a basic CI/CD pipeline for automated testing and deployment
- [ ] Explore integration with OpenAPI for automated API documentation.
- [ ] Research gRPC for potential future use.
- [ ] Build docs/contributing.md and include [Trunk Based Development](https://trunkbaseddevelopment.com/) and [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [ ] Create ADR and docs for [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)
- [ ] [Open Telemetry implementation](https://github.com/open-telemetry/opentelemetry-go)
- [ ] Investigate [websocket](https://thinhdanggroup.github.io/websocket-vs-http2/) implementation
## Done

- [x] Set up project repository.
- [x] Set up basic server and client with mTLS
- [x] Setup scripts for certificates and dependencies
- [x] Implement mTLS in server-client communication
- [x] Investigate Make on Windows issue.
- [x] Set up initial `docs/` folder structure.
- [x] Setup the docs/setup.md

## Docs & Reads

- [ ] Read and investigate [Team Topologies](https://teamtopologies.com/)
- [ ] [System design](https://bytebytego.com/courses/system-design-interview/foreword)
- [ ] [Diagram as code](https://www.youtube.com/watch?v=jCd6XfWLZsg)
- [ ] Logging, Tracing & Metrics [1](https://www.youtube.com/channel/UCZgt6AzoyjslHTC9dz0UoTw/community?lb=Ugkx6hFPFXZvllSwQ6UTJ7QYIiyzMyD9RTSS) [2](https://microsoft.github.io/code-with-engineering-playbook/observability/log-vs-metric-vs-trace/) [3](https://peter.bourgon.org/blog/2017/02/21/metrics-tracing-and-logging.html)

## References

- [C4 model](https://c4model.com/)
- [Clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [SOLID principes](https://www.freecodecamp.org/news/solid-principles-for-better-software-design/) ([wiki](https://en.wikipedia.org/wiki/SOLID))
- [Structurizr](https://structurizr.com/)
