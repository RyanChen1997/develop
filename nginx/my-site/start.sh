docker run -d --network=host --rm  --name my-web\
    -v ./nginx/my-site/nginx.conf:/etc/nginx/conf.d/default.conf \
    -v ./css-study/my_site/:/usr/share/nginx/html \
    -v ./nginx/my-site/log:/var/log/nginx nginx
