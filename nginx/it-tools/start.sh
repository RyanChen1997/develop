docker run -d --name it-tools --restart unless-stopped -p 8080:443 \
 	-v /root/qisimiaoxiang.site_nginx/qisimiaoxiang.site_bundle.crt:/etc/nginx/qisimiaoxiang.site_bundle.crt \
 	-v /root/qisimiaoxiang.site_nginx/qisimiaoxiang.site.key:/etc/nginx/qisimiaoxiang.site.key \
 	-v ./nginx/it-tools/nginx.conf:/etc/nginx/conf.d/default.conf \
	-v ./nginx/my-site/log:/var/log/nginx \
 	corentinth/it-tools:latest

# docker run -d --name it-tools --restart unless-stopped -p 8081:80 \
# 	corentinth/it-tools:latest
