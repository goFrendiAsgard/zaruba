# Documentation

`ZtplAppDirectory` is a microservice-ready monolith application. It is built on top of [FastAPI](https://fastapi.tiangolo.com/), a modern, fast (high-performance) web framework for building APIs with Python 3.7+ based on standard Python-type hints

In this documentation, you will see `ZtplAppDirectory`'s [motivation and architecture](motivation-and-architecture/README.md) as well as how to manage and extend `ZtplAppDirectory`.

<!--startToc-->
- [Motivation and architecture](motivation-and-architecture/README.md)
  - [Directory structure](motivation-and-architecture/directory-structure.md)
  - [Feature flags](motivation-and-architecture/feature-flags.md)
  - [Interface and layers](motivation-and-architecture/interface-and-layers.md)
  - [Connecting components](motivation-and-architecture/connecting-components.md)
- [Conventions](conventions.md)
- [Adding a new module](adding-a-new-module/README.md)
  - [Adding a CRUD handler](adding-a-new-module/adding-a-crud-handler.md)
  - [Adding a new column](adding-a-new-module/adding-a-new-column.md)
  - [Adding an API endpoint](adding-a-new-module/adding-an-api-endpoint.md)
  - [Adding a page](adding-a-new-module/adding-a-page.md)
  - [Adding an event handler](adding-a-new-module/adding-an-event-handler.md)
  - [Adding an RPC handler](adding-a-new-module/adding-an-rpc-handler.md)
- [Authentication and Authorization](authentication-and-authorization.md)
- [User interface](user-interface/README.md)
  - [Menu](user-interface/menu.md)
  - [Jinja templates](user-interface/jinja-templates.md)
  - [Vue.js](user-interface/vuejs.md)
<!--endToc-->