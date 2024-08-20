# Fo UTILS

## Strapi Message Code
| Action                   | Example |
--------------------------|---------|
| Send strapi message code | POST https://dev-cms.finan.one/api/message-code   |

Header: Authorization - Beader + Strapi Token \
Request JSON BODY

    code:        0,  // << required
    http_code:    0,  // << required
    message:     "", // << required/unique
    description: "",

Using in FO-BUSINESS by this api

    POST {fo-business-prefix}/internal/migrate 

Inside that API function name InitMigrate adding message code \
Change ID if it's existed before \
Adding new message code to list to multi created


		{
			ID: "202407301359005",
			Migrate: func(db *gorm.DB) error {
				// Create tables without foreign key
				ids = append(ids, "202407301359005")
				// Create list message_code
				messageCodes := []messagecode.CreateMessageCodeReq{
					{
						Code:        0,  // << required
						HttpCode:    0,  // << required
						ViMessage:     "", // << required/unique
						EnMessage:     "", // << required/unique
						Description: "",
					},
				}
				for _, v := range messageCodes {
					_, err := h.messClient.PublishMessageCode(ctx, v)
					if err != nil {
						return err
					}
				}
				return nil
			},
		},