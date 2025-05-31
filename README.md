# aws-lambda-go

## Call lambda

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" \
  -d '{"httpMethod":"GET", "path":"/health"}'
```