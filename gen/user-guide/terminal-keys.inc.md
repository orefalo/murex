# Terminal Hotkeys

## `tab`: autocomplete

Provides autocompletion suggestions. Press `esc` to hide suggestions.

## `ctrl`+`c`: kill foreground process

Pressing this will send a kill (SIGINT) request to the foreground process.

## `ctrl`+`d`: end of file

Send EOF (end of file). If the shell is sat on the prompt then this will exit
that running session.

## `ctrl`+`f`: search autocomplete suggestions

This will allow you to perform a regexp search through the autocompletion
suggestions. Thus allowing you to quickly navigate complex command options or
jump to specific sub-directories.

Press `esc` to cancel regexp search.

## `ctrl`+`r`: search shell history

This brings up your timestamped shell history as an autocomplete list with
regexp search activated. Using `ctrl`+`r` you can rapidly rerun previous
command lines.

Press `esc` to cancel history completion.

## `ctrl`+`u`: clear line

Clears the current line.

## `ctrl`+`\`: kill all running processes

This will kill all processes owned by the current _murex_ session. Including
any background processes too.

This function is a effectively an emergency kill switch to bring you back to
the command prompt.

Use sparingly because it doesn't allow processes to end graceful.

## `ctrl`+`z`: suspend foreground process

Suspends the execution of the current foreground process. You can then use job
control to resume execution in either the foreground or background. ([read more](../commands/fid-list.md))

## `esc` (aka "vim keys")

Pressing `esc` while no autocomplete suggestions are shown will switch the
line editor into **vim keys** mode.

Press `i` to return to normal editing mode.

### Supported keys

* `a`: insert after current character
* `A`: insert at end of line
* `b`: jump to beginning of word
* `B`: jump to previous whitespace
* `d`: delete mode
* `D`: delete characters
* `e`: jump to end of word
* `E`: jump to next whitespace
* `h`: previous character (like `🠔`)
* `i`: insert mode
* `I`: insert at beginning of line
* `l`: next character (like `🠖`)
* `p`: paste after
* `P`: paste before
* `r`: replace character (replace once)
* `R`: replace many characters
* `u`: undo
* `v`: visual editor (opens line in `$EDITOR`)
* `w`: jump to end of word
* `W`: jump to next whitespace
* `x`: delete character
* `y`: yank (copy line)
* `Y`: same as `y`
* `[`: jump to previous brace
* `]`: jump to next brace
* `$`: jump to end of line
* `%`: jump to either end of matching bracket
* `0` to `9`: repeat action _n_ times. eg `5x` would delete five (`5`) characters (`x`)

### Full Screen Editing via `$EDITOR`

When in "vim keys" mode, press `v` to bring up the visual editor. The editor
will be whichever command is stored in the `$EDITOR` environmental variable.