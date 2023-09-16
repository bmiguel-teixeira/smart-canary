
# Canary Error Rate in Percentage

```
sum(irate(http_requests_total{job="smart-canary", stage="canary", statusCode!="200"}[30s]))
/
sum(irate(http_requests_total{job="smart-canary", stage="canary"}[30s]))
*
100
``````


# Global Error Rate in Percentage

```
sum(irate(http_requests_total{job="smart-canary", statusCode!="200"}[30s]))
/
sum(irate(http_requests_total{job="smart-canary"}[30s]))
*
100
```