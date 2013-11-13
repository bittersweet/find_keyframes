# find_keyframes

A useful tool/script that uses [ffprobe](http://www.ffmpeg.org/ffprobe.html) to
find the nearest keyframe so you can split movie files efficiently.

## Installation

``` bash
go build
```

optionally move it to `/usr/local/bin/`

## Usage

``` bash
$ find_keyframes input_file HH:MM:SS
```

If you want to split a movie at the 00:05:38 mark:

``` bash
$ ./find_keyframes input.avi 00:05:38
Processing input.avi
Input HH:MM:SS: 00:05:38
Start in seconds: 338.000000
Keyframes: 62
Closest: 331.959991 Next: 338.399994
```

Then to split the movie file, make sure you start a bit before the keyframe:

``` bash
$ ffmpeg -i input.avi -ss 331.000000 -acodec copy -vcodec copy output.avi
```



