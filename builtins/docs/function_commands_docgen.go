package docs

func init() {

	Definition["function"] = "# _murex_ Language Guide\n\n## Command Reference: `function`\n\n> Define a function block\n\n### Description\n\n`function` defines a block of code as a function\n\n### Usage\n\n    function: name { code-block }\n    \n    !function: command\n\n### Examples\n\n    » function hw { out \"Hello, World!\" }\n    » hw\n    Hello, World!\n    \n    » !function hw\n    » hw\n    exec: \"hw\": executable file not found in $PATH\n\n### Detail\n\n#### Allowed characters\n\nFunction names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n#### Undefining a function\n\nLike all other definable states in _murex_, you can delete a function with\nthe bang prefix (see the example above).\n\n#### Order of preference\n\nThere is an order of preference for which commands are looked up:\n1. Aliases - defined via `alias`. All aliases are global\n2. _murex_ functions - defined via `function`. All functions are global\n3. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n4. variables (dollar prefixed) - declared via `set` or `let`\n5. auto-globbing prefix: `@g`\n6. murex builtins\n7. external executable files \n\n### Synonyms\n\n* `function`\n* `!function`\n* `func`\n* `!func`\n\n\n### See Also\n\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathmatical function and assign to variable\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [source](../commands/source.md):\n  "

}