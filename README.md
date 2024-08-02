# Fo UTILS

## Strapi Message Code
| Action                   | Example |
--------------------------|---------|
| Send strapi message code | POST https://dev-cms.finan.one/api/message-code   |

Header: Authorization - Beader + Strapi Token \
Request JSON BODY

    Code:        0,  // << required
    HttpCode:    0,  // << required
    Message:     "", // << required/unique
    Description: "",