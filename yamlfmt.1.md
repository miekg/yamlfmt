% YAMLFMT(1)
% Miek Gieben
% March 2019

# NAME

yamlfmt – format YAML

# SYNOPSIS

**yamlfmt** [**-strict**] [FILE...]

# DESCRIPTION

**yamlfmt** formats YAML. If no files are given it will read from standard input, otherwise for each file on
the command line it will print formatted YAML to standard output.

Multiple YAML documents in a single input/file are supported.

# OPTIONS

**--help**
:   show a short help

**-strict**
:   return an error on invalid YAML.
