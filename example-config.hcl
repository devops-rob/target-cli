###### option 1 #######

nomad "dev" {
    endpoint  = "https://dev-nomad.com:4646"
    region    = "uk"
    namespace = "dev"
}

nomad "staging" {
    endpoint  = "https://staging-nomad.com:4646"
    region    = "uk"
    namespace = "staging"
}

vault "dev" {
    endpoint  = "https://dev-vault.com:8200"
    namespace = "dev"
}

vault "staging" {
    endpoint  = "https://staging-vault.com:8200"
    namespace = "staging"
}


###### option 2 #######
nomad {
  profile "dev" {
    endpoint  = "https://dev-nomad.com:4646"
    region    = "uk"
    namespace = "dev"

  }
  profile "staging" {
    endpoint  = "https://staging-nomad.com:4646"
    region    = "uk"
    namespace = "staging"

  }
}

vault {
  profile "dev" {
    endpoint  = "https://dev-vault.com:8200"
    namespace = "dev"

  }
  profile "staging" {
    endpoint  = "https://staging-vault.com:8200"
    namespace = "staging"
  }
}
