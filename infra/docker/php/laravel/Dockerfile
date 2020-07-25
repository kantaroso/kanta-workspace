FROM kantaroso/kanta-workspace-php-base:latest

COPY ./app/laravel/ /var/www/html/
COPY ./infra/docker/php/laravel/config/php/.env /var/www/html/.env
COPY ./infra/docker/php/laravel/config/apache/000-default.conf /etc/apache2/sites-available/000-default.conf
