grammar SwiftGrammar;

// tokens

// reserved words

VAR: 'var'; 
LET: 'let';
BREAK: 'break';
CONTINUE: 'continue';
RETURN: 'return';
REF: '&';
INOUT: 'inout';

// primitives
DECIMAL:    [0-9]+'.'[0-9]+;
ENTERO:     [0-9]+;
CADENA:     '"'('\\''"'| ~["])*'"';
CARACTER:   '\''(~[\\']|'\\t'|'\\n'|'\\\\'|'\\r'|'\\\'')'\''; //https://regex101.com/r/edu3zA/1
ID:         [a-zA-Z_][a-zA-Z0-9_]*;

// symbols
MAS:           '+';
MENOS:         '-';

// skip
COMENTARIO_L    : '//' ~[\r\n]* -> skip;//'//'.* -> skip;
COMENTARIO_M    : [/][*]~[*]*[*]+(~[/*]~[*]*[*]+)*[/] -> skip;
WS              : [ \r\n\t]+ -> skip;

// rules

prog: block EOF;

block: (instr ';'?)*;

instr: declaracion
    | asignacion
    | print_instr
    | if_instr
    | switch_instr
    | while_instr
    | for_instr
    | guard
    | BREAK
    | CONTINUE
    | RETURN expr?
    | funcion
    | llamada_func
    | dec_vector
    | copia_vector
    | modificacion_vector
    | append
    | removeLast
    | removeAt
    ;

declaracion
    : (VAR | LET) ID ':' tipo '=' expr # DeclaracionTipoValor
    | (VAR | LET) ID ':' tipo '?' # DeclaracionTipo
    | (VAR | LET) ID '=' expr # DeclaracionValor
    ;

tipo:   'Int' | 'Float' | 'Bool' | 'String' | 'Character' | '[' tipo ']';

asignacion
    : ID (MAS | MENOS)? '=' expr
    ;

print_instr
    : 'print' '(' ( expr (',' expr)*)? ')'
    ;

if_instr  
    : 'if' expr '{' block '}'                       # If
    | 'if' expr '{' block '}' 'else' '{' block '}'  # IfElse
	| 'if' expr '{' block '}' 'else' if_instr       # ElseIf
    ;

switch_instr
    : 'switch' expr '{' ( case | default )+ '}'
    ;

case
    : 'case' expr ':' block 
    ;

default
    : 'default' ':' block
    ;

while_instr
    : 'while' expr '{' block '}'
    ;

for_instr
    : 'for' ID 'in' ( expr | rango )'{' block '}'
    ;

rango
    : expr '...' expr
    ;

guard
    : 'guard' expr 'else' '{' block '}'
    ;

funcion
    //: 'func' ID '(' ((ID | '_' )? ID ':' 'inout'? tipo (',' (ID | '_' )? ID ':' 'inout'? tipo )*)? ')' ( '->' tipo )? '{' block '}'
    : 'func' ID '(' parametros? ')' ( '->' tipo )? '{' block '}'
    ;

parametros
    : parametros ',' ID? ID ':' INOUT? tipo 
    | ID? ID ':' INOUT? tipo
    ;

llamada_func
    //:  ( (ID ':')? REF? expr (',' (ID ':')? REF? expr)*)? ')'
    : ( ID '(' parametros_llamada? ')')
    ;

parametros_llamada
    : parametros_llamada ',' (ID ':')? REF? expr
    | (ID ':')? REF? expr
    ;

dec_vector
    : 'var' ID ':' '[' tipo ']' '=' '[' (expr (',' expr)*)? ']'
    ;

copia_vector
    : 'var' ID ':' '[' tipo ']' '=' expr 
    ;

modificacion_vector
    : ID '[' expr ']' '=' expr
    ;

append
    : ID '.append' '(' expr ')'
    ;

removeLast
    : ID '.removeLast' '(' ')'
    ;

removeAt
    : ID '.remove' '(' 'at' ':' expr ')'
    ;

expr
    : <assoc=right> '!' expr          # NotExpr
    | <assoc=right> '-' expr          # UmenosExpr
    | '(' expr ')'                          # ParExpr
    | left=expr op=('*'|'/'|'%') right=expr # OpExpr
    | left=expr op=('+'|'-') right=expr     # OpExpr
    | left=expr op=('>='|'>'|'<='|'<') right=expr    # OpExpr
    | left=expr op=('=='|'!=') right=expr   # OpExpr
    | left=expr op='&&' right=expr          # OpExpr
    | left=expr op='||' right=expr          # OpExpr
    | ENTERO                                # IntExpr
    | DECIMAL                               # FloatExpr
    | ID                                    # IdExpr
    | CARACTER                              # CharExpr  
    | CADENA                                # StrExpr  
    | 'nil'                                 # NilExpr
    | ('true' | 'false')                    # BoolExpr
    | 'Int' '(' expr ')'                    # IntCastExpr
    | 'Float' '(' expr ')'                  # FloatCastExpr
    | 'String' '(' expr ')'                 # StringCastExpr
    | llamada_func                          # LlamadaFuncExpr
    | ID '[' expr ']'                       # AccesoVector
    | ID '.isEmpty'                         # IsEmpty
    | ID '.count'                           # Count
    ;
