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

## todo december 2025
- prepare for distribution to:
  - homebrew
  - winget
  - apt (self hosted repo, because debian sponsorship is too slow)
  - regular github releases -> links on website (although discourage using github releases, make users use homebrew, winget, apt)
  - installation instructions on website very nice
- make sklair actually more of a cli tool
  - think in terms of subcommands:
    - `sklair init` -> creates a sklair.json file in the current directory and answer a questionnaire
    - `sklair build` -> builds the website based on sklair.json file or default values (if no sklair.json then warn about defaults available on docs)
    - `sklair serve` -> starts a local dev server, watching for changes and auto rebuild also ensure that its not actually built EVERY time theres a change
    - `sklair clean` -> removes all build artifacts (build dir, static dir)
    - `sklair version` -> prints the current version of sklair (TODO: strict semantic versioning!!)
    - `sklair --help` or `sklair help` -> gives help
    - `sklair update` -> updates sklair to latest version (ALSO: ensure that on every run of sklair, it notifies the user of a new version unless auto update check is disabled in sklair config)
    - `sklair docs` -> opens sklair docs in browser
    - `sklair config` -> opens sklair config file in default editor. sklair config file should be in HOME/.sklair.config.json
    - to make sklair more maintainable as a CLI tool with subcommands, take application commands base from CommandRegistry in other numelon-proprietary projects
      - adapt it for use with cli, each subcommand is registered etc just like in CommandRegistry
  - improve CLI UX. perhaps at some point think of replacing our fancy "logs" (glorified prints) with a spinner animation and progress bar. spinner for doc/static discovery, then progress bar for compilign and copying static files etc.
  - then only finally print new empty line and then print build time stats etc (summary)
  - also add --silent flag to suppress all output except errors, perfect for ci/cd (todo: github actions for numelon (bespoke) websites) - PRIORITY: this is actually required in the short term!!!
- search for "TODO" in the entire project and attempt to fix all of those
- ensure that in main.go the default sklair.json fallnback is NOT src/sklair.json but rather just sklair.json. or just test both?
- long term: allow sklair to integrate third party stuff like tailwind compilation: sklair scans html, sees which classes are used, compiles css. likewise also scans css for tailwind class usage and adds them to css just in case, so that its also programmable.