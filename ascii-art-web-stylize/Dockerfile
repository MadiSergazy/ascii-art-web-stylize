from golang:1.17-alpine3.16

#MAINTAINER Madi856

WORKDIR /app

COPY . .
RUN apk add bash
#ENV PORT=4050
EXPOSE 4050 
#$PORT
# gdfg
CMD ["go", "run", "./Cmd"]