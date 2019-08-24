% YAMLFMT(1)
% Miek Gieben
% March 2019

# NAME

yamlfmt – format YAML

# SYNOPSIS

**yamlfmt** [**-strict**] [**-indent INDENT**] [FILE...]

# DESCRIPTION

**yamlfmt** formats YAML. If no files are given it will read from standard input, otherwise for each
file on the command line it will print formatted YAML to standard output.

Multiple YAML documents in a single input/file are supported, they are separated with the default
separator and an empty line.

# OPTIONS

**--help**
:   show a short help.

**-strict**
:   return an error on invalid YAML.

**-indent INDENT**
:   use **INDENT** as the indentation value, defaults to 2.
