package ast

import (
  "testing"
)



func TestParsingDeclaration(t *testing.T) {
  test_code := `val a = 0;`

  tree := ParseText(test_code)

  t.Fatalf("Fuck me bois")

}
