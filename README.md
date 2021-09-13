# EGO
## *E*ducational *GO* Parser
## Rough structure
```
val a = 0;
a = 1;
val b = a + b;

if (a) {
  print(a)
}
```
- line is terminated with semicolon
- `val` signals a variable declaration (only ints allowed for now)
- `=` is an assignment
- `+` is an addition
- `a`, `b` are names (only alphanumerical)
- `if (a)` checks if `a > 0`
- `print(a)` prints a's value

## EBNF Grammar
[What is EBNF?](https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form)

```
integer         = ? any integer ? ;
semicolon       = ";" ;
space           = " " | (" ", space) ;

binary_operator = "*" | "+" | "-" | "/" | "&" | "|";
expression      = ["("], (expression, binary_operator, expression, {binary_operator, expression}) | integer | name, [")"] ;

statement       = ("var" | "val", space, name, "=", expression, semicolon)
                | (name, "=", epxression, semicolon)
                | ("if", space, expression, "{", statement, "}", ["else", "{", statement, "}"])
                | ("while", space, epxression "{", statement "}") ;
```

## "features"
- strict left to right evaluation, you don't have to care about weird made up rules for numbers anymore
- no weird things like scopes, I never really liked them anyway, they just confused me
- functions are for lazy people that are afraid of "typing" too much, while me, as a real programmer, enjoy typing up to 2.000.000.000 words a day on my handmade mechanical keyboard

## rough flow
1) read line by line
2) use a recursive descent parser to build ast from text
3) profit
