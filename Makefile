
run:
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

build:
	make run 
	make html

html:
	# copy relevant files over to contents

	# compile templates 
	cd www && ./template-compiler -c content -t content -o output
