## Endpoint Checklist

This checklist is only a reference generated by chatGPT, and may be used when I want to convert this project to microservices/DDD project

### Public Endpoints (Exposed via API Gateway)

#### Aggregated (Unified) Endpoints
<!--- [ ] **Aggregated Entry List**  -->
<!--  `GET /api/aggregated/leaderboard/:leaderboardID/entries`  -->
<!--  *Returns entries for a leaderboard along with enrichment such as leaderboard field definitions and user details (for sorting, hidden fields, etc.).*-->

- [x] **Unified Leaderboard View**  
  `GET /api/aggregated/leaderboard/:leaderboardID`  
  *Returns leaderboard details, fields/options, entries, and enriched user information.*

- [x] **Aggregated Leaderboard Configuration**  
  `GET /api/aggregated/leaderboard/:leaderboardID/config`  
  *Returns complete leaderboard settings including fields, options, and external link metadata.*

- [x] **Aggregated Entry View** *(Optional)*  
  `GET /api/aggregated/entry/:entryID`  
  *Returns a single entry along with its associated leaderboard field data for proper rendering.*

- [x] **Aggregated Favorites**  
  `GET /api/aggregated/user/:userID/favorites`  
  *Returns a user’s favorite leaderboards enriched with basic leaderboard details.*

#### Domain-Specific CRUD Endpoints

**Leaderboard & Configuration:**
- [x] `GET /api/leaderboards`  
- [x] `POST /api/leaderboards`  
- [x] `PUT /api/leaderboards/:leaderboardID`  
- [ ] `DELETE /api/leaderboards/:leaderboardID`
<!--- [ ] `GET /api/leaderboards/:leaderboardID`  -->

<!--**Field & Option Management (within Leaderboard):**-->
<!--- [ ] `POST /api/leaderboards/:leaderboardID/fields`  -->
<!--- [ ] `GET /api/leaderboards/:leaderboardID/fields`  -->
<!--- [ ] `PUT /api/leaderboards/:leaderboardID/fields/:fieldID`  -->
<!--- [ ] `DELETE /api/leaderboards/:leaderboardID/fields/:fieldID`  -->
<!--- [ ] `POST /api/leaderboards/:leaderboardID/fields/:fieldID/options`  -->
<!--- [ ] `GET /api/leaderboards/:leaderboardID/fields/:fieldID/options`-->

**Entry Management:**
- [x] `POST /api/entries` *(May include orchestration/validation using leaderboard field definitions)*  
- [x] `GET /api/entries/:entryID` *(Could be the raw endpoint; the aggregated view is provided separately)*  
- [x] `DELETE /api/entries/:entryID`  
<!--- [ ] `PUT /api/entries/:entryID`  --> <!-- Entry not allowed to be edit-->

**User Management:**
- [x] `GET /api/users/:userID`  
- [x] `PUT /api/users/:userID`  
<!--- [ ] `DELETE /api/users/:userID`-->
<!--- [ ] `POST /api/users`  -->

**Verifier Management (if separate):**
- [x] `GET /api/leaderboards/:leaderboardID/verifiers`  
- [x] `POST /api/leaderboards/:leaderboardID/verifiers`
- [x] `DELETE /api/leaderboards/:leaderboardID/verifiers`

**Favorites Management:**
- [x] `POST /api/favorites`  
- [x] `DELETE /api/favorites/:userID/:leaderboardID`
<!--- [ ] `GET /api/favorites/:userID`  -->

**Auth Management:**
- [x] `POST /api/auth/login`  
- [x] `POST /api/auth/signup`  
- [x] `POST /api/auth/refresh`  
- [ ] `PUT /api/auth/password`  
- [ ] `PUT /api/auth/email`  
- [ ] `GET /api/auth/resetpassword`  
- [ ] `DELETE /api/auth/account`

<!------->
<!---->
<!--### Internal Endpoints (For Microservice-to-Microservice Communication Only)-->
<!---->
<!--**Leaderboard Service:**-->
<!--- [ ] `GET /internal/leaderboards/:leaderboardID`  -->
<!--  *Returns raw leaderboard configuration (fields, options, external links) for use by the aggregator.*-->
<!---->
<!--**Entry Service:**-->
<!--- [ ] `GET /internal/entries?leaderboardID=:leaderboardID`  -->
<!--  *Returns raw entry data for aggregation purposes.*  -->
<!--- [ ] `GET /internal/entries/:entryID`  -->
<!--  *Returns raw data for a single entry.*-->
<!---->
<!--**User Service:**-->
<!--- [ ] `GET /internal/users/mapping`  -->
<!--  *Returns a mapping of user IDs to usernames and/or additional user details (for enriching aggregated responses).*  -->
<!--- [ ] `GET /internal/users/:userID`  -->
<!--  *Returns raw user details.*-->
