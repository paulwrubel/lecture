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

startClause: START_CLAUSE TERMINATOR;

endClause: END_CLAUSE TERMINATOR;

statement: (declarationStatement | printStatement) TERMINATOR;

declarationStatement:
	DECLARATION_STATEMENT_INTRO SPACE identifier SPACE IS SPACE valueClause;

printStatement: PRINT_STATEMENT_INTRO SPACE valueClause;

identifier: ALPHANUMERICSTRING;

valueClause: value (SPACE operator SPACE valueClause)?;
value: literalClause | identifier;
literalClause: LITERALLY SPACE literal;
literal: number;

operator: PLUS;

number: INTEGER;

/**
 * TOKENS
 */

// reserved keywords and phrases
START_CLAUSE: 'okay, hear me out';
END_CLAUSE: 'i rest my case';
DECLARATION_STATEMENT_INTRO: 'let\'s say';
PRINT_STATEMENT_INTRO: 'then we have';
TERMINATOR: DOT;

IS: 'is';
LITERALLY: 'literally';
PLUS: 'plus';
SPACE: ' ';

// reserved symbols
ALPHANUMERICSTRING: ALPHA (ALPHANUMERIC)*;

// strings ALPHANUMERICSTRING: QUOTATION (ALPHANUMERIC | BASICSYMBOL)+ QUOTATION;
// QUOTEESCAPEDSTRING: QUOTATION NONQUOTEORESCAPED* QUOTATION;

// numbers
INTEGER: DIGIT+;
// FLOATINGPOINT: DIGIT+ DOT DIGIT* | DOT DIGIT+;

// lazy whitespace
WS: [ \r\t\n]+ -> skip;