# restaurant-browser

To execute the project the only dependency is [podman](https://podman.io/).


## How to run

```bash
make run
```

_This will create a pod with 4 containers running. (App, Server, DB, PgAdmin)_

Then open http://localhost:9000/

## To stop and clean

```bash
make clean
```