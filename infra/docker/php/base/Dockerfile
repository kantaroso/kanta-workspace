FROM php:7.4.1-apache

WORKDIR /var/www/html

RUN docker-php-ext-install pdo_mysql

RUN apt update -y

RUN apt install -y git

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
