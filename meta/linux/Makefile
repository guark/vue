# to install the app run: sudo make install

install:
	install -C {{ .ID }} /usr/local/bin/{{ .ID }}
	install -C {{ .ID }}.desktop /usr/local/share/applications/{{ .ID }}.desktop
	mkdir -p /usr/share/{{ .ID }}
	cp -r statics/ /usr/share/{{ .ID }}
	find /usr/share/{{ .ID }}/ -type d -exec chmod 755 {} \;
	find /usr/share/{{ .ID }}/ -type f -exec chmod 644 {} \;
	chmod +x /usr/local/bin/{{ .ID }}
