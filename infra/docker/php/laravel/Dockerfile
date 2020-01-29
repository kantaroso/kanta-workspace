FROM php:7.4.1-apache

COPY ../../../../app/laravel/ /var/www/html/
COPY ./config/php/.env /var/www/html/.env
COPY ./config/apache/000-default.conf /etc/apache2/sites-available/000-default.conf
