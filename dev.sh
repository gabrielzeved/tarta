# /bin/bash

wgo -file=.go -file=.templ -xfile=_templ.go go generate :: go run main.go & \
browser-sync start \
  --files './**/*.go, ./**/*.templ' \
  --ignore '*_templ.go' \
  --port 3001 \
  --reloadDelay '2500' \
  --proxy 'localhost:3000' \
  --middleware 'function(req, res, next) { \
    res.setHeader("Cache-Control", "no-cache, no-store, must-revalidate"); \
    return next(); \
  }'

# nodemon
# yarn nodemon -e go,templ --ignore '*_templ.go' --exec 'templ generate && go run .'