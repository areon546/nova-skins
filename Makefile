
run:
	cd goPageMaker && go build && go run .

hello:
	echo "Hello, World"

test:
	cd goPageMaker && go build && go test

git:
	git pull && git push