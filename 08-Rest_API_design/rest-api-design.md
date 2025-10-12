# REST API Design Guide

## Introduction to API Design and REST APIs

API design is a crucial activity for backend engineers, particularly focusing on REST APIs which dominate modern web services. This guide addresses common confusions and establishes best practices for REST API design.

### Key Areas Covered
- Resource and route design
- Success and error responses
- HTTP status codes
- Data acceptance patterns
- API documentation
- Payload design

## Historical Context and Origin of REST

### The Birth of the Web
- 1990: Tim Berners-Lee initiated the World Wide Web project
- Invented foundational technologies:
  - URI
  - HTTP
  - HTML
  - First web server and browser

### REST Architecture Development
Roy Fielding addressed web scalability challenges through REST (Representational State Transfer) in his 2000 PhD dissertation.

#### REST Constraints

| Constraint | Description | Purpose |
|------------|-------------|----------|
| Client-Server | Separation of UI/UX from data/logic | Independent evolution, improved scalability |
| Uniform Interface | Standardized communication via resources | Simplified architecture, consistency |
| Layered System | Hierarchical layers with limited interaction | Enhanced scalability and security |
| Cacheable | Explicit cache labeling | Reduced server load, better efficiency |
| Stateless | Self-contained requests | Improved reliability and scalability |
| Code on Demand (optional) | Server can send executable code | Added flexibility (rarely used) |

## Understanding REST Components

### REST Breakdown
1. **Representational**: Resources in various formats (JSON, HTML)
2. **State**: Current resource properties
3. **Transfer**: Moving resources via HTTP methods

## URL Structure and Best Practices

### URL Components

| Part | Description | Example |
|------|-------------|---------|
| Scheme | Protocol | https |
| Authority | Domain/subdomain | api.example.com |
| Path | Resource hierarchy | /books/123 |
| Query Params | Filters/pagination | ?limit=10&page=2 |
| Fragment | Client-side section | #section1 |

### URL Best Practices
- Use plural nouns (/books not /book)
- Use hyphens for slugs (harry-potter not harry_potter)
- Start with version (e.g., /v1/)
- Maintain hierarchical relationships

## Idempotency and HTTP Methods

### HTTP Methods Overview

| Method | Description | Idempotent? | Usage |
|--------|-------------|-------------|--------|
| GET | Retrieve data | Yes | Fetch resources |
| PUT | Replace resource | Yes | Full updates |
| PATCH | Partial update | Yes | Field updates |
| DELETE | Remove resource | Yes | Resource deletion |
| POST | Create/custom actions | No | Creation, custom operations |

### POST Method Usage
- Primary method for non-CRUD operations
- Used for custom actions
- Example: /users/{id}/send-email

## API Interface Design Workflow

### Design Process
1. Analyze UI wireframes and user stories
2. Identify resources (nouns)
3. Plan database schema
4. Design API routes
5. Document using tools (Insomnia, Postman)

## CRUD API Design Example: Organizations

### Standard Endpoints

| Operation | Method | Route | Response Code | Purpose |
|-----------|---------|-------|--------------|---------|
| List | GET | /organizations | 200 | Paginated list |
| Create | POST | /organizations | 201 | New resource |
| Get Single | GET | /organizations/{id} | 200 | Single resource |
| Update | PATCH | /organizations/{id} | 200 | Partial update |
| Delete | DELETE | /organizations/{id} | 204 | Resource removal |
| Archive | POST | /organizations/{id}/archive | 200 | Custom action |

### Pagination Features
- Parameters: limit, page
- Default values provided
- Metadata included in response

### Sorting and Filtering
- sortBy: field selection
- sortOrder: asc/desc
- Field-based filtering

## Error Handling
- Empty arrays return 200 for lists
- 404 for missing single resources
- Consistent error response format

## Custom Actions
- Use POST method
- Complex server-side operations
- Clear, descriptive endpoints

## Projects Resource Example

### Consistency Principles
- Plural resource names
- camelCase JSON keys
- Standard CRUD endpoints
- Similar pagination/sorting

### Custom Action Example
- POST /projects/{id}/clone
- Complex operations
- Returns 201 for new resources

## Best Practices Summary

### Documentation
- Use Swagger/OpenAPI
- Interactive testing
- Clear documentation

### Design Principles
- Intuitive interfaces
- Consistent naming
- Safe defaults
- Clear field names
- Design before coding

## Key Takeaways

1. Follow REST constraints for scalability
2. Use plural nouns consistently
3. Respect HTTP method semantics
4. Reserve POST for creation/custom actions
5. Implement pagination for lists
6. Maintain consistent naming
7. Design APIs before coding
8. Document thoroughly

This comprehensive guide provides backend engineers with the foundation needed to design scalable, maintainable, and user-friendly REST APIs that align with industry standards.