#!/bin/bash

set -e
db_host=$PGHOST
db_port=$PGPORT
db_name=$PGDATABASE

vault_token=$VAULT_TOKEN
vault_host=$VAULT_HOST

# get list of projects/approles
approle_response=$(curl -s\
	--header "X-Vault-Token: $vault_token" \
	--request LIST \
	"$vault_host/v1/auth/approle/role" \
	| jq -r '.data.keys')

for approle in $(echo "$approle_response" | jq -r '.[]'); do
  # get list of secret id accessors for each approle/project
  accessors=($(curl -s\
    --header "X-Vault-Token: $vault_token" \
    --request LIST \
    "$vault_host/v1/auth/approle/role/$approle/secret-id" \
    | jq -r '.data.keys' | tr -d '[]," '))

  # if multiple secret accessors exist, print approle for manual triage.
	# manually, find which secret belongs to an accessor, and save accessor ID to tokens DB.
  if [ "${#accessors[@]}" -gt 1 ]; then
		echo "WARN: Multiple accessors exist for AppRole: $approle"
		continue
  fi

  if [ "${#accessors[@]}" -eq 0 ]; then
		echo "WARN: Accessors do not exist for AppRole: $approle"
		continue
  fi

	set +e
	echo "ERROR: Inserting accessor into tokens table for AppRole: $approle"
	psql -h $db_host -p $db_port -w -d $db_name -c "INSERT INTO tokens(token_id, created_at, project) VALUES ('${accessors[0]}', current_timestamp, '$approle')"
	set -e
done

echo "Script complete"
