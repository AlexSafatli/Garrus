grammar Dice;

options { tokenVocab=DiceLexer; }

notation : dice | number | add ;
add : mult (ADDOPERATOR mult)* ;
mult : operand (MULTOPERATOR operand)* ;
operand : dice | number | LPAREN notation RPAREN ;
dice : ADDOPERATOR? DIGIT? DSEPARATOR DIGIT ;
number : ADDOPERATOR? DIGIT ;