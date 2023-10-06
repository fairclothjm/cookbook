# Cheatsheat ripgrep (rg)

| Flag    | Description |
| -------- | ------- |
| -I  | suppress file path |
| -N  | suppress line num |
| -l  | Only show file names    |
| --files-without-match | Only show non-matching file names     |
| -v    | show lines that do not match    |
| -c    | count matching lines |
| --count-matches    | count individual matches |
| -s    | case sensitive |
| -C NUM   | show NUM lines before and after |
| -F | treat pattern as fixed string |
| -x | match entire line only |
| --max-depth NUM | limit dir traversal depth |
| --no-ignore | don't respect ignore files |
| --no-messages | suppress all error messages |
| -o | print only matched parts |
| --passthru | cat file and highlight match |
| -q | quiet |
| -e PATTERN | regex pattern to search for |
| -r TEXT | replace matches with TEXT |
| --stats | print search stats |
| -t TYPE | only search files matching TYPE |
| --type-list | list available types |
| -T TYPE | don't search files matching TYPE |
| -w | only show matches surrounded by word boundaries |
