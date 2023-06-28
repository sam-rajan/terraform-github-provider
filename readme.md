# Custom Terraform Provider for creating GitHub Repository 

This repository serves as a sample project for implementing a custom Terraform provider for creating GitHub repositories. It demonstrates the basic structure and functionality required to create a provider that interacts with the GitHub API.

## Prerequisites

Before getting started, ensure that the following prerequisites are met:

- [Terraform](https://www.terraform.io/downloads.html) is installed on your local machine.
- [Go](https://golang.org/dl/) is installed on your local machine.
- A valid GitHub account and personal access token to authenticate with the GitHub API.

## Getting Started

To get started with this project, follow these steps:

1. Clone this repository to your local machine:

   ```shell
   git clone https://github.com/sam-rajan/terraform-github-provider.git
   ```

2. Change into the cloned directory:

   ```shell
   cd terraform-github-provider
   ```

3. Set up the required environment variables:

   ```shell
   export TF_VAR_GITHUB_TOKEN=your-personal-access-token
   ```

4. Build the provider plugin:

   ```shell
   go build -o terraform-provider-mygithub
   ```

5. Initialize the Terraform workspace:

   ```shell
   terraform init
   ```

6. Update the `main.tf` file with the desired repository configuration.

7. Plan and apply the Terraform configuration:

   ```shell
   terraform plan
   terraform apply
   ```

   Review the changes and confirm the creation of the GitHub repository.

8. When you are done experimenting, you can destroy the created resources:

   ```shell
   terraform destroy
   ```

## Customizing the Provider

To customize the provider implementation or extend its functionality, you can modify the `terraform/provider.go` file. This file contains the logic for interacting with the GitHub API and handling the necessary CRUD operations.

Feel free to explore the code and make changes according to your requirements.

## Contributing

Contributions to this sample project are welcome! If you find any issues or have ideas for improvements, please open an issue or submit a pull request.