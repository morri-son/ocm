## ocm ocm componentarchive create &mdash; Create New Component Archive

### Synopsis

```
ocm ocm componentarchive create [<options>] <component> <version> <provider> <path> {<label>=<value>}
```

### Options

```
  -f, --force         remove existing content
  -h, --help          help for create
  -t, --type string   archive format (default "directory")
```

### Description


Create a new component archive. This might be either a directory prepared
to host component version content or a tar/tgz file.


### SEE ALSO

##### Parents

* [ocm ocm componentarchive](ocm_ocm_componentarchive.md)	 - Commands acting on component archives
* [ocm ocm](ocm_ocm.md)	 - Dedicated command flavors for the Open Component Model
* [ocm](ocm.md)	 - Open Component Model command line client
