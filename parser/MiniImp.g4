grammar MiniImp;

fragment Digit : '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' ;
fragment Letter : [A-Z] ;
fragment Escape : EscapedQuote | EscapedBackslash ;
fragment EscapedQuote : '\\"' ;
fragment EscapedBackslash : '\\\\' ;

Identifier : Letter ( Letter | Digit | '_' )* ;
Number : Digit Digit* ;
truth : 'true' | 'false' | 'not' truth | 'is' Identifier expr | truth 'or' expr | truth 'and' expr ;

expr   : read | simpleExpr | term (('+' | '-') term)* ;
simpleExpr   : asInt | asStr ;
term   : factor (('*' | '/') factor)* ;
factor : ('(' expr ')') | truth | Identifier | Number | String ;

stmt   : select | iterat | set | write ;
select : 'if' expr 'then' scope 'else' scope ;
iterat : 'while' expr scope ;
set    : 'set' Identifier '=' expr ';' ;
write  : 'write' expr ';' ;
read   : 'read' expr? ;

decl     : variable ;
variable : 'var' Identifier ( '=' expr )? ';' ;
asStr : 'as str' expr ;
asInt : 'as int' expr ;

stmts : stmt stmt* ;
decls : decl decl* ;
scope : 'begin' decls? stmts? 'end.' ;
init  : 'program' Identifier ;
prog  : init scope ;

String : '"' (Escape | ~["\\])*? '"' ;
WS : [ \t\n\r] + -> skip ;
