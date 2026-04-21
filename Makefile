test:
	go test $(if $(PKG),$(PKG),./...) $(if $(RUN),-run $(RUN),)
