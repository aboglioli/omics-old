# Omics

## Services
- users
  - service.go
    - Register
    - Validate
    - Login
    - Deactivate
    - Update
    - GetByID
    - RecoverPassword
    - ChangePassword
    - Follow
    - Unfollow
- publications
  - service.go
    - Search
    - GetByID
    - Publish
    - Update
    - Accept
    - Reject
    - Delete
    - Favorites
    - Like
    - Unlike
    - Bookmark
    - Unbookmark
    - Rate
- collections
  - service.go
    - GetByID
    - Create
    - Update
    - Delete
- contract
  - service.go
    - Request
    - Accept
    - Reject
    - Cancel
    - GenerateSummary
    - GenerateStatistics
  - repository.go
- payment
  - service.go
    - Donate
    - Subscribe
    - CancelSubscription
    - PaySubscription
    - PayAuthors
- notifications
  - service.go
    - Notify...
- reports
  - service.go
    - GenerateReport
- infrastructure
  - api
  - repositories

## API

### User
- /users/validate/:code
- /users
  - /login
  - /register
  - /:id/publications
  - /:id/subscription
    - /current
  - /:id/follow

### Publication
- /catalogue
- /search
- /publications
  - /:id/like
  - /:id/review
  - /:id/view
  - /:id/follow
  - /:id/contracts
    - /current
  - /:id/summaries
    - /current
- /collections
  - /:id/follow
- /image
- /categories
- /tags

### Reports

## TODO

- Validations in entities
- DTOs generated from main domain entities
- Container for all components (Configuration, db, services, repositories, etc.)
  - Should have methods for lazy loading
