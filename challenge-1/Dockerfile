FROM golang

COPY flag.txt /app/
WORKDIR /app/website
COPY website/ .
RUN go build main.go

RUN mkdir /challenge && chmod 700 /challenge

# You can include additional files at the end
# This is what users download to help solve the problem
RUN tar zcvf /challenge/artifacts.tar.gz main.go

# This needs to be your flag. I pulled this from your code. double check.
RUN echo "{\"flag\":\"$(cat flag.txt)\"}" > /challenge/metadata.json

EXPOSE 8000

# PUBLISH 8000 AS port

CMD ["./main"]
