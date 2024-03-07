#!/bin/bash

# Specify the AWS Secret ID
secret_id="go-accounts"

# Retrieve the secret from AWS Secrets Manager
secret_value=$(aws secretsmanager get-secret-value --secret-id $secret_id --output json)

# Extract the secret string from the JSON response using jq
secret_string=$(echo "$secret_value" | jq -r '.SecretString')

# Check if secret string is not empty
if [ -n "$secret_string" ]; then
    # Create app.env file and write secret values to it
    echo "$secret_string" | jq -r 'to_entries | .[] | "\(.key)=\(.value)"' > app.env
    echo "Secrets successfully written to app.env file"
else
    echo "Failed to retrieve secret value from AWS Secrets Manager"
    exit 1
fi