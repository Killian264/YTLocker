import React from "react";
import { VideoListItem } from "../components/VideosListItem";
import { useVideo } from "../shared/api/useVideo";
import { BuildVideoUrl } from "../shared/urls";

export interface VideoListItemControllerProps {
	className?: string;
	videoId: number;
}

export const VideoListItemController: React.FC<VideoListItemControllerProps> = ({ className, videoId }) => {
	const [loading, video] = useVideo(videoId);

	if (loading || video == null) {
		return <div>Loading...</div>;
	}

	return <VideoListItem video={video} url={BuildVideoUrl(video.youtubeId)}></VideoListItem>;
};
