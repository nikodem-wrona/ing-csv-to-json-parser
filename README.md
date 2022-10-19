# ing-csv-to-json-parser
Tool for parsing ING account statement from CSV to JSON format

# Usage:
I'm using [Task](https://taskfile.dev/) for managing the scripts.

To build the tool run :

```
$ task build
```

After building you can find the binary file in `bin` folder.

To use it just run it in following way : 
```
$ ing-parser {path to ING csv export}
```
Ro run the tool without building it run : 
```
$ task run
```
