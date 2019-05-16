[![TravisCI build status](https://travis-ci.org/umccr/cli.svg)](https://travis-ci.org/umccr/cli)

# UMCCR cli tool

Organization command line tool to ease common operations, focusing on helping researchers to transition from HPLAC (High Performance Low Availability Computing) to cloud computing and reducing UX friction.

This CLI tool is based on the [Go cobra CLI framework](https://github.com/spf13/cobra), used by "many of the most widely used Go projects". The main motivation of writing this in Go is that no python virtual environment setup is required, just download the Go binary and off we go!

[This HN thread also helped with the Go-convincing](https://news.ycombinator.com/item?id=19459787).

# Quickstart

After downloading the CLI release for your platform:

```bash
$ wget https://github.com/umccr/cli/releases/download/umccr -O /usr/local/bin/umccr
```

Just run one of the available commands (assuming you have an active STS authenticated session):

```bash
$ awsdev
Google Password:
Open the Google App, and tap 'Yes' on the prompt to sign in ...
Assuming arn:aws:iam::<ACCT>:role/<ROLE>
Credentials Expiration: 2019-05-09 19:15:28+10:00

$ umccr find foo
foo-P016-merged.csv
foo-merged-template.yaml
```
