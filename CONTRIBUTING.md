# How To Contribute

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

We welcome contributions from external members of the team! We encourage you to float the idea for your PR first, via a GitHub issue or a post in Slack.

## Figma export-go instructions

- For export-go, the Figma export tool, see [export-go/README](export-go/README.md).

- After every new Figma export, we need to manually copy the icons to the `flight-icons/icons` directory right now.

```bash
rm flight-icons/icons/*
```

```bash
cp export-go/dist/* flight-icons/icons
```

## sprite instructions

- To generate the sprite file, [sprite.svg](ember-flight-icons/public/icons/sprite.svg), follow the instructions in the README of [edenspiekermann/sprite.sh](https://github.com/edenspiekermann/sprite.sh) e.g.

```bash
cd flight-icons/icons
```

```bash
npx spritesh
```

The `spitesh` command will great a `sprite.svg`. Copy this file to ember-flight-icons/

```bash
cd .. && cp flight-icons/icons/sprite.svg ember-flight-icons/public/icons && cp flight-icons/icons/_catalog.json ember-flight-icons/public/icons
```

## npm instructions

### @hashicorp/ember-flight-icons

```bash
cd ember-flight-icons
```

Bump the version number for the `ember-flight/package.json`.

After the change is merged to `main`, from the `ember-flight-icons/` directory, run:

```bash
npm publish
```

You will need 2FA on your npm account to publish.

### @hashicorp/flight-icons

```bash
cd flight-icons
```

- Bump the version number for the `flight-icons/package.json`.

After the change is merged to `main`, from the `flight-icons/` directory, run:

```bash
npm publish
```

You will need 2FA on your npm account to publish.

## How to test local addon changes

- `cd flight/ember-flight-icons` (or cd flight/flight-icons)
- Run `yarn link`. You'll get a response such as:

```bash
success Registered "@hashicorp/ember-flight-icons".
info You can now run `yarn link "@hashicorp/ember-flight-icons"` in the projects where you want to use this package and it will be used instead.
✨  Done in 0.06s.
```

(If necessary, run a `yarn unlink`.)

- In your external repo, e.g. https://github.com/hashicorp/design-system-playground-fastboot, run `yarn link "@hashicorp/ember-flight-icons"`
- In your external repo, manually add the path to the `package.json`. For example:

```json
"devDependencies": {
  "@hashicorp/ember-flight-icons": "link:~/your-path-here/flight/ember-flight-icons",
}
```

- Run `yarn` or `yarn install`
- You may need to copy code such as https://github.com/hashicorp/flight/blob/main/ember-flight-icons/tests/dummy/app/templates/application.hbs into the external app's `application.hbs` to see the results.
- If you want to test local changes to `ember-flight-icons`, add `isDevelopingAddon` to `ember-flight-icons/index.js`. The file will look something like the following:

```js
'use strict';

module.exports = {
  name: require('./package').name,
  isDevelopingAddon() {
    return true;
  },
};
```
