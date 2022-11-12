dev:
    docker-compose up

build-dev:
    docker build -t app-dev . --target development

build-prod:
    docker build -t app-prod . --target production

start:
    docker run -p 80:4000 --name app-prod app-prod 