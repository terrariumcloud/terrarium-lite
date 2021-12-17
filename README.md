# Terrarium - Self hosted private registry

Terrarium is a Terraform registry implementation designed to provide a subset of Terraform Cloud functionality in a self hosted manner

## Development Requirements

**This is subject to change!**

For development the following assumptions are made:

* Development is done on AWS EC2 instance with IAM role full access to S3 and DynamoDB. This machine also requires `docker` and `docker-compose`
* An S3 bucket has been manually created. The name of which can be specified via `--storage-backend-name` on the CLI. Defaults to `terrarium-dev`
* Terrarium on startup will created the required DynamoDB tables

To get started:

1. Ensure your chosen S3 bucket has been created and your EC2 instance has a correct IAM role (you may need to create this)
1. SSH into your EC2 instance
1. Review the `docker-compose.yml` and make any changes needed for your development
1. Run `docker-compose up`

## Contributing

To get started contributing please:

1. Clone the project
1. Run `godoc -notes ".*"`
1. Visit [http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/](http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/) 
1. Review the [drivers package](http://localhost:6060/pkg/github.com/dylanrhysscott/terrarium/pkg/registry/drivers/) as this contains interfaces for third parties to implement and extend Terrarium. See current [database implementations](./internal/database) and [storage implementations](./internal/storage) for examples and add your own!
1. Once your implementation is ready don't forget to modify the [module command](./cmd/module.go) to expose your support in the CLI

