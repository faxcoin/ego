# EGO
## *E*ducational *GO* Parser

## Rough structure
```
val a = 0;
a = 1;
val b = a + b;
```
- line is terminated with semicolon
- `val` signals a variable declaration (only ints allowed for now)
- `=` is an assignment
- `+` is an addition
- `a`, `b` are names (only alphanumerical)
