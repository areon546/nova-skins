run:
	make compile-md 
	make html

copy:
	cd www && make cp-local 
	cd www && make reload

md compile-md:
	cd goPageMaker && go build && go run .

test:
	cd goPageMaker && go test -cover

coverage:
	cd goPageMaker && go test -coverprofile cover.out && go tool cover -html=cover.out

get:
	cd goPageMaker && go get -u

git:
	git pull && git push

setup-hooks:
	git config core.hooksPath hooks


html compile-html:
	# copy relevant files over to contents

	# compile templates 
	./www/template-compiler -c www/content -t www/content -o www/output
