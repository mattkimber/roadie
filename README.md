# Roadie

Roadie is (yet another) template-based system for producing OpenTTD NML files.
It is designed so that after initial setup of appropriate templates and
basic configuration, new sprites can be added simply by appending to a .csv
tracking table.

## Prerequisites

You will need the following prerequisites to create an NML file:

* A `set.json` describing the set and templates
* A `table.csv` tracking table for the sprites

The files don't have to be named exactly that, but all examples will
use these conventions.

## set.json

`set.json` contains the following mandatory elements:

* `grf`: information about the GRF itself

And the following optional elements:

* `cargotable`: cargo tables

Each element is described in detail below.

### grf

Example:

```json
"grf": {
    "author_id": "RDE",
    "grf_id": 1,
    "name": "Roadie Example Vehicle Set",
    "description": "A set demonstrating the basic features of Roadie",
    "min_compat_version": 1,
    "filename": "example.nml",
    "language": "english"
}
```

The `grf` element has the following fields:

* `author_id`: a three-character author ID you use across all of your `.grf` files
* `grf_id`: the ID of this `.grf` file in particular, from 0-255
* `name`: the name of the vehicle set
* `description`: a short description of the set
* `min_compat_version`: the earliest version of the set this instance of the set is compatible with
* `filename`: the output NML filename to produce
* `language`: the language in which to output string tables. Defaults to `english` if not set.

### cargotable

Example

```json
"cargo_table": ["COAL","IORE","PASS"]
```

A cargo table to be defined. This is a string array, and is translated
directly to the NML output if it is present.