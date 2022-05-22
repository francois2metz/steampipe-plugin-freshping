# Table: freshping_user

List users who have access to the Freshping organization.

## Examples

### List users

```sql
select
  id,
  name,
  role
from
  freshping_user;
```

### List admin users

```sql
select
  id,
  name
from
  freshping_user
where
  role='Admin';
```

### List read-only users

```sql
select
  id,
  name
from
  freshping_user
where
  role='Read Only';
```
