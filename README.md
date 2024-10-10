# 1 billion row challenge - golang

The challenge is not new, I found the challenge in this
[repo in github](https://github.com/gunnarmorling/1brc/tree/main) and I decided
to do it in golang.

Why in golang? Because I'm learning it and I want to improve my skills.

## The challenge

The text file contains temperature values for a range of weather stations. Each
row is one measurement in the format <string: station name>;<double:
measurement>, with the measurement value having exactly one fractional digit.
The following shows ten rows as an example:

```text
Hamburg;12.0
Bulawayo;8.9
Palembang;38.8
St. John's;15.2
Cracow;12.6
Bridgetown;26.9
Istanbul;6.2
Roseau;34.4
Conakry;31.2
Istanbul;23.0
```

The task is to write a Go program which reads the file, calculates the min,
average, and max temperature value per weather station, and emits the results on
stdout like this (i.e. sorted alphabetically by station name, and the result
values per station in the format <min>/<average>/<max>, rounded to one
fractional digit.

## How to generate the file

Follow the official instructions in the
[original repo](https://github.com/gunnarmorling/1brc/tree/main?tab=readme-ov-file#running-the-challenge)
to generate the file.

Additional, you can generate the file by run the following command:

```bash
python ./scripts/create_measurements.py 1000000000
```

## How to run the code

```bash
go run main.go --file <path_to_file> --version <version>
```

## Scoreboard

- **v1**: Time taken to complete the challenge:  3m19.621960181s
- **v2**: Time taken to complete the challenge:  3m3.186535997s
- **v3**: Time taken to complete the challenge:  26.28551254s
