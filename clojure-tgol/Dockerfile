# Build stage
FROM clojure:lein AS builder
COPY . /code
WORKDIR /code
RUN lein uberjar

# Runtime stage
FROM openjdk:alpine
WORKDIR /app
COPY --from=builder /code/target/*-standalone.jar /app/app.jar
ENTRYPOINT [ "java", "-jar", "app.jar" ]
