# Todo list

## todo list transferred from `numelon-proprietary/website` repo

- create numelon web packing tool for creating websites
  - this web packing tool should have hot reload
  - allow for components
    - e.g. `/components/someComponent.html` and `/components/someComponent.css`
    - numelon web packing tool will search for components dir
      - if html found then it will replace the tag inside the html with the actual content of the html file
      - if css found for component in component dir then it will be automatically copied to the build directory into css dir and also the link href stylesheet injected into the head of the html file where the component was used
      - if component not found then just assume its a component registered from within js and leave it as-is, but issue a warning in the logs.
      - if js file is found then assume regular html component registered via js and add the js to head tag. ofc if css with same name is found then import that too
      - therefore the order from most important to least is like this:
        1. componentName.js -> import js (& optional css, if html is present then ignored) in head
        2. componentName.html -> replace all occurences of `<componentName/>` or `<componentName>` or `<componentName></componentName>` with contents of componentName.html. (& optional css in head)
            - note that if it is `<componentName>aaa</componentName>` then `aaa` would be passed to `$$COMPONENT_BODY` var inside componentName.html.
  - reusable component for menubar and footer
  - allow components to be used inside of eachother (in src) but hard fail on circular component usage bc infinite loop

- support for replacing $$COMPONENT_BODY inside of javascript too, since its only supported in html right now with `<!-- $$COMPONENT_BODY -->`

- create a JSON schema for `sklair.json` files:
- <https://json-schema.org/understanding-json-schema/reference/index.html>
- note that any file paths like `input` and `output` are RELATIVE to the sklair.json file
- allow components to be FULL html files with a head and body. If it is a full file (as opposed to just a regular component body that is bare), then basically build a "cache" of things from the head that will be inserted into the source document (deduplication) and the rest of the body just gets inserted as usual
- create separate timers for actually processing the files - ie file discovery, then compiling. then separate timer for copying static files since that heavily inflates the build time.