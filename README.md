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

* `sprites`: information about sprites
* `cargotable`: cargo tables
* `sprite_templates`: an array of sprite size templates 

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
    "language": "english", 
    "parameters": []
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
* `parameters`: any parameters this grf will define

### parameters

Example:

```json
"parameters": [
  {
    "id": "example",
    "name": "Example Behaviour",
    "desc":  "Set an example behaviour somehow",
    "type": "int",
    "def_value": 1,
    "min_value": 0,
    "max_value": 2,
    "names": [ "Low", "Medium", "High" ]
  }
]
```

The `parameters` element has the following fields, closely following the
NML specification:

* `id`: a short identifier for this parameter
* `name`: the name of this parameter
* `desc`: a description of this parameter
* `type`: type of parameter, `int` or `bool`
* `default_value`: default value
* `min_value`: minimum value (`int` parameters only)
* `max_value`: maximum value (`int` parameters only)
* `value_names`: if this is set the numerical values will be overridden by the selected names. The length of this array should be (`max_value` - `min_value`).


### sprites

Example:

```json
"sprites": {
  "table": "example/table.csv",
  "template_directory": "example/templates"
}
```

The `sprites` element controls how to turn a `.csv` tracking table into NML
sprites. This is done by specifying a tracking table and a directory for
templates.

Under the hood this is a thin layer on top of the Golang templating system,
and templates are standard Go template files. Each column from the tracking
table will be passed as a Go map, allowing for an almost unlimited variety
of sprite handling options (at the cost of needing to write potentially
complex templates).

`sprites` has the following elements:

* `table`: the path to the tracking table (defaults to `table.csv`)
* `template_directory`: the root path in which templates can be found (defaults to the current path)

A tracking table must have the following **mandatory** fields:

* `name`: name of sprite, used for populating language files
* `id`: a unique identifier which is used when creating various resources
* `template`: the name of the template to use. This will automatically be prepended by the template directory and appended by `.tmpl`

As Roadie devolves all other responsibilities to the templates, no other
fields are mandatory. However some additional map fields will be generated
and passed to the template to make life easier:

* `name_string`: the `string(STR_NAME)` reference to the name in the language file

Note: it is a good idea to leave some trailing newlines in your templates.

### cargotable

Example:

```json
"cargo_table": ["COAL","IORE","PASS"]
```

A cargo table to be defined. This is a string array, and is translated
directly to the NML output if it is present.

### sprite_templates

Example:

```json
"sprite_templates": [
  {
    "names": ["template"],
    "location_scales": [1],
    "size_scales": [1],
    "offset_scales": [0.25],
    "locations" : [[7,0],[40,0],[80,0],[120,0],[167,0],[200,0],[240,0],[280,0]],
    "sizes": [[10,26],[26,26],[32,18],[26,26],[10,26],[26,26],[32,18],[26,26]],
    "offsets": [[-16,-47],[-67,-52],[-75,-56],[-26,-52],[-16,-55],[-66,-52],[-55,-55],[-28,-52]]
  }
]
```

The way Roadie handles sprite templates is a little complex at first glance, 
but with this complexity comes a lot of flexibility and the ability to
quickly adjust things like offsets across multiple zoom levels at once.

`sprite_templates` is an array of individual template block objects, each of
which are configured by several sub-arrays:

* `names` is the names of the templates this particular template block will product
* `location_scales` is the list of values to scale the x,y locations within spritesheets by
* `size_scales` is the list of values to scale the x,y heights within spritesheets by
* `offset_scales` is the list of values to scale the x,y offsets within spritesheets by
* `locations`: are the x,y pairs for each sprite in the spritesheet (typically there will be 1, 4 or 8 pairs depending on the type of `.grf`)
* `sizes`: are the x,y pairs for the height of each sprite
* `offsets`: are the x,y pairs for the offset of each sprite

The `locations`, `sizes` and `offsets` will be multiplied by the relevant
scale value when generating the output templates. For example, if you were
working on a set with 2x and 4x zoom your template declaration might look
something like this:

```json
"sprite_templates": [
  {
    "names": ["templates_1x", "templates_2x", "templates_4x"],
    "location_scales": [1,2,4],
    "size_scales": [1,2,4],
    "offset_scales": [0.25,0.5,1],
    "locations" : [[10,15]],
    "sizes": [[15,15]],
    "offsets": [[16,-12]]
  }
]
```

Which will produce the following output:

```
template templates_1x() {
    [ 10, 15, 15, 15, 4, -3 ]
}

template templates_2x() {
    [ 20, 30, 30, 30, 8, -6 ]
}

template templates_4x() {
    [ 40, 60, 60, 60, 16, -12 ]
}
```

## Language files

Roadie will output its language file in the `lang` directory, as expected
by NML. If this directory does not exist it will be created.

Roadie will attempt to write the correct language pragma per
https://newgrf-specs.tt-wiki.net/wiki/NML:Language_files if you use
the lowercase name of the language, strip any brackets and replace
spaces with underscores. In case it cannot identify the language
it will default to English(GB), 0x01.

It is the business of the future to be dangerous. It is *not* the business
of Roadie to handle multiple translations. Roadie will output a single 
strings file in the language of your choice, but you will need to create 
your own translations manually.