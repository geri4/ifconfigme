# Ifconfigme

This is simple go web service that return your ip address.
Web service listening port 9090.

## Installation

Reside this service behind nginx balancer. Set proxy_pass directive, also add
X-Real-Ip header. Example of nginx location:
```
location / {
    proxy_pass http://127.0.0.1:9090;
    proxy_http_version 1.1;
    proxy_set_header Connection "";
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header Host $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
}
```
