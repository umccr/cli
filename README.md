# UMCCR cli tool

Organization command line tool to ease common operations, focusing on helping researchers to transition from HPLAC (High Performance Low Availability Computing) to cloud computing.

This CLI tool is based on the [Go cobra CLI framework](https://github.com/spf13/cobra), used by "many of the most widely used Go projects".

# Quickstart

After downloading the CLI release for your platform:

```bash
$ wget https://github.com/umccr/cli/releases/download/umccr -O /usr/local/bin/umccr
```

Just run one of the available commands:

```bash
$ umccr find foo
{
  ResultSet: {
    ResultSetMetadata: {
      ColumnInfo: [{
          CaseSensitive: true,
          CatalogName: "hive",
          Label: "key",
          Name: "key",
          Nullable: "UNKNOWN",
          Precision: 2147483647,
          Scale: 0,
          SchemaName: "",
          TableName: "",
          Type: "varchar"
        }]
    },
    Rows: [
      {
        Data: [{
            VarCharValue: "key"
          }]
      },
      {
        Data: [{
            VarCharValue: "foo-P016-merged.csv"
          }]
      },
      {
        Data: [{
            VarCharValue: "foo-merged-template.yaml"
          }]
      },
```
