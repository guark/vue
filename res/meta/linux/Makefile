# to install the app run: sudo make install

install:
	install -C bin/{{ .ID }} /usr/local/bin/{{ .ID }}
	mkdir -p /usr/share/{{ .ID }}
	cp -r datadir/ /usr/share/{{ .ID }}
	find /usr/share/{{ .ID }}/ -type d -exec chmod 755 {} \;
	find /usr/share/{{ .ID }}/ -type f -exec chmod 644 {} \;
	chmod +x /usr/local/bin/{{ .ID }}