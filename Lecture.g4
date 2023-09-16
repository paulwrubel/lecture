grammar Lecture;

/**
 * FRAGMENTS
 */

// string parsing helpers
fragment NON_QUOTE: ~["];
fragment ESCAPED_QUOTE: '\\"';
fragment ALPHANUMERIC: [a-zA-Z0-9];
fragment ALPHA: [a-zA-Z];
fragment BASICSYMBOL: [\-@_.];
fragment NONQUOTEORESCAPED: (NON_QUOTE | ESCAPED_QUOTE);

// number parsing helpers
fragment DOT: '.';
fragment DIGIT: [0-9];

/**
 * RULES
 */

// initial rule
lecture: program EOF;

program: startClause (statement)+ endClause;

startClause: OKAY_HEAR_ME_OUT TERMINATOR;

endClause: I_REST_MY_CASE TERMINATOR;

// contructs

function: functionStatement (atomicStatement)* returnStatement;

// statements

statement: (atomicStatement | function);

atomicStatement: (declarationStatement | printStatement) TERMINATOR;

functionStatement:
	WE_CAN_USE_A_PROCESS_KNOWN_AS SPACE identifier (
		parametersDeclaration
	)? SPACE AND_PROCEEDS_AS_FOLLOWS TERMINATOR;

returnStatement: FINALLY_WE_GET SPACE valueClause TERMINATOR;

declarationStatement:
	LETS_SAY SPACE identifier SPACE IS SPACE valueClause;

printStatement: THEN_WE_HAVE SPACE valueClause;

// sub-statements

parametersDeclaration:
	COMMA SPACE WHICH_NEEDS SPACE parameterDeclarationClause COMMA;

parameterDeclarationClause:
	parameterDeclaration (
		SPACE AND SPACE parameterDeclarationClause
	)?;

parameterDeclaration:
	A SPACE type SPACE CALLED SPACE identifier;

// valueClause

valueClause: value (SPACE operator SPACE valueClause)?;

value: literalClause | identifier | functionCall;

literalClause: LITERALLY SPACE literal;

functionCall:
	THE_RESULT_OF SPACE identifier (
		SPACE WITH SPACE parametersClause
	)?;

parametersClause: value (SPACE AND SPACE parametersClause)?;

// simples

type: NUMBER;
operator: PLUS;

identifier: ALPHANUMERICSTRING;

literal: number;
number: INTEGER;

/**
 * TOKENS
 */

// reserved keyphrases
OKAY_HEAR_ME_OUT: 'okay, hear me out';
I_REST_MY_CASE: 'i rest my case';
LETS_SAY: 'let\'s say';
THEN_WE_HAVE: 'then we have';
WE_CAN_USE_A_PROCESS_KNOWN_AS: 'we can use a process known as';
WHICH_NEEDS: 'which needs';
AND_PROCEEDS_AS_FOLLOWS: 'and proceeds as follows';
FINALLY_WE_GET: 'finally, we get';
THE_RESULT_OF: 'the result of';

// reserved keywords
A: 'a';
CALLED: 'called';
IS: 'is';
AND: 'and';
WITH: 'with';
LITERALLY: 'literally';
NUMBER: 'number';
PLUS: 'plus';

// reserved symbols
COMMA: ',';
SPACE: ' ';
TERMINATOR: DOT;

// strings
ALPHANUMERICSTRING: ALPHA (ALPHANUMERIC)*;
// ALPHANUMERICSTRING: QUOTATION (ALPHANUMERIC | BASICSYMBOL)+ QUOTATION

// QUOTEESCAPEDSTRING: QUOTATION NONQUOTEORESCAPED* QUOTATION;

// numbers
INTEGER: DIGIT+;
// FLOATINGPOINT: DIGIT+ DOT DIGIT* | DOT DIGIT+;

// lazy whitespace
WS: [ \r\t\n]+ -> skip;