# pic2ascii

[![Build Status](https://travis-ci.org/wzshiming/pic2ascii.svg?branch=master)](https://travis-ci.org/wzshiming/pic2ascii)
[![Go Report Card](https://goreportcard.com/badge/github.com/wzshiming/pic2ascii)](https://goreportcard.com/report/github.com/wzshiming/pic2ascii)
[![Docker Automated build](https://img.shields.io/docker/automated/wzshiming/pic2ascii.svg?maxAge=2592000?style=plastic)](https://github.com/wzshiming/pic2ascii/)
[![GitHub license](https://img.shields.io/github/license/wzshiming/pic2ascii.svg)](https://github.com/wzshiming/pic2ascii/blob/master/LICENSE)

- [English](https://github.com/wzshiming/pic2ascii/blob/master/README.md)
- [简体中文](https://github.com/wzshiming/pic2ascii/blob/master/README_cn.md)

## Requirements

Go version >= 1.5.

FFmpeg version >= 3.X

## Download and install

``` shell
# Not support video
go get -u -v github.com/wzshiming/pic2ascii/cmd/pic2ascii

# Support video (Depends ffmpeg)
brew install ffmpeg # mac
go get -tags=support_video -u -v github.com/wzshiming/pic2ascii/cmd/pic2ascii
```

or

[Download releases](https://github.com/wzshiming/pic2ascii/releases) Not support video (Compile or use docker image if necessary.)

[Docker image](https://hub.docker.com/r/wzshiming/pic2ascii/) Support video

## Usage

### example1

``` shell
pic2ascii -f ./demo/src.gif -w 80 -h 40 -r
```

or

``` shell
docker run --rm -it wzshiming/pic2ascii -f https://github.com/wzshiming/pic2ascii/blob/master/demo/src.gif?raw=true -w 80 -h 40 -r -t gif
```

![src](https://github.com/wzshiming/pic2ascii/blob/master/demo/src.gif?raw=true)
![dist](https://github.com/wzshiming/pic2ascii/blob/master/demo/dist.gif?raw=true)

### example2

``` shell
pic2ascii -c "MMWNXK0Okxou=:\"'.  " -f https://avatars0.githubusercontent.com/u/6565744 -w 90 -h 40
```

or

``` shell
docker run --rm -it wzshiming/pic2ascii -c "MMWNXK0Okxou=:\"'.  " -f https://avatars0.githubusercontent.com/u/6565744 -w 90 -h 40
```

[![pic](https://avatars0.githubusercontent.com/u/6565744)](https://github.com/wzshiming)

``` log
  ..  .
kx=::'
MMMMMNo.
MMMMMMMXu
MMMMMMMMMx'
MMMMMMMMMMMXk="..
MMMMMMMMMMMMMMMMXo'
MMMMMMMMMMMMMMMMMMWKko=".                                                             ..':
MMMMMMMMMMMMMMMMMMMMMMMMWK=..                                                    . .uk0NWM
MMMMMMMMMMMMMMMMMMMMMMMMMMMXku"  .                                            '"u0XWMMWMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNKOo".    ..                               .=kXWMMMMMMNNNWM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0ko:''.                         .':x0NMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0x=:'..              ."u0WMMMMMMMMMMMMMNkooOWM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNXK0Okou==:":uxXMMMMMMMMMMMMMWk:''   .xM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWMMMMMMWk" .     .oM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMKu' .       'XM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM=""kMMMMMMMMN..     .    oMM
MMMMMMMMMMMMMMMMMMNXXK00Okkouooooxk0NMMMMMMMMMMMMMMMMMMMMWXko"   'KMMMMMMMk.          'KMM
MMMMMMMMMMMWNKOo:'.         .     . .=kXMMMMMMMMMMMMMNKk:..      :XMMMMMMMO .        'OMMM
MMMMMMNKx"'.....                        '=kNMMWKOxu:'            'OMMMMMMMX.        .uWMMM
MMXk='.  ...=ONNK'                      .  "NMO .                 :KMMMMMMN"     .=xXMMMMM
=".      "okkuOMM0.           ..."=xx='   . oMW=.                  =WMMMMMW:    'KMMMMMMMM
        "XK"..uMMM"          ..xKN0o0M0"    .OMX".                 :NMMMMMM"  . "MMMMMMMMN
       .x:. . kMMMk.       .=0Ku"..uMMM0'.   "NMk.                 "NMMMMMW'  . .oONMMMXx'
    . .o0. .. :NMMx   .   .x0=..  'XMMMN".    kMW=                 :XMMMMMX'      ..:=='
  .  .:0:  . '=WMW"       .=. . ..oMMMNu.    ."MMO.                .kMMMMM0.      .
           .  "==' ..   .    .   .kNWO"    .  'NMN"              . "0MMMMMX
            .:". .===:""'.     '  ..'         .OMW:              .xNMMMMMNu..
          .:u:'  'xO0KKK00k:  .o"             .oWM:            .'OMMMMMMK:.
         'uu"""""::::::"""'.  .ou..           .xMM:           'oXMMMMMMX'
        .uo"0kuooooxxkkO00X0ku":O:            .OMW:           OMMMMMMMW:  .
        .uo'uXKkx=::"""":==o00o=x=.          ."NMX.       . "kMMMMMMMx'
         =x''oXMMMMMMMMMMMMK=.'o='.           xMMo         =NMMMMMMXu.
         .==  .=oxkO000Okx:.':u..           .oWMX:     . "xWMMMMMK=..
.                 .    . ..=:' .            kMMO..     .xMMMMMMXu
0o.                       ''            ."uKMXo.     .=0MMMMMMx'
MMKu.                                 ."kWMNx' .     =WMMMMNk:..
":".                                  ."xo:'         :xkkkx:.
                                       .
       ..         ..      ..       ...

```

## Support formats

- [x] Ascii
- [ ] Color (unix like & windows)
- [x] Picture
  - [x] jpeg
  - [x] png
  - [x] bmp
  - [x] tiff
  - [x] webp
  - [x] gif
- [x] Video (Depends ffmpeg)
  - [x] mp4
  - [x] ts
  - [x] rtmp
  - [x] rtsp
  - [x] flv
  - [x] aac


## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/wzshiming/pic2ascii/blob/master/LICENSE) for the full license text.
