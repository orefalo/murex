package docs

func init() {

	Definition["addheading"] = "# _murex_ Shell Docs\n\n## Command Reference: `addheading` \n\n> Adds headings to a table\n\n## Description\n\n`addheading` takes a list of parameters and adds them to the start of a table.\nWhere `prepend` is designed to work with arrays, `addheading` is designed to\nprepend to tables.\n\n## Usage\n\n    <stdin> -> addheading: value value value ... -> <stdout>\n\n## Examples\n\n    » tout: jsonl '[\"Bob\", 23, true]' -> addheading name age active                                                                                   \n    [\"name\",\"age\",\"active\"]\n    [\"Bob\",\"23\",\"true\"]\n\n## See Also\n\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [commands/`count`](../commands/count.md):\n  Count items in a map, list or array\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/`regexp`](../commands/regexp.md):\n  Regexp tools for arrays / lists of strings"

}