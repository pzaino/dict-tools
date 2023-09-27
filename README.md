# DICT-TOOLS

This is a collection of tools for working with unicode text files.

## Tools

### extract-alphabet

This tool extracts the alphabet from a text file. It is useful for
creating a list of characters to use in a dictionary.

It is also useful to quickly identify which characters have been used
in a text file. This helps determining which type of encryption was
used. For example, if the text file contains only characters from the
English alphabet, then it is likely that the text was encrypted using
a substitution cipher.

### words-occur

This tool counts the number of occurrences of each word in a text
file. It is useful for creating a list of words to use in a
dictionary.

It is also useful to determine which words are used most often in a
text file. This helps determining which type of encryption was
used. For example, if the text file contains only the word "the", then
it is likely that the text was encrypted using a substitution cipher.

It is also useful to determine the incidence of used instructions in a
program (for example in an assembly program).

### substr-occur

This tool counts the number of occurrences of each substring in a text
file. It is useful for creating a list of substrings and their number
of occurrences.

## Installation

```bash
git clone <this repo>

cd dict-tools

cd <tool-dir>

go build
```
