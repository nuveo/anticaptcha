# go-anticaptcha

Golang wrapper for [anti-captcha](https://anti-captcha.com/) online service which provides google recaptcha decodings.

This wrapper was implemented based on anti-captcha API [documentation](https://anticaptcha.atlassian.net/wiki/display/API/Documentation+in+English) and follow the steps described in [here](https://anticaptcha.atlassian.net/wiki/pages/viewpage.action?pageId=6029327).

As the documentation describes, the service receives a url and a recaptcha key, solves it and then return the `gRecaptchaResponse` key, that needs to be submitted in the website that has the captcha you are trying to solve, in the parameter `g-recaptcha-response`. The parameter name may be different, you can check more information on the following [link](https://anticaptcha.atlassian.net/wiki/display/API/Reproducing+Recaptcha+validation+without+digging+the+HTML+source).

Usually, the recaptcha key can be found in a parameter like that:

```
<div class="g-recaptcha" data-sitekey="6Lc_aCMTAAAAABx7u2W0WPXnVbI_v6ZdbM6rYf16"></div>
```

This [link](https://anticaptcha.atlassian.net/wiki/display/API/Reproducing+Recaptcha+validation+without+digging+the+HTML+source) shows how to find the key if it isn't in the format above.

Example usage:

```
package main

import (
    "fmt"
    "github.com/nuveo/anticaptcha"
)

func main() {
    // Go to https://anti-captcha.com/panel/settings/account to get your key
    c := &anticaptcha.Client{APIKey: "your-key-goes-here"}

    key := c.SendRecaptcha(
        "http://http.myjino.ru/recaptcha/test-get.php", // url that has the recaptcha
        "6Lc_aCMTAAAAABx7u2W0WPXnVbI_v6ZdbM6rYf16", // the recaptcha key
    )
    fmt.Println(key)
}

```

Here's the example for regular catpchas (image to text):

```
package main

import (
    "fmt"
    "github.com/nuveo/anticaptcha"
)

func main() {
    // Go to https://anti-captcha.com/panel/settings/account to get your key
    c := &anticaptcha.Client{APIKey: "your-key-goes-here"}

    text := c.SendImage(
        "your-base64-string", // the image file encoded to base64
    )
    fmt.Println(text)
}


```