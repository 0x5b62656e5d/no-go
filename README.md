# no-go

## No as a service

Use it to deny access or say no with some seasoning.

## Usage

Endpoint: `https://no-go.benkou.dev`

### JSON Structure:

```json
{
    "success": true, // Boolean to indicate if the request was successful or not
    "data": {
        "lang": "en", // ISO 639 language code
        "no": "No. I'm tired." // The reason
    },
    "error": null, // Will contain an error message if something goes wrong
    "message": "Seasoned a no properly." // A response message
}
```

### Parameters:

- Add a query parameter `?lang=<CODE>` to choose one of the supported language. Defaults to English (`en`).

## Supported languages:

| ISO 639 language code | Support |
|-----------------------|---------|
|          en           |    ✅   |
|          de           |    ✅   |
|          zh           |    ✅   |

Legend:

- ✅: Fully supported
- 🟡: Partial support (Only a few reasons written)

## Self hosting:

Just run the container with the given 

### Environment variables:

```
PORT=8080 # Optional, the port to run the API on
```

## Credits:

### Inspiration

- Inspired by [hotheadhacker/no-as-a-service](https://github.com/hotheadhacker/no-as-a-service)
- Inspired by [noahdunnagan/nope-rs](https://github.com/noahdunnagan/nope-rs)

### Contributors

Contributors will be listed here.

## License

This is a "Do what you want" license.

You may use, modify, and distribute this project however you like. The only requirement: ***credit me somewhere***.