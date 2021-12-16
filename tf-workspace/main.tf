module "test" {
    source = "api.terrarium.cloud/test/test/aws"
    version = "0.0.1"
}

module "test2" {
    source = "api.terrarium.cloud/nonexistent/test/aws"
    version = "0.0.1"
}