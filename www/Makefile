WEB-CONF=nova-skins

reload:
	make stop

link:
	ln -s /etc/nginx/sites-available/$(WEB-CONF) /etc/nginx/sites-enabled/$(WEB-CONF)

cp-local:
	cp -f ./localConfig /etc/nginx/sites-available/$(WEB-CONF)

cp-prod:
	cp -f ./prodConfig /etc/nginx/sites-available/$(WEB-CONF)

stop:
	sudo nginx -s reload
