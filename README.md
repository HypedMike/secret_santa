# SECRET SANTA

## Description
A simple program that takes as input a dataset composed like this

```json
{
    "people": [
        "jhon",
        "marta",
        "paul",
        "sarha",
        "peter",
        "laura",
        "michele",
    ],
    "relationships": [
        {
            "p1": "michele",
            "p2": "marta"
        },
        {
            "p1": "jhon",
            "p2": "sarha"
        }
    ]
}
```

The program will match each person with another person in the dataset taking into consideration
the relationships between people. Partners will not be matched with each other.

The program will also obscure th results using a simple substitution cipher.

## Usage

You will first need to set an env variable as
```bash
export FILE_PATH=<your_key>
```

### Build
```bash
go build -o secret_santa
```

### Run
Get the results
```bash
go run secret_santa run
```

### Decypher
```bash
go run secret_santa decypher <cypher_text>
```
