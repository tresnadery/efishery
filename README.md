# Endpoint Documentation
You can see the Endpoint documentation in [Endpoint Documentation]file API.MD

# Manual Installation

## Requirement

* golang installed

* docker
see [docker documentation](https://docs.docker.com/get-docker/)

* postgres database

# Run Application

```docker-compose up```

Use sudo If your docker installed on root :

```sudo docker-compose up```

## Migration Database

```cat dump_21-07-2020_08_59_58.sql |docker exec -i efishery-db psql -U postgres```

Use sudo If your docker installed on root :

```cat dump_21-07-2020_08_59_58.sql |sudo docker exec -i efishery-db psql -U postgres```

# Production Architecture

The below diagram is my proposal deploying for production environment

[Diagram Syste](https://drive.google.com/file/d/1w3XitSQZN6oe_pEhqxPJIRrFDbvjmwT_/view?usp=sharing)



