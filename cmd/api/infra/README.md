# Terraform drift management

## Manipulating terraform state drift
```
terraform plan -out drift.plan -no-color 
terraform show -json drift.plan drift.json  > myd.json 
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "update" | wc -l
# do some if statements to check for create, update and delete
```

Terraform plan options
```
terraform plan -out drift.plan -no-color 
terraform show -json drift.plan drift.json  > myd.json 
cat myd.json | jq .

terraform plan -out drift.tfplan -no-color 
terraform show -json drift.tfplan drift.json  > myd.json 
cat myd.json | jq . | cat myd.json | jq "." | cat myd.json | jq "."

# check create resources changes
cat myd.json | jq -r ".resource_changes[].change.actions[]"
# Grep any action that says create
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create"
# count the number of resources to be created
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l | 's/^[[:space:]]*//g'

# check update resources changes
cat myd.json | jq -r ".resource_changes[].change.actions[]"
# Grep any action that says update
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "update"
# count the number of resources to be updated
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l | 's/^[[:space:]]*//g'

# check delete resources changes
cat myd.json | jq -r ".resource_changes[].change.actions[]"
# Grep any action that says update
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "delete"
# count the number of resources to deleted
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l
cat myd.json | jq -r ".resource_changes[].change.actions[]" | grep "create" | wc -l | 's/^[[:space:]]*//g'



# other options to manipulate terraform plan file
terraform show -json drift.plan drift.json
terraform show -no-color -json drift.plan drift.json  > myd.json 
terraform plan -refresh-only -no-color > drift.json
terraform plan -no-color > drift.json
terraform plan -refresh-only -out drift.json
terraform plan -out drit2.json
terraform plan -no-color -out drit2.json
terraform plan -no-color > data.txt 
```

```
terraform apply -refresh-only
terraform state list
```

What I need to parse the drift.json

- jq expression
- sed command 
- grep command 
- wc -l 
curl 'http://stash.compciv.org/congress-twitter/json/joni-ernst.json' > ernst.json