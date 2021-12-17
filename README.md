# Terrarium - Self hosted private registry

Terrarium is a Terraform registry implementation designed to provide a subset of Terraform Cloud functionality in a self hosted manner

## Contributing

To get started contributing please:

1. Clone the project
1. Run `godoc -notes ".*"`
1. Visit [http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/](http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/) 
1. Review the [drivers package](http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/pkg/registry/drivers/) as this contains interfaces for third parties to implement and extend Terrarium. See current [database implementations](./internal/database) and [storage implementations](./internal/storage) for examples and add your own!
1. Once your implementation is ready don't forget to modify the [module command](./cmd/module.go) to expose your support in the CLI

