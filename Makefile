.PHONY: gen-swagger-html
gen-swagger-html:
	redoc-cli bundle -o public/index.html docs/swagger.json