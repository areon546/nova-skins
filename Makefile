
run:
	cd goPageMaker && go build && go run .


test:
	cd goPageMaker && go test

hello:
	echo "Hello, World"

git:
	git pull && git push