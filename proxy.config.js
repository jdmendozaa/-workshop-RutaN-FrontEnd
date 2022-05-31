const PROXI_CONFIG = {
  "/api/": {
    "target" :"http://${BACKEND_HOST}:8080",
    "secure": false,
    "logLevel": "debug"
  }
}

module.export = PROXI_CONFIG;

