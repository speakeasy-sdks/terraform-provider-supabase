# V1DatabaseResponse

## Example Usage

```typescript
import { V1DatabaseResponse } from "supabase/models/components";

let value: V1DatabaseResponse = {
  host: "bright-agreement.net",
  version: "<value>",
  postgresEngine: "<value>",
  releaseChannel: "<value>",
};
```

## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `host`             | *string*           | :heavy_check_mark: | Database host      |
| `version`          | *string*           | :heavy_check_mark: | Database version   |
| `postgresEngine`   | *string*           | :heavy_check_mark: | Database engine    |
| `releaseChannel`   | *string*           | :heavy_check_mark: | Release channel    |