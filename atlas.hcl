env "local" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
}

env "dev" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
}

env "prod" {
  src = "ent://ent/schema"
  dev = "mysql://user:password@localhost:3307/dbname_dev"
  url = "mysql://user:password@localhost:3306/dbname"
} 