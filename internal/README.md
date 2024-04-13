# Internal Package

This is the heart of our appplication. All it's internal logic is stored here. the internal directory is not imported into other applications or libraries or modules. The code written here is for internal use within the codebase. This is were we store our business logic, i.e all logic associated with our app. This is going to contain three layers, the transport, business and database layers. The communication should be from top to bottom without having any skips in-between. The transport layer maybe the network layer of the app where our user interacts with. The business logic layer contains the logic that supports our app's core function. The database layer is responsible for interacting with our databases or other non-busineess related information processing.

## Structure

- /app: the point where all our dependencies and logic are collected and run the app. Typically our run function is found here which the main function calls.

- /config: Initializations of the app's configuration.

- /database: The files containing database/information storage related logic.

- /models: The structures for the database tables

- /services: The business logic files of our app.

- transport: The http-server settings, handlers, ports, etc.