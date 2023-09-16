
.PHONY: all grammar build test

ANTLR4_JAR=/usr/local/lib/antlr-4.13.1-complete.jar
ANTLR4=antlr4

GRAMMAR_NAME=Lecture
GRAMMAR_FILE=$(GRAMMAR_NAME).g4
PARSER_FOLDER=parser
GRAMMAR_FOLDER=lecture
GOLANG_PACKAGE=lecture

BINARY_NAME=lecture


GO_SOURCES = $(shell find . -type f -name '*.go') go.mod

all: grammar build

grammar: $(PARSER_FOLDER)

build: $(BINARY_NAME)

$(BINARY_NAME): $(PARSER_FOLDER) $(GO_SOURCES)
	go build -o $(BINARY_NAME) *.go

$(PARSER_FOLDER): $(PARSER_FOLDER)/java $(PARSER_FOLDER)/go

$(PARSER_FOLDER)/go: $(GRAMMAR_FILE)
	$(ANTLR4) -Dlanguage=Go -package $(GOLANG_PACKAGE) -o $(PARSER_FOLDER)/go/$(GRAMMAR_FOLDER) $(GRAMMAR_FILE)

$(PARSER_FOLDER)/java: $(GRAMMAR_FILE)
	$(ANTLR4) -o $(PARSER_FOLDER)/java/$(GRAMMAR_FOLDER) $(GRAMMAR_FILE)
	javac -cp "$(ANTLR4_JAR)" $(PARSER_FOLDER)/java/$(GRAMMAR_FOLDER)/$(GRAMMAR_NAME)*.java -d $(PARSER_FOLDER)/java/$(GRAMMAR_FOLDER)

test: $(GRAMMAR_FILE)
	go test -v -count=1 ./...
