# Connect

An advertising backend server and its web interface.
Built for SFC CLIP.

# Usage

- `/console`
  - Web interface
- `/api`
  - The root of API, which follows [JSON API](http://jsonapi.org) scheme
  - Endpoints:
    - `POST/GET/PATCH/DELETE /api/units`
    - `POST/GET/PATCH/DELETE /api/groups`
- `GET /any/{GroupID}`
  - Return an ID of the fewest accessed unit in the specified group
- `GET /img/{UnitID}`
  - Redirect to the registered image of the specified unit,
    and record an access
- `GET /open/{UnitID}`
  - Redirect to the registered URL of the specified unit,
    and record an access

# Requirements

- go
- Node.js
- MySQL / SQLite3

# Development

```sh
cd console
npm install
npm run dev
cd ..

go run main.go
```

# Build and run

```sh
cd console
npm install
npm run build
cd ..

go build -o connect .
./connect --port=PORT --production=1
```
