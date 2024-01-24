# flugo-api-golang

## Initial project setup

Create .env (or .dev.env) file. Project should be launched with the correct environment set.

## Formatters setup

Prettier extension should be installed. And the .prettierrc file should be created in the project root.

The contents of the .prettierrc should be as follows:

```json
{
  "plugins": ["prettier-plugin-go-template", "prettier-plugin-tailwindcss"],
  "overrides": [
    {
      "files": ["*.html"],
      "options": {
        "parser": "go-template"
      }
    }
  ]
}
```

### Golang Templates formatter

https://github.com/NiklasPor/prettier-plugin-go-template

Command to install prettier plugin: `npm install --save-dev prettier prettier-plugin-go-template`

### Tailwind CSS formatter

https://tailwindcss.com/docs/editor-setup

Command to install prettier plugin: `npm install --save-dev prettier-plugin-tailwindcss`
