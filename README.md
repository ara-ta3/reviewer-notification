reviewer notification
---

[![Build Status](https://travis-ci.com/ara-ta3/reviewer-notification.svg?branch=master)](https://travis-ci.com/ara-ta3/reviewer-notification)

[![](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/ara-ta3/reviewer-notification)

# Settings

## Github Webhooks

**Payload URL**

```
https://your-app-name.herokuapp.com/
```

**Content type*  

```
application/json
```

*Which events would you like to trigger this webhook?*  

Let me select individual events.  

- Issues
- Pull requests

## Environment Variable

### ACCOUNT_MAP

Create a mapping of github user and slack user to link them.

Example

```
ara-ta3:<@UXXXXXXX|arata>
```

Use comma to link multiple users.

```
ara-ta1:<@UXXXXXX1|arata1>,ara-ta2:<@UXXXXXX2|arata2>,ara-ta3:<@UXXXXXX3|arata3>,
```

### SLACK_CHANNEL

Where to notify.

Example

```
#random
```

### SLACK_WEBHOOK_URL

**NOTE!** Currently only legacy incoming webhook is supported.  

https://api.slack.com/legacy/custom-integrations/incoming-webhooks  

Example

```
https://hooks.slack.com/services/TXXXXXXX/YYYYYY/ZZZZZZ...
```

### TARGET_LABELS

The label to send notification.  
When this label is added your Pull Request, the notification will be sent.  

Example 

```
dummy-label
```

@see https://github.com/ara-ta3/reviewer-notification/labels  


