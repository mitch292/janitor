## TODO (README)

- Cover .janitor config file structure
- Cover .janitors_closet (all managed files are symlinked there)
- Cover .janitors_closet/.janitors_error_log
- Can mention initial use case for managing public keys

## TODO (Code)

- Sync
  - Test empty files
  - Handle directories as source
  - Handle when desitnation is directory..previously wrote file, but now force exact destination.
  - Handle when destination directory doesn't exist (should create it?)
  - Handle when we don't have permission to destination (just write to error log)
  - Handle when source file doesn't exist
  - Handle when we don't have access to the source file
  - Create Overwrite flag, on if we should overwrite the file if it exists
  - Allow people to pass the source / destination via the command line rather than config file
  - Support remote locations that require authentication.
- Mop
  - Remove any dead symlinks from the janitors_closet
- Transh
  - All - trash all files we manage (maybe just symlinks?)
  - Single file or dir can be passed as well
