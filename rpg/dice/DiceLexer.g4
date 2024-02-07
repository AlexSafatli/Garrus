lexer grammar DiceLexer;

DSEPARATOR : ( 'd' | 'D' ) ;
DIGIT : ('0'..'9')+ ;
ADDOPERATOR : ( ADD | SUB ) ;
MULTOPERATOR : ( MULT | DIV ) ;

LPAREN : '(' ;
RPAREN : ')' ;

fragment ADD : '+' ;
fragment SUB : '-' ;
fragment MULT : '*' ;
fragment DIV : '/' ;

WS : [\t\r\n]+ -> skip ;