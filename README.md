# go-ffmpeg-practice

Practice of video processing with [ffmpeg-go](https://github.com/u2takey/ffmpeg-go), especially video compression with Codec H.265.

## Requirements

- docker compose

## How to run

The container includes Go and ffmpeg binary.

```bash
$ docker compose up --build && docker compose rm -fsv
```

Check the output video file `test_output.mp4` in the `media/` directory.

## Videos used in this project

- Pexels Video (test_video.mp4): https://www.pexels.com/video/a-woman-busy-writing-on-a-paper-4778723/
