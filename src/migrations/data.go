package migrations

import (
    "github.com/drumer2142/microWeb/src/api/models"
)

var sites = []models.Site{
    models.Site{
        URL:          "https://www.google.com/",
        Domain:       "google",
        Description:  "",
    },
    models.Site{
        URL:          "https://www.facebook.com/",
        Domain:       "facebook",
        Description:  "",
    },
}
