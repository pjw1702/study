FROM ubuntu:18.04
RUN apt-get update
RUN apt-get install -y apache2
ENV APACHE_RUN_USER www-data
ENV APACHE_RUN_GROUP www-data
ENV APACHE_LOG_DIR /var/log/apache2
ENV APACHE_PID_FILE /var/run/apache2/apache2.APACHE_PID_FILE
EXPOSE 80
CMD ["/usr/bin/apache2ctl", "-D", "FOREGROUND"]

FROM alpine:latest
WORKDIR /var/www/html
COPY --from=0 /var/www/html .
CMD ["/usr/bin/apache2ctl", "-D", "FOREGROUND"]