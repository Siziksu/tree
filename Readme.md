# Tree

`tree` displays the folder structure of a drive or path in a graphical format.

```text
Usage: tree [-d] [-e folder1,folder2,...] [depth]

Flags:
  -d    Show directories only
  -e string
        Comma-separated list of folders to exclude (e.g., .git,node_modules)
  -h    Show hidden files and folders
```

Examples

```powershell
# Default usage without flags and depth
tree

# Show hidden entries with a depth of 5
tree -h 5

# Excluding folders, showing only no hidden entries and using a depth of 3
tree -e .git,android,.idea,.dart_tool 3

# Excluding folders, showing only no hidden entries and without depth
tree -e .git,node_modules

# Just directories, including hidden entries
tree -d -h

# Export the result to a file
tree -d > list.txt
```
