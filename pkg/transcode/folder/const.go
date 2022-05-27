package folder

const (
	masterHeader  = "#EXTM3U\n#EXT-X-VERSION:3\n"
	subtitleGroup = "subtitle"
	stereoGroup   = "stereo"

	codecs = "\"avc1.42c00d,mp4a.40.2\""

	stereoTemplate = "%s#EXT-X-MEDIA:TYPE=AUDIO,GROUP-ID=\"%s\",LANGUAGE=\"%s\",NAME=\"%s\"," +
		"DEFAULT=%s,AUTOSELECT=YES,URI=\"audios/%s/audio.m3u8\"\n"
	subtitleTemplate = "%s#EXT-X-MEDIA:TYPE=SUBTITLES,GROUP-ID=\"%s\",LANGUAGE=\"%s\"," +
		"NAME=\"%s\",DEFAULT=%s,AUTOSELECT=YES,FORCED=NO,URI=\"subtitles/%s/sub.m3u8\"\n"
	videoTemplate = "%s#EXT-X-STREAM-INF:BANDWIDTH=%s,CODECS=%s,RESOLUTION=%dx%d," +
		"AUDIO=\"%s\",SUBTITLES=\"%s\"\nvideos/%dp/video.m3u8\n"
)
