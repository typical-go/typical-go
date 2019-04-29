# Typical Code Generator (WIP)

## Commands 

### Typi-Go Tool

`typi-go` used to bootstrap new project and migrate common go project to typical go project

General:
- [ ] `typi-go version`: show version of `typi-go`
- [ ] `typi-go archtype`: show list of available architecture type

Core functional:
- [ ] `typi-go new [archtype] [name]`: create new project based on architecture type 
- [ ] `type-go init [archtype]`: generate `typi-gen` and meta information of the project in current directory
- [ ] `type-go init [archtype] --force`: same with `init`, but will override everything

 
### Typi-Gen Tool

`typi-gen` should be included and committed in every typical go project.

`typi-gen` mainly help to generate entity/layer, put it the right place, setup the Dependency Inject and update the project readme/documentation.

General:
- [ ] `typi-gen update`: update metadata in current directory
- [ ] `typi-gen upgrade`: upgrade `typi-gen`
- [ ] `typi-gen version`: show current version of `typi-gen`
- [ ] `typi-gen about`: show general information of this project

Core functional:
- [ ] `typi-gen add [type] [name]`: add new entity to the project
- [ ] `type-gen type`: show list of type
- [ ] `typi-gen mock`: generate mock class


## Architecture Type

Currently only support `rest` architecture. 

In the future, each architecture will have 2 repository: 
1. `typical-[archtype_name]-go`: act as experimental and complete example of respective architecture
2. `archtype-[archtype_name]`: act as handler to `typi-gen` tool

## Metadata

The underlying information of typical go project will be store at `.typical-go` folder which is contain appcontext and entity json file

Detail of `_appctx.json`
```js
{
  "name":"[name]",
  "architecture": {
    "type": "[type]",
    "version": "[version]"
  }
}
```

### Contributing

It's a baby born project right now. Please contact me directly at <iman.tung@gmail.com> for any contribution. Any help would be most welcome.

### License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details




