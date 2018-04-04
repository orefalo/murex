# _murex_ Language Guide

## Command reference: murex-docs

> Displays the man pages for _murex_ builtins

### Description

Displays the man pages for _murex_ builtins.

### Usage

    murex-docs: [ flag ] command -> <stdout>

### Examples

    # Output this man page
    murex-docs: murex-docs

### Flags

* `--digest`: returns an abridged description of the command rather than the
    entire help page.

### Detail

These man pages are compiled into the _murex_ executable.
