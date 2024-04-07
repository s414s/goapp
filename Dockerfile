# Cogemos imagen inicial
FROM golang:1.19-alpine
# Dentro del container creamos un directorio donde meter nuestro código
RUN mkdir /app
# Añadimos el código fuente a la carpeta
ADD . /app
# Establecemos el diretorio de trabajo
WORKDIR /app
# Compilamos el código fuente para obtener el binario
RUN go build -o main .
# Añadimos el puerto que vamos a exponer
EXPOSE 8080
# Especificamos el comando de consola para ejecutar el binario
CMD ["/app/main"]