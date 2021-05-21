import React from "react";
import { VideoListItem } from "../components/VideosListItem";
import { usePlaylists } from "../shared/api/usePlaylists";
import { useVideo } from "../shared/api/useVideo";
import { BuildVideoPlaylistUrl } from "../shared/urls";

export interface VideoListItemControllerProps {
	className?: string;
	videoId: number;
	playlistYoutubeId: string;
}

export const VideoListItemController: React.FC<VideoListItemControllerProps> = ({
	className,
	videoId,
	playlistYoutubeId,
}) => {
	const [loading, video] = useVideo(videoId);
	const [loadingP, playlists] = usePlaylists();

	if (loading || video == null) {
		return <div>Loading...</div>;
	}

	if (loadingP || playlists == null) {
		return <div>Loading...</div>;
	}

	const videos = playlists.map((playlist, index) => {
		if (!playlist.videos.includes(videoId)) {
			return "";
		}

		return (
			<VideoListItem
				key={index}
				video={video}
				url={BuildVideoPlaylistUrl(video.youtubeId, playlistYoutubeId)}
				playlistColor={playlist.color}
			></VideoListItem>
		);
	});

	return <>{videos}</>;
};
