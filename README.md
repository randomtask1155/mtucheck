# mtucheck
check for mtu packet drops on container network



# CLI Usage

checking for header response with the promise of no data being sent in response body

```
curl -k -X GET https://mtucheck.apps.domain/header
```

Checking for response body reaponse by specifing how many characters you want returned

```
curl -k -X POST https://mtucheck.apps.domain/data -d { "dataString": 8192 }
```



