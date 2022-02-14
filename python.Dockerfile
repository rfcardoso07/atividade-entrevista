FROM python:3

WORKDIR /app

RUN pip3 install requests

COPY . .

ENTRYPOINT [ "python", "runner.py" ]