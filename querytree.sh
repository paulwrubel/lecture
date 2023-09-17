#!/bin/bash

ANTLR4_JAR="/usr/local/lib/antlr-4.13.1-complete.jar"

GRAMMAR_NAME="LectureLang"
GRAMMAR_FOLDER="lecture"

ANTLR4="java -cp "$ANTLR4_JAR" org.antlr.v4.Tool"
ANTLR4_JAVA_GUI="java -classpath "parser/java/$GRAMMAR_FOLDER:$ANTLR4_JAR" org.antlr.v4.gui.TestRig"

echo $ANTLR4_JAVA_GUI

$ANTLR4_JAVA_GUI $GRAMMAR_NAME lecture -tokens -gui -tree <<< $(cat -)