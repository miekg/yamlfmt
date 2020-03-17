file=$(mktemp)
cat <<EOF > $file
repos:
- repo: local
  hooks:  
  - id: documentation_checker
    name: documentation_checker
    entry: bash ./scripts/pre-commit/docgen.sh
    language: system
-   repo: https://github.com/pre-commit/mirrors-yapf
    rev: v0.29.0
    hooks:
    - id:   yapf
EOF

gold=$(mktemp)
cat <<EOF > $gold
repos:
- repo: local
  hooks:
  - id: documentation_checker
    name: documentation_checker
    entry: bash ./scripts/pre-commit/docgen.sh
    language: system
- repo: https://github.com/pre-commit/mirrors-yapf
  rev: v0.29.0
  hooks:
  - id: yapf
EOF

go install
test $? -eq 0 || { echo "Failed to compile"; exit 1; }

go run yamlfmt.go -w $file
test $? -eq 0 || { echo "Failed to run with replace mode"; exit 1; }

cmp $file $gold || { echo "Unexpected output"; diff $file $gold; exit 1; }

out=$(mktemp)
cat $file | go run yamlfmt.go > $out
test $? -eq 0 || { echo "Failed to run reading stdin"; exit 1; }

cmp $out $gold || { echo "Unexpected output"; diff $out $gold; exit 1; }

file1=$(mktemp)
cp $file $file1
go run yamlfmt.go -w $file $file1 || { echo "Failed to replace multiple files"; exit 1; }

cmp $file1 $gold || { echo "Unexpected output from replacing multiple files"; diff $file1 $gold; exit 1; }
