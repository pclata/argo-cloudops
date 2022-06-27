#!/bin/bash -e


vaultToken="1234"
vaultHost="https://url:8200/"

# call vault for projects/approles
appRoleResponse=$(curl \
	--header "X-Vault-Token: $vaultToken" \
	--request LIST \
	"$vaultHost/v1/auth/approle/role" \
	| jq -r '.data.keys')

#for appRole in $(echo "$appRoleResponse" | jq -r '.[]); do
#done

# call vault for secret id accessor for each approle/project
# if multiple, triage
#secretIdAccessors=$(curl \
#--header "X-Vault-Toke: $vaultToken" \
#--request LIST \
#"$vaultHost/v1/auth/approle/role/something/secret-id")



# insert into tokens DB