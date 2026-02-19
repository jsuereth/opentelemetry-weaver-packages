# Go Codegen Templates

This package contains Jinja templates for generating Go semantic convention constants and instruments using the Weaver V2 data model.

## Usage

```bash
weaver code generate -p templates/codegen/go <registry>
```

## Templates

- `attribute_group.go.j2`: Generates attribute keys and helper functions.
- `metric.go.j2`: Generates metric instruments and helper types.

## Configuration

The `weaver.yaml` file supports the following parameters:

- `tag`: The semantic convention version tag (e.g., `v1.30.0`).
- `excluded_namespaces`: Root namespaces to exclude from generation.
- `excluded_attributes`: Specific attribute keys to exclude.
