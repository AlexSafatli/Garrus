grammar Dice;

D       : [dD] ;
SIGN    : [\-+] ;
LPAREN  : '(' ;
RPAREN  : ')' ;
COMMA   : ',' ;
SPACE   : ' ' ;

WS      : [\r\n\t] -> skip ;

Integer       : [0-9]+ ;
Id            : [a-zA-Z][A-Za-z0-9]+ ;
StringLiteral : '"' ~('\'' | '\\' | '\n' | '\r')+ '"' ;

notation  : count? sides modifier? ;
count     : Integer ;
sides     : Id ;
modifier  : SIGN Integer ;