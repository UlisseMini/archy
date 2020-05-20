FROM nginx:alpine

RUN apk add --no-cache ca-certificates

COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d /etc/nginx/conf.d

CMD ["nginx", "-g", "daemon off;"]
