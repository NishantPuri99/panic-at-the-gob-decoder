FROM golang

ARG FLAG
RUN apt-get update && apt-get install -y python3

COPY setup-challenge.py /app/

WORKDIR /app/website
COPY website/ .
RUN go build main.go

RUN mkdir /challenge && chmod 700 /challenge
RUN python3 ../setup-challenge.py

# You can include additional files at the end
# This is what users download to help solve the problem
RUN tar zcvf /challenge/artifacts.tar.gz main.go

EXPOSE 8000

# PUBLISH 8000 AS port

CMD ["./main"]