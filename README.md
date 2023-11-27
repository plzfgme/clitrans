# clitrans

```text
A command-line translator.
It translate text from one language to another. When no text is provided, stdin will be used instead.

Currently supported translators: google.
Please refer to each translator's documentation for language codes used in 'from' and 'to' flags.

Currently supported preprocessors: remove_newlines (useful for text copied from a pdf).

Usage:
  clitrans [flags] [text]

Flags:
  -f, --from string             The language to translate from, or 'auto' (default "auto")
  -h, --help                    help for clitrans
  -p, --preprocessors strings   The preprocessors to use
  -t, --to string               The language to translate to
  -l, --translator string       The translator to use (default "google")
```
