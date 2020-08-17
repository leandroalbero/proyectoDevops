Construir contenedor con Dockerfile:
    docker build . -t goserver:1.0
Ejecutar el contenedor:
    docker run -p 80:80 -d --name goserver goserver:1.0

