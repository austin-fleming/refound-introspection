# ~~Introspection~~ Users + Funding Pools

## TODO:

- Decide if we want to resolve user verification with a ZK solution.
    - If Celo, ask contact if Plumo is ready. 
    - If Celo, should account record be tied to phone number instead? What might be our account recovery path?
    - If Polygon, ask contact about PolygonID roadmap.
- Cache server for query aggregates to improve on-chain query response.
    - (assuming we cover cost of some actions) command event -> on-chain write -> validate success -> cache invalidation event -> cache de-normalized query
    - This actually needs to be two services: cache + contract events
      - contract event log should be polled to trigger events across the bus, allowing easy checks for cache staleness. This should live beside the cache to reduce network latency.
        - We need to decide what should be banned vs purged
- Implement on-chain repos instead of Postgres mock
- Should users be charged for creating profile/pool contracts, or should the platform eat the cost?
- When beneficiary attempts a death claim on account, Twilio and email notifications should be sent to primary if contact details were provided.
- Should beneficiaries have an entity, or just be a function param?
- Query chain to verify account/beneficiary/contract actually exists. Validation currently only checks address can exist.
- Voting politics within an initiative pool
- Commenting

## Services:

- account
- initiative
- notification
- asset-upload (should tie into Twilio)
    - Needs to reference account phone for Celo
- TBD: on-chain event monitor

## Entities:

- account - a user's entire account (use viewable only)
- profile - a user's publicly viewable information
- initiative - an initiative pool owned by one or more accounts

## Endpoints:

- /api/v1/account

  - "CREATE"
  - 🔐 "ADD_BENEFICIARY" (address) -> add beneficiary to member's contract
  - 🔐 "GET_ACCOUNT" (address) -> get account
  - 🔐 "LOG_IN"
  - 🔐 "DELETE"
  - 🔐 "UPDATE" (kv)
  - 🔐 "WITHDRAW" (amount)

- /api/v1/profile

  - "GET" (address)
  - 🔐 "FOLLOW" (address)
  - 🔐 "UNFOLLOW" (address)

- /api/v1/initiative

  - "GET"
  - 🔐 "CREATE"
  - 🔐 "FOLLOW"
  - 🔐 "UNFOLLOW"
  - 🔐 "DELETE"
  - 🔐 "FUND"
  - 🔐 "EDIT"
  - 🔐 "WITHDRAW"
  - 🔐 "UP_VOTE"
  - 🔐 "DOWN_VOTE"
