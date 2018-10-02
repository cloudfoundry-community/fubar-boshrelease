# FUBAR'd BOSH Release

## Instructions

This BOSH release is broken in a number of ways. Its purpose is
to provide an exercise for people new to BOSH to become familiar with how
BOSH releases work, through finding errors, fixing them, and deploying dev
releases, until the application is fully functional. Keep in mind, after the
[Contributing](#contributing) section, the documentation in this
README may, itself, be intentionally accurate and misleading --
documentation is a part of a BOSH release, after all.

To get started with this, follow the usage instructions to generate a manifest
and start deploying this bosh release. Keep iterating through any errors encountered
until it deploys. That's not the end of it though! You will be finally done when you
can visit the deployed proxy VM IP in the browser, and receive the
respone "You win!".

# Contributing

Since this FUBAR BOSH release is designed to be broken, please
**DO NOT** submit pull requests / issues for bugs you find.
Please do not upload any new blobs to the blobstore.  Such
contributions will be summarily closed.

The only contributions that will be accepted are those that
improve the broken-ness of this release, with an eye towards
teaching and instruction.  Clarifying failures, new problems for
new BOSH features, etc., are all acceptable.

## Usage

Normally, you would run the following commands to compile a
development build of this release, and upload it to your BOSH
director:

```
bosh create-release --force --timestamp-version
bosh upload-release
```

This step will need to be re-run after every change you make
locally, to get a new version of the compiled release available
for deployment.

Next, you would deploy:

```
bosh deploy -d fubar manifests/fubar.yml
```

**INSTEAD**, we've got a handy script for you to run:

```
./try-again
```

This prefixes the name of the release / deployment with your
username.  For example, if `sam` runs it, the release will be named
`sam-fubar`, and the deployment will be `sam-fubar`.

Once you've fixed all the bugs, you should be able to visit the IP
of the deployed proxy VM in your web browser, and you'll see "You win!".
