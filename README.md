# chunky
Chunky is a helper tool to download TTV videos including subscribers only , deleted videos for free.
The tool depend on the https://pogu.live/ so if the website is down the tool will go down also.
> It's in the early stages of developement so many errors are excpected , also many manual job need to be done that will be automated later.

# Dependency
this tool require `ffmpeg` to run , you will get an error if ffmpeg is not detected in the path when launching tool , so please make sure you have ffmpeg installed properly in your system.

# Info
In our example I am going to use this [video](https://www.twitch.tv/videos/1199813680) because it's free on youtube , but not on twitch.
if we go to this URL we can see that it's for subscribers only

![image](https://user-images.githubusercontent.com/35725554/145275174-684dc376-06dc-4108-af56-9525849610aa.png)

now go to [pogu](https://pogu.live/) type the name of the twitch channel in the search box , we can see that the video is available

![image](https://user-images.githubusercontent.com/35725554/145275418-fb902c6f-c3f6-4215-b371-ea617176260e.png)

click on the link will open a player , solve capatcha wait few seconds and click clip it a small chunk will start to download cancel it and copy download link

```
https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/0.ts
```

in our case we got this link it doesn't matter number of chunk the tool will take care of that. Now time of my tool `chunky`

# Getting tool

you can download binaries from the release page or compile it by your self , to compile : 

first clone the repository
```bash
git clone https://github.com/khatibomar/chunky.git
cd chunky
```

now you need to have go installed if not please install using this [link](https://go.dev/dl/)

now in the terminal run

```bash
go mod tidy
go build
```

# Usage:

```bash
./chunky
```
> note: for windows users instead of ./chunky replace with chunky.exe in the cmd

now you will see options
```
chunky is a tool that will allow you to download subscribers only videos from twitch.
The tool is under current developement so many bugs will occur , and many missing features and many hard coded stuff.
So please if you found any bug or missing features feel free to open an issue on project page:
https://github.com/khatibomar/chunky

Usage:
  -dir string
        specify a download path , for *nix users use $HOME instead of ~ . In case no absolute path specified the folder will be created in same dir as the tool folder
  -dwn
        by default true , false if you just want to get chunks size without downloading files (default true)
  -name string
        the name you want to save the video with without .mp4
  -url string
        provide a link that have a chunk , example:
        https://d2nvs31859zcd8.cloudfront.net/70c102b5b66dbeac89e4_channel_name_blaabllablablabl/chunked/X.ts
```

those are the options for now

```bash
./chunky -name=bb -dir=$HOME/BabelOnBabylyon -url=https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/0.ts 
```

```
Info : 2021/12/10 06:02:31 Nb of chunks is 1026 , from 0 to 1025
Info : 2021/12/10 06:02:31 Downloding /home/okpc/BabelOnBabylyon/bb0.ts...
Info : 2021/12/10 06:02:32 Downloding /home/okpc/BabelOnBabylyon/bb1.ts...
Info : 2021/12/10 06:02:33 Downloding /home/okpc/BabelOnBabylyon/bb2.ts...
^C
```
as we can see the app guessed the number of chunks very fast and start downloading chunks , after the download complete , the app will assemble all the parts to the name you specificed in `-name=...` into an assembled mp4 video , and will clean all the chunks and will only keep the mp4 video in the directory you specified with `-dir=...`
