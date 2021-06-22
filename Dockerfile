FROM alpine
ADD bj38-service /bj38-service
ENTRYPOINT [ "/bj38-service" ]
