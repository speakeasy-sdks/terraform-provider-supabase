# SnippetMeta

## Example Usage

```typescript
import { SnippetMeta } from "supabase/models/components";

let value: SnippetMeta = {
  id: "<id>",
  insertedAt: "<value>",
  updatedAt: "1737523146867",
  type: "sql",
  visibility: "project",
  name: "<value>",
  project: {
    id: 118727,
    name: "<value>",
  },
  owner: {
    id: 317983,
    username: "Harry_Pouros",
  },
  updatedBy: {
    id: 93940,
    username: "Kira13",
  },
};
```

## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `insertedAt`                                                           | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `updatedAt`                                                            | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `type`                                                                 | [components.Type](../../models/components/type.md)                     | :heavy_check_mark:                                                     | N/A                                                                    |
| `visibility`                                                           | [components.Visibility](../../models/components/visibility.md)         | :heavy_check_mark:                                                     | N/A                                                                    |
| `name`                                                                 | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `description`                                                          | *string*                                                               | :heavy_minus_sign:                                                     | N/A                                                                    |
| `project`                                                              | [components.SnippetProject](../../models/components/snippetproject.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `owner`                                                                | [components.SnippetUser](../../models/components/snippetuser.md)       | :heavy_check_mark:                                                     | N/A                                                                    |
| `updatedBy`                                                            | [components.SnippetUser](../../models/components/snippetuser.md)       | :heavy_check_mark:                                                     | N/A                                                                    |