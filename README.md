jCfg
---

Load configuration data from a json file.
If the file does not exist or there are missing fields, the user can create it interactively.

##### Example
```
type Settings struct {
  TestField string `json:"testField"`
}

cfg := Parse("./settings.json", &Settings{}).(*Settings)
```

##### Caveats

* Only supports flat json files (no nesting whatsoever)
* Only supports string values
