# canvafix
Tool that fixes the Canva Bulk Create missing Audio bug on exported MP4 files

## Issue

When you create a Project that has will use Bulk Create to connect data and make multiple pages (from the first)... when you download (as seperate files ticked) only the first MP4 will have the audio. The rest of the MP4s are silent.

I have looked at the data & properties when making this tool and have noticed that they all have an AAC Audio stream which is the proper length, but onyl the first one has a good bitstream (greater than 17k), the others have a bad bitstream (less than 2.3k).

### Community Reports

- https://www.reddit.com/r/canva/comments/13txr11/bulk_create_not_retaining_audio/
- https://www.reddit.com/r/canva/comments/1755cp7/canva_bulk_create_with_audio_files/

## The Hack

So this tool will grab the good audio from the first exported file (1.mp4), cache it in memory then loop through all other MP4's that do not have a matching bitstram and replace them with the first.

I also added "prepend" argument that will allow you to better prepare the file for Youtube uploading.

so, if you want to set the title or your youtube upload and still have an index you can do something like this:

```sh
canvafix --p My youtube title
```

This will rename all files (processed) look like this:

1.mp4 -> My youtube title (1).mp4
...
20.mp4 -> My youtube title (20).mp4

## How to use

If you have golang installed you can clone this repo and build your own:


make help to display output below.

```sh
Usage:
  make windows      Build for Windows
  make linux        Build for Linux
  make mac          Build for macOS
  make clean        Remove all binaries and the bin directory
  make all          Build for all platforms
```

### Build release for all platforms

```sh
make
```

### Build release for a specific (your) platform

```sh
make windows
or
make linux
or
make mac
```

The compiled ready to use executables will be in the bin folder

### Download the binary

I will make each release available here on github. Look to your right to see the Releases section. Download for your OS, rename to canvafix (canvafix.exe for windows) and leave it in your downloads folder. Then copy it and paste it in your MP4's folder that need fixing and run it as stated above.

#### Latest Releases

- [Download for Windows](https://github.com/AndrewDonelson/canvafix/releases/download/0.1.0/canvafix-windows-amd64.exe)
- [Download for Linux](https://github.com/AndrewDonelson/canvafix/releases/download/0.1.0/canvafix-linux-amd64)
- [Download for MacOS](https://github.com/AndrewDonelson/canvafix/releases/download/0.1.0/canvafix-darwin-amd64)


### Final word

Its not perfect at this point, i have it working and stable with only a 3 hours of work in it but I plan to finish it up and polish a bit since it seems Canva is not going to fix this (over one year old) issue anytime soon.


#### known Issues

1. It does not "prepend" the first file so it will stay 1.mp4.
2. The binary has to be copied into the same folder as your MP4s then removed when done.

I will be working on both of these in my spare time and hopefuly have them done in a few days.


Hope this helped you
