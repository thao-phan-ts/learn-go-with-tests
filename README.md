# learn-go-with-tests

## 1. Setup 

```bash
    go version 

    go mod init <project_name>

    mkdir <folder_name>
```


## 2. Use cases

```text
1. 
input: lender_config_id
output: 
- Is activate or deactivate?
- Which lead_source is being used?
- Which flow is being used?
- Which ui_version is being used?

2.
input: lead_source
output: 
- 


vds dang work o nhung flow gi? 
platform la gi? 
telco nao duoc accept?
co check duoc online/offline khong?

Read de -> get 



```

```textmate

0. active:
- true -> onboarding flow, false -> user drop-off & comeback flow
- list of ui_version is used/not used
1. from telco_code ->
- current flow_type of app
- detect telco_code is allowed
2. from lead_source -> get flow_type of app
3. from tags.[i]product_code -> get platform of app:
- product_code = tpb_evo_native -> mobile app
- product_code = tpb_01 -> web app
4. from ui_flow -> write step by step
5. from ui_flow_settings:
- if null -> skip
- if not null -> get value ...

```




# Navigate to the 'data' directory or create it if it doesn't exist
mkdir -p data/<app_id>_<flow_type>

# Create the 'credit_assignment' directory and its files
mkdir -p data/<app_id>/credit_assignment
cd data/<app_id>/credit_assignment
touch evaluation_path.json application_matrix.json

# Create the 'risk_grade' directory and its files
mkdir -p data/<app_id>/risk_grade
touch data/<app_id>/risk_grade/evaluation_path.json
touch data/<app_id>/risk_grade/application_matrix.json

# Create the 'evaluate' directory and its files
mkdir -p data/<app_id>/evaluate
touch data/<app_id>/evaluate/evaluation_path.json
touch data/<app_id>/evaluate/application_matrix.json