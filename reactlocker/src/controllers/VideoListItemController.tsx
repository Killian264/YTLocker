import React from "react";
import { LoadingListItem } from "../components/LoadingListItem";
import { VideoListItem } from "../components/VideosListItem";
import { useVideo } from "../hooks/api/useVideo";
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
	const [isLoadingVideo, video] = useVideo(videoId);

	if (isLoadingVideo || video == null) {
		return <LoadingListItem></LoadingListItem>;
	}

	return (
		<VideoListItem
			className={className}
			video={video}
			url={BuildVideoPlaylistUrl(video.youtubeId, playlist.youtubeId)}
			color={playlist.color}
		></VideoListItem>
	);
};
