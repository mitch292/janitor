## What is janitor?

Janitor is a cli tool that will let you sync remote files to your local computer or server. This can help you keep files in sync from one single repository.

## Why should I use janitor?

Janitor's orignal use case was for managing public SSH keys, or the `allowed_hosts` file on a server. Sometimes we can have many servers that need to keep a single file in sync. In addition, sometimes using a shared drive is not an option. So with janitor, you can store a single file in static storace (like amazon S3) and keep it in sync on all of your servers.

## How do I keep my files in sync with janitor?

Run `janitor sync` on a cron job! This will keep your files up to date on a time schedule that you define!

## How do I tell janitor what files to manage?

You use the `janitor.yml` file with a structure like below.

```yaml
files:
  myFile:
    source: "https://your-file-location.com/test.txt"
    destination: "/home/stras37/Documents/testJanitor.txt"
    safeMode: true
  anotherFile:
    source: "https://your-file-location.com/something_random.csv"
    destination: "/home/stras37/Documents/ws19mvp.csv"
  oneLastFile:
    source: "https://your-file-location.com/lets_go_nats.txt"
    destination: "~/Documents/letsGoNats.txt"
```

## What are the options for the `janitor.yml` file?

- Under files, name the each file you would like to have managed by the janitor.
- Each file will have a `source` and a `desitnation`
  - The `source`, is the rmote location to pull the file from.
  - The `destination`, is the locaiton on your local computer where the file should be stored.
  - If you add the `safeMode` option, you can indicate to janitor, whether you should overwrite an existing file or not. If `safeMode` is not present, `false` will be assumed and the local file "destination" will be overriden.

## Anything else I should know about Janitor?

- Janitor will manage all files through the "Janitor's closet". Located in the ~/.janitors_closet directory. Within this directory there will be symlinks to all of the files managed by Janitor.
- Within the janitors closet, there is a file .janitor_error_log, where you can find any errors that janitor has run into.

## Janitor CLI commands

- `sync` - Take the "source" files from your `.janitor.yml` file and "sync" them to the "destination" specified in your `.janitor.yml` file.
- `mop` - Janitor uses symlinks to keep track of all the files it manages. `mop` will remove all the dead symlinks from your `.janitor.yml` file.
- `trash` - Will remove all files managed by Janitor from your computer. It will not remove the `.janitors_closet`.
