# YAMLfmt

`yamlfmt` is a little Go binary that reads YAML from file(s) or standard input, formats it and then
writes it to standard output.

## Use in Editors

### Vim

Use the following configuration in Vim, to create a "Fmt" command that formats your YAML.

~~~ viml
au FileType yaml command! Fmt call YamlFmt(120)
let yaml_fmt = "yamlfmt /dev/stdin"
au FileType yaml let &l:formatprg=yaml_fmt
~~~
