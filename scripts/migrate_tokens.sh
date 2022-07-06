#!/bin/bash

set -e
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
		echo "ERROR: Multiple accessors exist for AppRole: $approle"
		continue
  fi

  if [ "${#accessors[@]}" -eq 0 ]; then
		echo "ERROR: Accessors do not exist for AppRole: $approle"
		continue
  fi


	#TODO: won't work for internal	
	approle_prefix="argo-cloudops-projects-"
	project=${approle#$approle_prefix}

	exists=$(psql -h $PGHOST -p $PGPORT -d $PGDATABASE -U $PGUSER -w -t -c "SELECT EXISTS(SELECT project from projects where project='$project')")
	if [ $exists == "f" ]; then
		# TODO: figure out why project doesn't exist in DB
		echo "ERROR: Project does not exist in table: $approle"
	fi


	set +e
	#echo "Inserting accessor ID into tokens table for AppRole: $approle"
	psql -h $PGHOST -p $PGPORT -d $PGDATABASE -U $PGUSER -w <<'SQL' | ...
	  INSERT INTO tokens(token_id, created_at, project)
		VALUES ('${accessors[0]}', current_timestamp, '$approle')
SQL
	
	# TODO: catch error
	set -e
done

echo "Script complete"
