# GetApiKeyRequest

## Example Usage

```typescript
import { GetApiKeyRequest } from "supabase/models/operations";

let value: GetApiKeyRequest = {
  ref: "<value>",
  id: "<id>",
  reveal: false,
};
```

## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `ref`              | *string*           | :heavy_check_mark: | Project ref        |
| `id`               | *string*           | :heavy_check_mark: | N/A                |
| `reveal`           | *boolean*          | :heavy_check_mark: | N/A                |