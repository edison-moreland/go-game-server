FROM python:3.7

WORKDIR /locust

RUN pip install locustio

RUN curl "https://raw.githubusercontent.com/edison-moreland/go-game-server/master/scripts/locustfile.py" > locustfile.py

EXPOSE 8089:8089
EXPOSE 5557
EXPOSE 5558

ENV SERVERADDRESS localhost
ENV SERVERPORT 8080

ENTRYPOINT locust
CMD ["--host=${SERVERADDRESS}:${SERVERPORT}"]
