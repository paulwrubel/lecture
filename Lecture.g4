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

program: (function)* mainFunction (function)*;

// main function

mainFunction:
	mainStartStatement statementBlock mainEndStatement;

mainStartStatement: OKAY_HEAR_ME_OUT TERMINATOR;

mainEndStatement: I_REST_MY_CASE TERMINATOR;

// functions

function: functionSignature (statement)* returnStatement;

functionSignature:
	WE_CAN_USE_A_PROCESS_KNOWN_AS SPACE identifier (
		parametersDeclaration
	)? SPACE TO_PRODUCE_A SPACE type TERMINATOR IT_PROCEEDS_AS_FOLLOWS TERMINATOR;

parametersDeclaration:
	COMMA SPACE WHICH_NEEDS SPACE parameterDeclarationClause COMMA;

parameterDeclarationClause:
	parameterDeclaration (
		SPACE AND SPACE parameterDeclarationClause
	)?;

parameterDeclaration:
	A SPACE type SPACE CALLED SPACE identifier;

// statements

statement: (
		assignmentStatement
		| reassignmentStatement
		| printStatement
		| ifChainStatement
	) TERMINATOR;

statementBlock: statement+;

returnStatement: FINALLY_WE_GET SPACE valueClause TERMINATOR;

assignmentStatement:
	LETS_SAY SPACE identifier SPACE IS SPACE (
		valueClause
		| functionCall
	);

reassignmentStatement:
	NOW_LETS_SAY SPACE identifier SPACE IS SPACE (
		valueClause
		| functionCall
	);

printStatement: THEN_WE_HAVE SPACE valueClause;

ifChainStatement:
	ifStatement elseIfStatement* elseStatement? ifClosingStatement;

ifStatement: ifSignature statementBlock;

ifSignature:
	IF SPACE conditionClause COMMA SPACE HERES_WHAT_WE_NEED_TO_DO TERMINATOR;

elseIfStatement: elseIfSignature statementBlock;

elseIfSignature: OTHERWISE COMMA SPACE ifSignature;

elseStatement: elseSignature statementBlock;

elseSignature:
	OTHERWISE COMMA SPACE HERES_WHAT_WE_NEED_TO_DO TERMINATOR;

ifClosingStatement:
	NOW_THAT_WEVE_DONE_THAT COMMA SPACE WE_CAN_MOVE_ON;

conditionClause: valueClause SPACE comparator SPACE valueClause;

// valueClause

valueClause: value (SPACE operator SPACE valueClause)?;

value: literalClause | identifier;

literalClause: LITERALLY SPACE literal;

functionCall:
	THE_RESULT_OF SPACE identifier (
		COMMA? SPACE USING SPACE parametersClause
	)?;

parametersClause: parameter (SPACE AND SPACE parametersClause)?;

parameter: valueClause;

// simples

type: NUMBER;
operator: PLUS | MINUS;
comparator: IS;

identifier: ALPHANUMERICSTRING;

literal: number;
number: INTEGER;

/**
 * TOKENS
 */

// reserved keyphrases
OKAY_HEAR_ME_OUT: 'okay, hear me out';
I_REST_MY_CASE: 'i rest my case';
NOW_LETS_SAY: 'now let\'s say';
LETS_SAY: 'let\'s say';
THEN_WE_HAVE: 'then we have';
WE_CAN_USE_A_PROCESS_KNOWN_AS: 'we can use a process known as';
WHICH_NEEDS: 'which needs';
TO_PRODUCE_A: 'to produce a';
IT_PROCEEDS_AS_FOLLOWS: 'it proceeds as follows';
FINALLY_WE_GET: 'finally, we get';
THE_RESULT_OF: 'the result of';
HERES_WHAT_WE_NEED_TO_DO: 'here\'s what we need to do';
NOW_THAT_WEVE_DONE_THAT: 'now that we\'ve done that';
WE_CAN_MOVE_ON: 'we can move on';

// reserved keywords
A: 'a';
CALLED: 'called';
IS: 'is';
AND: 'and';
USING: 'using';
IF: 'if';
OTHERWISE: 'otherwise';
LITERALLY: 'literally';
NUMBER: 'number';
PLUS: 'plus';
MINUS: 'minus';

// reserved symbols
COMMA: ',';
SPACE: ' ';
TERMINATOR: DOT;

// strings
ALPHANUMERICSTRING: ALPHA (ALPHANUMERIC)*;
// ALPHANUMERICSTRING: QUOTATION (ALPHANUMERIC | BASICSYMBOL)+ QUOTATION
// 
// QUOTEESCAPEDSTRING: QUOTATION NONQUOTEORESCAPED* QUOTATION;

// numbers
INTEGER: DIGIT+;
// FLOATINGPOINT: DIGIT+ DOT DIGIT* | DOT DIGIT+;

// lazy whitespace
WS: [ \r\t\n]+ -> skip;