# Typinger

## Idea
Typinger is simple and lightweight golang CMS.

## Usage
### Config
| Environment variable | Usage                                                             | Default                    |
| -------------------- | ----------------------------------------------------------------- | -------------------------- |
| HOST                 | IP of interface that is used for serving                          | -                          |
| POST                 | Port for server                                                   | 80                         |
| DOMAIN               | Domain is used for image URL generating                           | 0.0.0.0                    |
| DATABASE_URL         | URL for PostgreSQL                                                | -                          |
| WRITE_TIMEOUT        | Timeout for write network operations                                                                  | 10                         |
| READ_TIMEOUT         | Timeout for read network operations                                                                     | 5                          |
| IDLE_TIMEOUT         | Timeout for connection idle                                                                 | 120                        |
| JWT_SECRET           | JWT secret (not recomended to set for security reasons)                                                                  | &lt;randomly generated&gt; |
| IMG_DIR              | Virtual directory for images                                                                  | img                        |
| PROTOCOL             | Default protocol for serving for path generation for static files | http                       |