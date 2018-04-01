FROM iron/go
WORKDIR /app
ADD alpine_scoreservice /app/
ENTRYPOINT ["./alpine_scoreservice"]