
run:
	cd goPageMaker && go build && go run .


test:
	cd goPageMaker && go test -cover

coverage:
	cd goPageMaker && go test -cover && go tool cover -html=cover.out


hello:
	echo "Hello, World"

git:
	git pull && git push