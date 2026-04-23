# Runbook

## Common Failures and Recovery
- **Invalid API Keys**: Ensure `MAPBOX_API_KEY` and `TOLLGURU_API_KEY` are correctly set in the environment.
- **API Rate Limits**: Check the usage dashboards for Mapbox and TollGuru if requests are failing with 429 status codes.
- **Detailed Errors**: Inspect the JSON response body from TollGuru; it often contains descriptive error messages regarding polyline parsing or parameter mismatches.
- **Network Issues**: Ensure the server has outbound internet access to reach `api.mapbox.com` and `apis.tollguru.com`.

## Rollback Basics
- Use Git to revert to a previous stable commit: `git revert <commit-hash>`.

## On-call / Logs / Monitoring
- **Logs**: Standard output (STDOUT) and standard error (STDERR) capture the API responses and any errors.
- **Monitoring**: 
  - [Mapbox Dashboard](https://account.mapbox.com/)
  - [TollGuru Dashboard](https://tollguru.com/developers/dashboard)

## Cron Jobs / Scheduled Jobs
- None.

## Data Backfill Scripts
- None.
