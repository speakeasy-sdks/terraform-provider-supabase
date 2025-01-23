# V1DeleteNetworkBansRequest

## Example Usage

```typescript
import { V1DeleteNetworkBansRequest } from "supabase/models/operations";

let value: V1DeleteNetworkBansRequest = {
  ref: "<value>",
  removeNetworkBanRequest: {
    ipv4Addresses: [
      "<value>",
    ],
  },
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ref`                                                                                    | *string*                                                                                 | :heavy_check_mark:                                                                       | Project ref                                                                              |
| `removeNetworkBanRequest`                                                                | [components.RemoveNetworkBanRequest](../../models/components/removenetworkbanrequest.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |