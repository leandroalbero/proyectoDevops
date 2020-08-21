Pequeña aplicación escrita en Go que muestra en el navegador un saludo personalizado con el nombre del parámetro que se le pasa por la URL.
Se pide que en cada push al repositorio de Github se realicen tests, se muestre el coverage y se rechace el test en caso de que se detecte un problema de linting en el código fuente.
## Construir contenedor con Dockerfile:
    docker build . -t goserver:1.0
## Ejecutar el contenedor:
    docker run -p 80:80 -d --name goserver goserver:1.0
## Navegar a la URL:
    http://localhost/?key=leandro
Tras cada run de Github Actions se generan dos artifacts; uno con el binario compilado y otro con el resultado del análisis de cobertura.