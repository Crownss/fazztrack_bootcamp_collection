FROM golang:alpine

WORKDIR /week9

ENV DBUSER = ${USER}
ENV DBPASSWORD = ${PASSWORD}
ENV DBNAME = ${DBNAME}
ENV SSLMODE = ${SSLMODE}
ENV DBHOST = ${HOST}
ENV DBPORT = ${DBPORT}
ENV DBTIMEZONE = ${DBTIMEZONE}

ENV HOST = ${HOST}
ENV PORT = ${PORT}

COPY . .

RUN "go mod tidy && go build -o start"

ENTRYPOINT [ "/week9/start" ]