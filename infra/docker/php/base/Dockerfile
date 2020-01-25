FROM php:7.4.1-apache

WORKDIR /var/www/html

RUN apt update -y

RUN apt install git -y

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
