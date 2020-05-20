FROM nginx:alpine

COPY data /var/www/data
COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d /etc/nginx/conf.d

CMD ["nginx", "-g", "daemon off;"]
