<img src="./resources/icon-xs.png">

# Double Tap
A small development utility that helps minimize Electron build sizes.

## Usage

1. Download the approapriate version from the [releases page](https://github.com/christopherwk210/double-tap/releases), and put it in its own folder.
2. Create a `resources` folder next to it, and inside there place your application code (either an `app` folder or `app.asar`)
3. Create a `target.json` in your `resources` folder, with an object that looks like this:

```json
{
  "version": "19.0.4"
}
```

where `version` is the Electron version you want to use.

4. Simply run the Double Tap executable and you will have a functioning Electron app. The first run may take a moment since it has to download the correct binary, but all future runs will be quick.

Your folder should look like this when it's all setup:
```
.
├── double-tap.exe
└── resources
    ├── app
    │   ├── your app code goes in here...
    └── target.json
```

You are free to rename the executable file to whatever you'd like.

## The problem
Electron builds, as we all know, are large. A "hello world" app can be 70-100MB+ on its own. For production builds this isn't really avoidable, and that's fine for the most part. But for sharing development builds back and forth with your team, this can be troublesome. Especially when you're iterating often. Uploading huge zip files multiple times a day just so someone can review some minor changes is annoying and dumb. There are plenty of solutions one could probably come up with, but for my circumstance developing Double Tap made sense enough. Also it was fun to make.

## How it works
Packaging an Electron app basically just means taking your application code and sticking it with an Electron binary and calling it a day. I figured, why not just use the same binary every time? That way we'd only have to send the updated application code and save some time. Well, that's what Double Tap does. Stick your application code next to the Double Tap binary, run it, and it will automatically download an Electron binary to a cached location. For all future runs using that same version, it'll reuse that same binary. Now you only need to resend the Double Tap binary which is only 8MB or so, instead of Electron which is around 100MB.

## Disclaimer
This was made just for quickly passing around development builds with your team or clients. It's probably not suited for production builds since it's lacking in the UX department. That being said, it would be cool for a strategy like this to be used for all Electron apps. Could save us all a lot of disk space or whatever. Also, this is the first thing I've ever made using Golang. It's a great language, but I probably didn't do things according to convention. There's plenty of room for improvement, but for now it works just fine for my workflow.

## Development
If you want to contribute, clone the repo and do a quick `npm i`. I use node for a small step in the Windows build process so I can use [node-rcedit](https://github.com/electron/node-rcedit) to update the executable icon. For that to work, you'll also need to have Wine installed, you can get it from homebrew:

```
brew install --cask wine-stable
```

Use `dev.sh` or `npm start` to run in development mode. When running in development mode, the application will look in `./temp` for your `resources` folder, so you'll have to create that folder and put some app in there.
