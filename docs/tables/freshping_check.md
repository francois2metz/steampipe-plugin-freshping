# Table: freshping_check

In Freshping, a Check refers to the specifications configured to monitor an end-point. Checks can be performed for (HTTP/ HTTPS, ICMP ping, UDP, TCP, WebSocket, DNS)

## Examples

### List all checks

```sql
select
  id,
  name,
  url
from
  freshping_check;
```

### List checks in failure

```sql
select
  id,
  name,
  url,
  status
from
  freshping_check
where
  status not in ('AV', 'PS')
```

### List checks with degraded performance

```sql
select
  id,
  name,
  url,
  status
from
  freshping_check
where
  performance_status='DP'
```
