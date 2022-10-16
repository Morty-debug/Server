# imagen base
FROM golang:1.12-alpine

# Agregamos git, bash y openssh a la imagen
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Agregamos informacion
LABEL maintainer="Ricardo Antonio Valladares Renderos <r_a_v_r_@hotmail.com>"

# Establecemos el directorio de trabajo en el contenedor
WORKDIR /docker

# Copiamos go.mod
COPY go.mod ./

# Descargamos todas las dependencias. Las dependencias se almacenaran en cache si los archivos go.mod no se modifican.
RUN go mod download

# Copie la fuente del directorio actual al directorio de trabajo dentro del contenedor
COPY . .

# Construimos la aplicacion
RUN go build -o main .

# Exponemos el puerto 8080
EXPOSE 8080

# Ejecutamos la aplicacion
CMD ["./main"]
