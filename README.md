# Batch file renamer

**Pietro Mascolo - iz4vve (at) gmail.com**

Batch rename files with option to remove the originals.

Usage:
```bash
$ rename -glob=<golb-of-files-to-rename> [-name=<new-files-name-prefix> | -outputDir=<output-directory> | -max=<maximum-number-for-counter> | -remove=<remove-old-files>]
```

### Defaults

Parameter | default
-----| ---- |
name | counter |
outputDir | same as glob directory |
max  | 1000  (0001, 0002, ...) |
remove | false | 