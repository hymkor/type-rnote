type-rnote
==========

Read your releases of github repository and cat the release notes to STDOUT.

- `type-rnote [-r REVISION] USERNAME REPOSITORY`
- `type-rnote [-r REVISION] USERNAME/REPOSITORY`
- `type-rnote [-r REVISION] https://github.com/USERNAME/REPOSITORY`

Square brackets `[ ]` indicate optional arguments.

![image](./screenshot.png)

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/type-rnote/releases) and extract the executable.


### User go install

```
go install github.com/hymkor/type-rnote@latest
```

### Use scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/type-rnote/master/type-rnote.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install type-rnote
```
