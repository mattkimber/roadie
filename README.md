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

## Usage

To process the set JSON file `set.json`:

`roadie set.json`

## set.json

`set.json` contains the following mandatory elements:

* `grf`: information about the GRF itself

And the following optional elements:

* `sprites`: information about sprites
* `cargotable`: cargo tables
* `sprite_templates`: an array of sprite size templates 
* `static_templates`: templates to process that are not related to sprites

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
* `value_names`: if this is set the numerical values will be overridden by the selected names. The length of this array should be (`max_value` - `min_value`), and `min_value` must be 0 when this is used.


### sprites

Example:

```json
"sprites": {
  "table": "example/table.csv",
  "template_directory": "example/templates",
  "additional_text_field": "text",
  "additional_text_format": "Usage: %s"
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

There are some additional built-in functions on top of the standard Go
template functions to make life easier for some common `.nml` scenarios.

`sprites` has the following elements:

* `table`: the path to the tracking table (defaults to `table.csv`)
* `template_directory`: the root path in which templates can be found (defaults to the current path)
* `additional_text_field`: if specified, a field in the .csv which contains additional text (e.g. historical details or usage hints) for the sprite
* `additional_text_format`: if specified, a format string for additional text - useful for format elements which are consistent across all sprites. If this does not contain a `%s` string location character, one will be appended.
* `nestable_templates`: if specified, templates that should be available for nesting. This allows you to set up templates which can be repeatedly called from other templates. (See the example `set.json` and attached template files for more details on how this works in practice)

A tracking table must have the following **mandatory** fields:

* `name`: name of sprite, used for populating language files
* `id`: a unique identifier which is used when creating various resources
* `template`: the name of the template to use. This will automatically be prepended by the template directory and appended by `.tmpl`

As Roadie devolves all other responsibilities to the templates, no other
fields are mandatory. However, some additional map fields will be generated
and passed to the template to make life easier:

* `name_string`: the `string(STR_NAME)` reference to the name in the language file
* `additional_text_string`: the `string(STR_ADDITIONAL_TEXT)` reference to the additional text in the language file

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

### static_templates

Example:

```json
"static_templates": [
  {
    "template":  "example/comment.tmpl"
  },
  {
    "template":  "example/disable_items.tmpl",
    "data" : {
      "Feature": "FEAT_ROADVEHS",
      "Start": "0",
      "End": "254"
  }
}
],
```

Static templates allow you to produce elements that are not related to
sprites and not covered by any of the core Roadie functionality - for
example, adjusting `basecosts` or including a `disable_item` block.

Data from the `data` block will be passed to the template. All items in
the map must be strings, even if they represent integer data in the final
`.nml` file.

Template names do not auto-resolve as with sprites; you can put any path 
and extension here.

Given the following templates: 

```text
// example/comment.tmpl
/* Example of a custom template used to output a string */

// example/disable_items.tmpl
disable_item({{.Feature}}, {{.Start}}, {{.End}});
```

The above example will produce the following output:

```text
/* Example of a custom template used to output a string */
disable_item(FEAT_ROADVEHS, 0, 254);
```

## Built-in functions

The following built-in template functions are offered:

* `altsprites`
* `concat`
* `parseint`
* `slice`

These are described below.

### altsprites

Example:

```text
{{ altsprites "bus" "template_rv" 2 }}
```

This is a convention-based function for producing alternative sprites
for 32bpp and extra zoom levels. It uses the standard GoRender conventions
on directory naming and file suffixes.

The parameters are, in order:

* Name of sprite
* Name of template to reference (will be suffixed with scale identifier)
* Maximum zoom level (only 1, 2 and 4 are meaningful)

The example above will produce the following output:

```text
spriteset (spriteset_bus, "1x/bus_8bpp.png" )
{ template_rv_1x() }

alternative_sprites(spriteset_bus, ZOOM_LEVEL_NORMAL, BIT_DEPTH_32BPP, "1x/bus_32bpp.png", "1x/bus_mask.png")
{ template_rv_1x() }

alternative_sprites(spriteset_bus, ZOOM_LEVEL_IN_2X, BIT_DEPTH_8BPP, "2x/bus_8bpp.png")
{ template_rv_2x() }

alternative_sprites(spriteset_bus, ZOOM_LEVEL_IN_2X, BIT_DEPTH_32BPP, "2x/bus_32bpp.png", "2x/bus_mask.png")
{ template_rv_2x() }
``` 

### concat

Example:

```text
{{ concat "foo" "bar" }}
```

This is a simple function to concatenate two strings. The above will
output the single string `foobar`.

### parseint

Example:

```text
{{$intro_yr := parseint .intro_year}}{{if le $intro_yr 1920 -}}
```

This is a simple function to parse a value from the map as an integer
for use with comparison functions

### slice

Example:

```text
{{ $cargotypes := slice "goods,coal" }}
```

This function generates an iterable slice of values from a
comma-separated string.

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

## Versioning

Roadie can handle versioning of the NML output files in one of two ways:

* Auto-incrementing using a `.roadie_version` file
* Via command line flag

If the command line flag `-version` (`-v` for short) is passed thusly:

```text
roadie -version 23 set.json
```

Then the version in the output file will be set to 23. This is useful
for integration in automated pipelines where you want a consistent version
number across multiple files, including those which Roadie is unaware of.

If no command line flag is passed, then Roadie will track the current
version by creating a `.roadie_version` file (if it does not already
exist) and incrementing it each time Roadie runs. If you are converting
an existing set to Roadie it is possible to edit this file manually
to set the starting point.