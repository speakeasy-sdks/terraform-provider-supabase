# SnippetList

## Example Usage

```typescript
import { SnippetList } from "supabase/models/components";

let value: SnippetList = {
  data: [
    {
      id: "<id>",
      insertedAt: "<value>",
      updatedAt: "1737533880126",
      type: "sql",
      visibility: "user",
      name: "<value>",
      project: {
        id: 110375,
        name: "<value>",
      },
      owner: {
        id: 656330,
        username: "Britney9",
      },
      updatedBy: {
        id: 96098,
        username: "Weston_Wolff73",
      },
    },
  ],
};
```

## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `data`                                                             | [components.SnippetMeta](../../models/components/snippetmeta.md)[] | :heavy_check_mark:                                                 | N/A                                                                |