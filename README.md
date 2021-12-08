# chunky
Chunky is a helper tool to download TTV videos including subscribers only , deleted videos for free.
The tool depend on the https://pogu.live/ so if the website is down the tool will go down also.
> It's in the early stages of developement so many errors are excpected , also many manual job need to be done that will be automated later.


# Usage
In our example I am going to use this [video](https://www.twitch.tv/videos/1199813680) because it's free on youtube , but not on twitch.
if we go to this URL we can see that it's for subscribers only

![image](https://user-images.githubusercontent.com/35725554/145275174-684dc376-06dc-4108-af56-9525849610aa.png)

now go to [pogu](https://pogu.live/) type the name of the twitch channel in the search box , we can see that the video is available

![image](https://user-images.githubusercontent.com/35725554/145275418-fb902c6f-c3f6-4215-b371-ea617176260e.png)

click on the link will open a player , solve capatcha wait few seconds and click clip it a small chunk will start to download cancel it and copy download link

```
https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/0.ts
```

in our case we got this link it doesn't matter number of chunk the tool will take care of that. Now time of my tool `chunky` , the tool job is to guess the number of chunks that this video have , I use a modified `binary search` algorithm to find the number of chunks so it's superfast

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
        specify a download path
  -dwn
        by default true , false if you just want to get chunks size without downloading files (default true)
  -max int
        provide the excpected max number of files, zero or negative numbers will be treated as max int (default -1)
  -url string
        provide a link that have a chunk , example:
        https://d2nvs31859zcd8.cloudfront.net/70c102b5b66dbeac89e4_handmade_hero_40072241627_1633745055/chunked/155.ts
```
The downloader at this point is very simple implementation so I encourage you to only use the tool to guess number of chunks and use external downloader.
So I decided to just get only number of chunks
```
./chunky -dwn=false -url=https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/0.ts -max=2000
```
max here is totally random number because I dont think the number of chunks will exceed the max value if you didnt specify max the max will try the biggest integer possible , In case of number of chunks is more than 2000 an message will be printed on screen so don't worry :)
```
Info : 2021/12/08 22:11:39 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/2000.ts
Info : 2021/12/08 22:11:40 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1000.ts
Info : 2021/12/08 22:11:41 Highest Guess: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1000.ts
Info : 2021/12/08 22:11:41 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1000.ts
Info : 2021/12/08 22:11:41 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1500.ts
Info : 2021/12/08 22:11:41 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1250.ts
Info : 2021/12/08 22:11:42 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1125.ts
Info : 2021/12/08 22:11:42 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1062.ts
Info : 2021/12/08 22:11:43 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1031.ts
Info : 2021/12/08 22:11:44 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1015.ts
Info : 2021/12/08 22:11:45 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1023.ts
Info : 2021/12/08 22:11:50 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1027.ts
Info : 2021/12/08 22:12:01 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1025.ts
Info : 2021/12/08 22:12:02 Trying: https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/1026.ts
Info : 2021/12/08 22:12:03 Nb of chunks from 0 to 1026
```
the app got the number of chunks with very minimal number of checks instead of trying 2000 link we only checked 14 link!!

now use any downloader you want to download chunks from `0` to number outputed by the app , in our case is `1026`
or delete `-dwn=false` in case you want to use built in installer
> size of the videos are big ~= 9GB

```
./chunky -url=https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/0.ts -max=2000 -dir=Video1
```
I specified directory in this case you will see a new folder created in the same folder of the tool.
After downloading all chunks successfully , now we are in the last step , combining all those chunks into 1 video
```
cd Video1
```
> NOTE: I didn't yet implemented the auto combining part but it's in the plan.

```
export NAME="[Handmade Hero] Babel On Babylion"
export LENGTH=1026
export BASEURL=https://d2nvs31859zcd8.cloudfront.net/a05ed2eeabcd1d053260_handmade_hero_40195752827_1636398126/chunked/
set -e
rm mylist.txt *.mp4
for i in $(seq 0 $LENGTH); do echo file $i.ts >> mylist.txt;done
ffmpeg -safe 0 -f concat -i mylist.txt -c copy all.ts
ffmpeg -i all.ts -acodec copy -vcodec copy all.mp4
mv all.mp4 "$NAME.mp4"
rm *.ts
```

after this step a new video with name of `[Handmade Hero] Babel On Babylion.mp4` will appear, ENJOY!!
