# StorageConfigResponse

## Example Usage

```typescript
import { StorageConfigResponse } from "supabase/models/components";

let value: StorageConfigResponse = {
  fileSizeLimit: 407183,
  features: {
    imageTransformation: {
      enabled: false,
    },
  },
};
```

## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `fileSizeLimit`                                                          | *number*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `features`                                                               | [components.StorageFeatures](../../models/components/storagefeatures.md) | :heavy_check_mark:                                                       | N/A                                                                      |