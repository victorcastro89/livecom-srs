#!/bin/sh

# Remove existing containers if they are running
docker rm -f redis srs 2>/dev/null

# Run redis container
docker run --name redis --rm -it -v $HOME/data/redis:/data -p 6379:6379 -d redis

# Create config files if they don't exist and add a comment if they are empty
mkdir -p platform/containers/data/config
touch platform/containers/data/config/srs.server.conf platform/containers/data/config/srs.vhost.conf
echo "TODO: FIXME: Remove it after SRS supports empty config file."

if [ ! -s platform/containers/data/config/srs.server.conf ]; then
    echo '# OK' > platform/containers/data/config/srs.server.conf
fi

if [ ! -s platform/containers/data/config/srs.vhost.conf ]; then
    echo '# OK' > platform/containers/data/config/srs.vhost.conf
fi

# Run srs container
docker run --name srs --rm -it \
    -v $(pwd)/platform/containers/data/config:/usr/local/srs/containers/data/config \
    -v $(pwd)/platform/containers/conf/srs.release-ubuntu-change-your-ip.conf:/usr/local/srs/conf/srs.conf \
    -v $(pwd)/platform/containers/objs/nginx:/usr/local/srs/objs/nginx \
    -p 1935:1935 -p 1985:1985 -p 8080:8080 -p 8000:8000/udp -p 10080:10080/udp \
    --env CANDIDATE=$(ip addr show eth0 | grep 'inet ' | awk '{print $2}' | cut -d/ -f1) \
 ossrs/srs:5
