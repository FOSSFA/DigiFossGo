BIN_NAME=bot

compile:
	go build -o ${BIN_NAME} ./cmd/DigiFoss/main.go

run:
	./${BIN_NAME} --config ./build/config/config-main.yml

build_and_run: compile run

clean:
	go clean
	rm ${BIN_NAME}