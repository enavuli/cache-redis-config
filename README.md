# cache-redis-config
====================

## Description
------------

A configuration management tool for Redis cache, providing a simple and intuitive way to manage Redis connections, data expiration, and other configuration settings.

## Features
------------

*   **Multi-Environment Support**: Seamlessly switch between different Redis environments with ease.
*   **Configurable Connection Settings**: Easily modify Redis connection parameters, such as host, port, and password.
*   **Data Expiration Management**: Set TTL (time-to-live) for Redis keys to automatically expire data after a specified period.
*   **Key Prefix Management**: Add a prefix to all Redis keys to improve cache organization and prevent key collisions.
*   **Monitoring and Logging**: Integrate with popular monitoring tools for real-time performance insights and error logging.

## Technologies Used
--------------------

*   **Node.js**: The project is built using Node.js, utilizing its asynchronous and event-driven nature to handle Redis connections efficiently.
*   **Redis**: The primary cache storage engine, providing a fast and scalable solution for storing and retrieving data.
*   **Lodash**: Utilized for utility functions and data manipulation.

## Installation
------------

### Prerequisites

*   Node.js (version 14 or higher)
*   Redis (version 6 or higher)

### Step 1: Clone the repository

```bash
git clone https://github.com/your-username/cache-redis-config.git
```

### Step 2: Install dependencies

```bash
cd cache-redis-config
npm install
```

### Step 3: Configure Redis connection settings

Create a new file named `config.json` in the project root directory with the following structure:

```json
{
    "redis": {
        "host": "localhost",
        "port": 6379,
        "password": "your-redis-password"
    },
    "ttl": {
        "default": 3600 // default TTL in seconds
    },
    "keyPrefix": {
        "default": "cache:" // default key prefix
    }
}
```

### Step 4: Run the application

```bash
node app.js
```

## Usage
-----

The `cache-redis-config` tool provides a simple command-line interface for managing Redis connections, data expiration, and other configuration settings.

### Commands

*   `start`: Start the Redis connection and begin monitoring data expiration.
*   `stop`: Stop the Redis connection and terminate data expiration monitoring.
*   `config`: Display the current Redis connection settings and data expiration configuration.
*   `log`: Display the last 10 lines of Redis log output.

## Contributing
------------

Contributions are welcome! Please fork the repository, make your changes, and submit a pull request.

## License
-------

This project is licensed under the MIT License.