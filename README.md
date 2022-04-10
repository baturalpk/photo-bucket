# Getting Started
A demonstrative HTTP API, which follows REST standards, is basically a kind of photo upload service.
Configure, run, and create a profile to start uploading your favorite photos. <br><br>
Follow the guide to complete the installation. Upon the successful completion of [Setup](#setup) & [Deployment](#deployment) phases, the app will be
accessible from http://localhost:8080. See [Endpoints](#endpoints) section to learn more about its usage.

## Endpoints
Most **request bodies** can be delivered either as `JSON payload` or in `x-www-form-urlencoded` format until stated otherwise.

### Profiles Endpoints:
#### - POST /profiles/signup/email : <br>
*Request body (all fields are **required**) :*
```json
{
  "email":    "...@...",
  "username": "...",
  "password": "***" 
}
```
#### - POST /profiles/signin : <br>
*Request body ('password' and " 'username' **OR** 'email' " are required) :*
```json
{
  "email":    "...@...",
  "username": "...",
  "password": "***" 
}
```
#### - GET /profiles/{username} : <br>
Retrieves a specific profile by username.
### Photos Endpoints:
#### - POST /photos/ : <br>
**⚠️Expects multi-part form data** <br>
*Request body ('photo' is **required**) :*
```text
{
  "photo":       < multi-part image file > // jpeg (jpg) is currently the only supported format ( < 30 MB )
  "tags":        < Array[string] >
  "description": < string >
}
```
#### - GET /photos/owner/{id} : <br>
Retrieves a photo collection by its owner's ID.
#### - GET /photos/{id} : <br>
Retrieves a specific photo's details by ID.

## Prerequisites
1. Access to an Amazon S3 or S3 compatible (e.g., Backblaze B2) storage service
2. Docker (engine >= 18.06.0)
3. Git CLI
4. Unix-like environment that supports building Makefiles, e.g., MacOS, Linux, Windows Subsystem for Linux, Cygwin
5. (Optional) PostgreSQL DBMS >= 10.x (DBaaS or local)

## Setup
1. Open a shell and clone the repository with:<br>
    ```git clone https://github.com/baturalpk/photo-bucket.git```
2. Change the current directory to the newly created `photo-bucket` directory:<br>
    ```cd photo-bucket```
3. Create a file named `secrets.yaml` with the same content of `config/secrets.template.yaml` under `config` directory:<br>
    ```cp config/secrets.template.yaml config/secrets.yaml```
4. (If **prerequisite 5** is not ensured) Create a file named `db.env` with the same content of `config/db.template.env` under `config` directory:<br>
    ```cp config/db.template.env config/db.env```
5. Edit the newly created files `config/secrets.yaml` and `config/db.env` using the specifications given under [Configurations](#configurations) section.

## Configurations

### config/secrets.yaml
**Example:**
```yaml
env: "dev"
es256:
  privateKey: |
    -----BEGIN EC PRIVATE KEY-----
    MHcCy3/TPeoMpGsC7WuUDt4Xvw+nRLRasfasfasfBai3t5ZzUCEQuWoAoGCqGSM49
    AwEHoUQDQgAEAQEEIOy+eNKDTo4uOImTK9p2tgv78OGWGD9MWtnuroh6mUqvMHuJc
    or/E9SlJ/jS2XWRWrRsZkbEMTzx/Ptf3Jg==
    -----END EC PRIVATE KEY-----
  publicKey: |
    -----BEGIN PUBLIC KEY-----
    MFkwEwYHKoZIzj0CAQYsfasfaasfasdgsdagsdfSAFASFASEeoMpGsC7WuUDCSD9
    MWto6upWknurXvw+nRLqvMHuJcor/E9SlJ/jS2XWRWrRsZkbEMTzx/Ptf3Jg==
    -----END PUBLIC KEY-----
postgresURI:
  test: "postgres://<USERNAME>:<PASSWORD>@localhost:5432/photo-bucket-testDB"
  dev: "postgres://<USERNAME>:<PASSWORD>@localhost:5432/photo-bucket-devDB"
  prod: "postgres://<USERNAME>:<PASSWORD>@pgsql:5432/photo-bucket"
s3:
  bucketNames:
    photos:
      test: "my-photos_test"
      dev: "my-photos_dev"
      prod: "my-photos"
  credentials:
    id: "***"
    secret: "***"
  endpoint: "https://s3.us-west-2.***"
  region: "us-west-2"
```
* `env:` "prod" | "dev" | "test" (The deployment sets env as `prod` by default.)
* `es256:` Firstly, you need to generate private and public Elliptic Curve (ES256) keys to be able to sign JSON Web Tokens. 
           Then paste contents of the key files into the proper fields of `secrets.yaml` as multi-line strings. <br>
          **Related section(s) in the following guide would be helpful: https://cloud.google.com/iot/docs/how-tos/credentials/keys**
* `postgresURI:` Postgres connection strings for test, development, and production environments respectively.
* `s3/bucketNames/photos:` S3 bucket names for test, development, and production environments respectively.
* `s3/credentials:` Your ID & API Secret that can be obtained by object storage provider.
* `s3/endpoint` & `s3/region:` The S3 protocol compatible endpoint and region values that can be obtained by object storage provider.

### config/db.env
**Example:**
```.env
POSTGRES_USER=docker
POSTGRES_PASSWORD=docker
POSTGRES_DB=photo-bucket
```
* POSTGRES_USER: Your Postgres username that will be used at the initialization phase.
* POSTGRES_PASSWORD: Your Postgres password that will be used at the initialization phase.
* POSTGRES_DB: Your Postgres database name that will be used at the initialization phase.

## Deployment
0. Ensure that you have completed the steps given under [Setup](#setup) section successfully.
1. Ensure that the Docker engine is operating healthy.
2. Open a shell at the project's root directory (`photo-bucket`).
3. Run the target `deployment` from `Makefile` with: ```make deployment```
4. After a short time, the application will be accessible from: ```http://localhost:8080```

### Building Binary Executables (Optional)
Optionally you can create binary executables for various OS platforms:

#### MacOS (darwin)
0. Ensure that you have completed the steps given under [Setup](#setup) section successfully.
1. Open a shell at the project's root directory (`photo-bucket`).
2. Run the target `darwin_build` from `Makefile` with: ```make darwin_build```
3. Upon the successful completion of build process, executable can be found under the newly created `build` directory.

#### Linux
0. Ensure that you have completed the steps given under [Setup](#setup) section successfully.
1. Open a shell at the project's root directory (`photo-bucket`).
2. Run the target `linux_build` from `Makefile` with: ```make linux_build```
3. Upon the successful completion of build process, executable can be found under the newly created `build` directory.

#### Windows
0. Ensure that you have completed the steps given under [Setup](#setup) section successfully.
1. Open a shell at the project's root directory (`photo-bucket`).
2. Run the target `windows_build` from `Makefile` with: ```make windows_build```
3. Upon the successful completion of build process, executable can be found under the newly created `build` directory.

## Tests
Unit tests for `repositories` and `clients` are available under the relevant directories: 
- Repositories: `services/photo/...` & `services/profile/...`
- Clients: `clients/s3client/...` & `clients/postgres/...` & `clients/entclient/...`

❗ Ensure that you have successfully completed the steps given under [Setup](#setup) section before executing tests.
