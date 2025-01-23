# SnippetResponse

## Example Usage

```typescript
import { SnippetResponse } from "supabase/models/components";

let value: SnippetResponse = {
  id: "<id>",
  insertedAt: "<value>",
  updatedAt: "1737595894148",
  type: "sql",
  visibility: "user",
  name: "<value>",
  project: {
    id: 576157,
    name: "<value>",
  },
  owner: {
    id: 592042,
    username: "Kianna84",
  },
  updatedBy: {
    id: 297437,
    username: "Raquel88",
  },
  content: {
    favorite: false,
    schemaVersion: "<value>",
    sql: "<value>",
  },
};
```

## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `id`                                                                                         | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `insertedAt`                                                                                 | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `updatedAt`                                                                                  | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `type`                                                                                       | [components.SnippetResponseType](../../models/components/snippetresponsetype.md)             | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `visibility`                                                                                 | [components.SnippetResponseVisibility](../../models/components/snippetresponsevisibility.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `name`                                                                                       | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `description`                                                                                | *string*                                                                                     | :heavy_minus_sign:                                                                           | N/A                                                                                          |
| `project`                                                                                    | [components.SnippetProject](../../models/components/snippetproject.md)                       | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `owner`                                                                                      | [components.SnippetUser](../../models/components/snippetuser.md)                             | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `updatedBy`                                                                                  | [components.SnippetUser](../../models/components/snippetuser.md)                             | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `content`                                                                                    | [components.SnippetContent](../../models/components/snippetcontent.md)                       | :heavy_check_mark:                                                                           | N/A                                                                                          |