# redpill

Get info about the current docker container

## CLI 

If the CLI is executed in a container, it will return the container ID:

    $ redpill
    ed807a7d59accf3b9e70c05d52cc494fc125377b82e472591211a83729c4b566

## API

### GetContainerID 

This will return the currently executing container ID, or an empty string.

    GetContainerID() (string, error)

### InContainer 

This will return true if executing in a container.

    InContainer() (bool, error)
