# lowply/warmcolors

[lowply.net/warmcolors](https://lowply.net/warmcolors/) - Used to be warmcolors.info

## Build and deploy

### 1. Install Ruby

Install [rbenv](https://github.com/rbenv/rbenv) and
[ruby-build](https://github.com/rbenv/ruby-build). On macOS, configure Xcode
and accept its license before installing Ruby:

```sh
sudo xcode-select --switch /Applications/Xcode.app
sudo xcodebuild -license accept
rbenv install
```

The Ruby version is specified in `.ruby-version`.

If Ruby was installed before configuring Xcode, uninstall it with
`rbenv uninstall 4.0.6` and run `rbenv install` again.

### 2. Install dependencies

```sh
bundle install
```

### 3. Build

```sh
./script/build
```

### 4. Develop locally

```sh
./script/server
```

Open <http://127.0.0.1:4000/warmcolors/>.

### 5. Deploy

```sh
./script/deploy
```
