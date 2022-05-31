const PROXY_CONFIG = {
  "/api/": {
    "target" :"http://"+process.env.BACKEND_HOST+":8080",
    "secure": false,
    "logLevel": "debug"
  }
}

module.exports = PROXY_CONFIG;

