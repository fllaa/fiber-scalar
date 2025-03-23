package scalar

const indexTmpl string = `
<!doctype html>
<html>
  <head>
    <title>{{ .Title }}</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>

  <body>
    <div id="app"></div>

    <!-- Load the Script -->
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>

    <!-- Initialize the Scalar API Reference -->
    <script>
      Scalar.createApiReference('#app', {
        {{- if .ContentURL }}
        url: '{{ .ContentURL }}',
        {{- else }}
        content: {{ .Content }},
        {{- end }}
        proxyUrl: '{{ .ProxyURL }}',
        title: '{{ .Title }}',
        layout: '{{ .Layout }}',
        theme: '{{ .Theme }}',
        darkMode: {{ .DarkMode }},
      })
    </script>
  </body>
</html>
`
