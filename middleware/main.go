package middleware

import "time"

const LOCALS_USER_ID = "user_id"
const CACHE_KEY = "Cache-Time"

// cache nothing by default
const CACHE_DEFAULT_TIME = time.Duration(0)
