FROM ubuntu:latest
RUN apt-get update && apt-get install -y nginx
RUN apt-get update && apt-get install -y \
  php \
  curl \
  php-cli \
  php-mbstring \
  git \
  unzip \
  php-fpm \
  php-zip \ 
  php-xml \
  php-mysqlnd \
  composer \
  php-pear \
  npm \
  && rm -rf /var/lib/apt/lists/*

