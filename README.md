# Discord Chat Exporter csv parsers

<sub> Now with a much faster version written in go! </sub>

<sub> Each go executable is in it's own sub dir while the python mains in the root dir</sub>

## How to use the go version

To run the go versions all you need to do is download the executable for your platform and run it from the terminal/powershell:

```bash
./IdParser[.exe] -i [Path to folder/file] -o [Path to output .csv (Defaults to ./out.csv)]
```

Use the same output between multiple runs works and takes into account old data to ensure no dupes

```bash
./LocateID[.exe] -i [Path to folder/fil] -id "[User ID in quotes]" -o [Path to output csv (Optional)]
```

## How to use the Python version

```bash
py ./main.py [csv file or dir] [output csv]
```

Use the same output between multiple runs works and takes into account old data to ensure no dupes

``` bash
py ./locate_id.py [User ID] [csv file or dir]
```
