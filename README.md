# Interviewtopia Go assessment policy API

Static household tax-policy fixtures for the change request in the Interviewtopia Go assessment.

- Swagger UI: <https://masseyis.github.io/interviewtopia-go-tax-policy-api/docs/>
- OpenAPI definition: <https://masseyis.github.io/interviewtopia-go-tax-policy-api/openapi.yaml>

The endpoint pattern is:

```text
https://masseyis.github.io/interviewtopia-go-tax-policy-api/policies/{residents}.json
```

`residents` is an integer from 1 to 6. Responses use USD integer cents and whole basis points. The calculator remains non-cumulative and rounds half up to the nearest cent; those are fixed rules from the base exercise rather than response fields.

Band order is not guaranteed. The six-resident fixture deliberately contains 1,000 bands so the assessment can exercise policy validation, one-time normalisation and ordered lookup.

The endpoint contains no personal, secret or production data.

