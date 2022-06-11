This kata is a simple web server written in golang.

## Hierarchy
Below is the hierarchy of route to file.
```mermaid
    flowchart LR

    server([Server])
    server --> root_route(["/"])
    server --> hello_route([/hello])
    server --> form_route([/form])

    root_route --> index_html([index.html])
    hello_route --> hello_func([hello func])
    form_route --> form_func([form func])

    form_func --> form_html([form.html])
```