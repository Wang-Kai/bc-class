FROM redis

COPY ./app /app
COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT /entrypoint.sh