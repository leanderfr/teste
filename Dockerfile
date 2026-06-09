
# Use the official lightweight Alpine-based Redis image
FROM redis:7.2-alpine

# Set the working directory inside the container
WORKDIR /data

# Copy your custom Redis configuration file into the container
COPY redis.conf /usr/local/etc/redis/redis.conf

# Expose the default Redis port
EXPOSE 6379

# Start Redis server using the custom configuration file
CMD [ "redis-server", "/usr/local/etc/redis/redis.conf" ]



#  docker build -t my-redis .
#  docker run -d --name redis-app -p 6379:6379 -v ~/redis_data:/data --restart always my-redis
