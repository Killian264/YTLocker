import React from "react";
import { VideoListItem } from "../components/VideosListItem";
import { useVideo } from "../shared/api/useVideo";
import { Playlist } from "../shared/types";
import { BuildVideoPlaylistUrl } from "../shared/urls";

export interface VideoListItemControllerProps {
	className?: string;
	videoId: number;
	playlist: Playlist;
}

export const VideoListItemController: React.FC<VideoListItemControllerProps> = ({
	className,
	videoId,
	playlist,
}) => {
	const [loading, video] = useVideo(videoId);

	if (loading || video == null) {
		return <div>Loading...</div>;
	}

	return (
		<VideoListItem
			video={video}
			url={BuildVideoPlaylistUrl(video.youtubeId, playlist.youtubeId)}
			playlistColor={playlist.color}
		></VideoListItem>
	);
};
