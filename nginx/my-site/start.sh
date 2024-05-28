# 启动nginx
# 同时配置https ssl证书
# 证书地址根据需要改
docker run -d --network=host --rm  --name my-web\
    -v /home/dev/dev/qisimiaoxiang.site_nginx/qisimiaoxiang.site_bundle.crt:/etc/nginx/qisimiaoxiang.site_bundle.crt \
    -v /home/dev/dev/qisimiaoxiang.site_nginx/qisimiaoxiang.site.key:/etc/nginx/qisimiaoxiang.site.key \
    -v ./nginx/my-site/http.conf:/etc/nginx/conf.d/http.conf \
    -v ./nginx/my-site/https.conf:/etc/nginx/conf.d/https.conf \
    -v ./css-study/my_site/:/usr/share/nginx/html \
    -v ./nginx/my-site/log:/var/log/nginx nginx
