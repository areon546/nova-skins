
run:
	cd goPageMaker && go build && go run .

test:
	cd goPageMaker && go test -cover

coverage:
	cd goPageMaker && go test -coverprofile cover.out && go tool cover -html=cover.out

get:
	cd goPageMaker && go get -u

hello:
	echo "Hello, World"
git:
	git pull && git push

setup-hooks:
	git config core.hooksPath hooks
