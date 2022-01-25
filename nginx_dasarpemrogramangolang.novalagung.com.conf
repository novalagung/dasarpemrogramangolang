map $http_host $outbound {
    default allow;
    ~\s*live.rezync.com$ deny;
}

server {
    listen 80;
    listen [::]:80;

    server_name dasarpemrogramangolang.novalagung.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    listen [::]:443 ssl;

    server_name dasarpemrogramangolang.novalagung.com;

    ssl_certificate /etc/letsencrypt/live/dasarpemrogramangolang.novalagung.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/dasarpemrogramangolang.novalagung.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    include /etc/nginx/snippets/shared-security.conf;

    location ~ /.well-known {
        root /var/www/well-known;
    }

    include /etc/nginx/snippets/dasarpemrogramangolang-redirect.conf;

    location /images/ {
        access_log        off;
        log_not_found     off;
        expires           30d;
        add_header        Cache-Control "public";
        root              /var/www/dasarpemrogramangolang/_book;
    }

    location /gitbook/ {
        access_log        off;
        log_not_found     off;
        expires           30d;
        add_header        Cache-Control "public";
        root              /var/www/dasarpemrogramangolang/_book;
    }

    location /search_index.json {
        return 200;
    }

    location /dasarpemrogramangolang.pdf {
        add_header Content-Type "application/pdf";
        add_header Content-Disposition "inline; filename=\"Dasar Pemrograman Golang.pdf\"";

        root /var/www/dasarpemrogramangolang/_book;
        try_files $uri $uri/index.html;
    }

    location / {
        root /var/www/dasarpemrogramangolang/_book;
        try_files $uri $uri/index.html;
    }
}