# Table: freshping_check

In Freshping, a Check refers to the specifications configured to monitor an end-point. Checks can be performed for (HTTP/ HTTPS, ICMP ping, UDP, TCP, WebSocket, DNS)

## Examples

### List checks

```sql
select
  id,
  name,
  url
from
  freshping_check;
```
